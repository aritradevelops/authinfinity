package oauth

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// implements Schema
type Oauth struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID    uuid.UUID  `json:"user_id" gorm:"type:uuid" validate:"required"`
	AppID     uuid.UUID  `json:"app_id" gorm:"type:uuid" validate:"required"`
	AccountID uuid.UUID  `json:"account_id" gorm:"type:uuid" validate:"required"`
	Code      string     `json:"code" validate:"required"`
	ExpiresAt time.Time  `json:"expires_at" validate:"required"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"type:uuid"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"type:uuid"`
}

func Model() core.Model {
	return core.NewModel("oauths", []string{"name"})
}

func (u *Oauth) GetID() string {
	return u.ID.String()
}

func (u *Oauth) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *Oauth) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *Oauth) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *Oauth) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *Oauth) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *Oauth) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *Oauth) UnsetDeletedBy() {
	u.DeletedBy = nil
}

func (u *Oauth) SetAccountID(id string) {
	u.AccountID, _ = uuid.Parse(id)
}
