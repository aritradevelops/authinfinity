package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/app"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/emailverificationrequest"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/password"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/middlewares/translator"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/logger"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	PasswordHashingCost = 10
)

type AuthService struct {
	core.Service[*user.User]
}

var authService *AuthService

func Service() *AuthService {
	return authService
}

func (s *AuthService) Register(c *fiber.Ctx) error {
	conf := config.Instance()
	account, err := account.Service().GetAccountFromReq(c)
	if err != nil {
		return core.NewNotFoundError(c)
	}

	// parse the body
	var payload RegisterPayload
	err = c.BodyParser(&payload)
	if err != nil {
		return core.NewBadRequestError(c)
	}

	// validate email
	errs := validator.Validate(payload, c)
	if errs != nil {
		return core.NewRequestValidationError(c, errs)
	}

	// validate password specially
	err2 := validator.ValidatePassword(c, payload.Password)
	if err2 != nil {
		return core.NewRequestValidationError(c, []validator.ValidationError{*err2})
	}

	userData := &user.User{
		Name:          payload.Name,
		Email:         payload.Email,
		EmailVerified: false,
	}

	// insert the user
	userID, err := user.Repository().Create(userData)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// handle unique email
			return core.NewDuplicateKeyError(c, "email")
		}
		return err
	}
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), PasswordHashingCost)

	if err != nil {
		return err
	}

	passwordData := &password.Password{
		Password: string(hashedPassword),
		UserID:   uuid.MustParse(userID),
	}
	passwordData.SetAccountID(account.ID.String())
	passwordData.SetCreatedAt()
	passwordData.SetCreatedBy(account.CreatedBy.String())

	// insert the password
	_, err = password.Repository().Create(passwordData)

	if err != nil {
		return core.NewInternalServerError(c)
	}
	secretBytes := make([]byte, 32)
	_, err = rand.Read(secretBytes)

	if err != nil {
		return core.NewInternalServerError(c)
	}
	duration, _ := time.ParseDuration(conf.Env.EmailVerificationHashExpiry)

	emailVerificationRequestData := &emailverificationrequest.EmailVerificationRequest{
		Hash:      hex.EncodeToString(secretBytes),
		UserID:    uuid.MustParse(userID),
		ExpiredAt: time.Now().Add(duration),
	}
	emailVerificationRequestData.SetAccountID(account.ID.String())
	emailVerificationRequestData.SetCreatedAt()
	emailVerificationRequestData.SetCreatedBy(account.CreatedBy.String())
	_, err = emailverificationrequest.Repository().Create(emailVerificationRequestData)
	if err != nil {
		return err
	}
	// trigger a mail

	logger.Info().Msg(fmt.Sprintf("Email verification link: https://%s/verify-email?hash=%s\n", c.Request().URI().Host(), emailVerificationRequestData.Hash))

	// respond to the user
	return nil
}

func (s *AuthService) VerifyEmail(c *fiber.Ctx) error {

	account, err := account.Service().GetAccountFromReq(c)
	if err != nil {
		return core.NewNotFoundError(c)
	}

	hash := c.Query("hash")

	if hash == "" {
		return core.NewRequestValidationError(c, []validator.ValidationError{
			{Message: translator.Localize(c, "validation.required", map[string]string{
				"Field": "hash",
			})},
		})
	}
	_, err = emailverificationrequest.Repository().GetActiveEmailVerificationRequest(
		account.ID.String(),
		hash,
	)
	return err
}

func (s *AuthService) Login(c *fiber.Ctx) error {
	// get the account
	conf := config.Instance()
	account, err := account.Service().GetAccountFromReq(c)
	if err != nil {
		return core.NewNotFoundError(c)
	}

	// parse the body
	var payload LoginPayload
	err = c.BodyParser(&payload)
	if err != nil {
		return core.NewBadRequestError(c)
	}

	// validate email
	errs := validator.Validate(payload, c)
	if errs != nil {
		return core.NewRequestValidationError(c, errs)
	}
	existingUser := &user.User{}
	err = user.Repository().View(core.Filter{
		"email":          payload.Email,
		"deleted_at":     nil,
		"email_verified": true,
	}, &existingUser)

	if err != nil {
		return err
	}
	app, err := app.Service().GetSysAdminApp()

	if err != nil {
		return err
	}

	// sign tokens
	return nil
}
