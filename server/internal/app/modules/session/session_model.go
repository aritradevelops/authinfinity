package session

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// implements Schema
type Session struct {
	ID uuid.UUID `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid()"`
	// add your additional fields here
	IP           string    `json:"ip" validate:"required,ip"`
	UserAgent    string    `json:"user_agent" validate:"required"`
	RefreshToken string    `json:"refresh_token" validate:"required"`
	UserID       uuid.UUID `json:"user_id" gorm:"type:uuid" validate:"required"`
	AppID        uuid.UUID `json:"app_id" gorm:"type:uuid" validate:"required"`
	// system generated fields
	AccountID uuid.UUID  `json:"account_id" validate:"required" gorm:"type:uuid;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"type:uuid"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"type:uuid"`
}

var sessionModel core.Model

func Model() core.Model {
	return sessionModel
}

func (u *Session) GetID() string {
	return u.ID.String()
}

func (u *Session) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *Session) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *Session) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *Session) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *Session) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *Session) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *Session) UnsetDeletedBy() {
	u.DeletedBy = nil
}
func (u *Session) SetAccountID(id string) {
	u.AccountID, _ = uuid.Parse(id)
}
