package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/app/modules/account"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/emailverificationrequest"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/password"
	"github.com/aritradevelops/authinfinity/server/internal/app/modules/user"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	core.Service[*user.User]
	userRepository                     *user.UserRepository
	passwordRepository                 *password.PasswordRepository
	accountService                     *account.AccountService
	emailVerificationRequestRepository *emailverificationrequest.EmailVerificationRequestRepository
}

func Service() *AuthService {
	var authRepository = Repository()
	return &AuthService{
		Service:                            core.NewService(core.Repository[*user.User](authRepository)),
		userRepository:                     user.Repository(),
		passwordRepository:                 password.Repository(),
		accountService:                     account.Service(),
		emailVerificationRequestRepository: emailverificationrequest.Repository(),
	}
}

func (s *AuthService) Register(c *fiber.Ctx) error {
	conf, err := config.Load()
	if err != nil {
		return err
	}
	account, err := s.accountService.GetAccountFromReq(c)

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

	user := &user.User{
		Name:  payload.Name,
		Email: payload.Email,
	}

	// insert the user
	userID, err := s.userRepository.Create(user)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// handle unique email
			return core.NewDuplicateKeyError(c, "email")
		}
		return err
	}

	// hash the password
	salt := 10

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), salt)

	password := &password.Password{
		Password: string(hashedPassword),
		UserID:   uuid.MustParse(userID),
	}
	password.SetAccountID(account.ID.String())
	password.SetCreatedAt()
	password.SetCreatedBy(account.CreatedBy.String())

	// insert the password
	_, err = s.passwordRepository.Create(password)

	if err != nil {
		return core.NewInternalServerError(c)
	}
	secretBytes := make([]byte, 64)
	_, err = rand.Read(secretBytes)

	if err != nil {
		return core.NewInternalServerError(c)
	}
	duration, _ := time.ParseDuration(conf.Env.EmailVerificationHashExpiry)

	emailVerificationRequest := &emailverificationrequest.EmailVerificationRequest{
		Hash:      string(secretBytes),
		UserID:    uuid.MustParse(userID),
		ExpiredAt: time.Now().Add(duration),
	}
	emailVerificationRequest.SetAccountID(account.ID.String())
	emailVerificationRequest.SetCreatedAt()
	emailVerificationRequest.SetCreatedBy(account.CreatedBy.String())
	_, err = s.emailVerificationRequestRepository.Create(emailVerificationRequest)
	if err != nil {
		return core.NewInternalServerError(c)
	}
	// trigger a mail

	fmt.Printf("Email verification link: https://%s/verify-email?hash=%s", c.Request().URI().Host(), emailVerificationRequest.Hash)

	// respond to the user
	return nil
}
