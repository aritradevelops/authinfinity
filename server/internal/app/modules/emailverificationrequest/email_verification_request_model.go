package emailverificationrequest

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// implements Schema
type EmailVerificationRequest struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid()"`
	// add your additional fields here
	Hash      string    `json:"hash" validate:"required,min=3"`
	UserID    uuid.UUID `json:"user_id"`
	ExpiredAt time.Time `json:"expired_at"`
	// system generated fields
	AccountID uuid.UUID  `json:"account_id" validate:"required" gorm:"type:uuid;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"type:uuid"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"type:uuid"`
}

func Model() core.Model {
	return core.NewModel("email_verification_requests", []string{"name"})
}

func (u *EmailVerificationRequest) GetID() string {
	return u.ID.String()
}

func (u *EmailVerificationRequest) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *EmailVerificationRequest) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *EmailVerificationRequest) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *EmailVerificationRequest) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *EmailVerificationRequest) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *EmailVerificationRequest) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *EmailVerificationRequest) UnsetDeletedBy() {
	u.DeletedBy = nil
}
func (u *EmailVerificationRequest) SetAccountID(id string) {
	u.AccountID, _ = uuid.Parse(id)
}
