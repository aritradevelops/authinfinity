package user

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// implements Schema
type User struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid()"`
	// add your additional fields here
	Name      string     `json:"name" validate:"required,min=3"`
	// system generated fields
	AccountID uuid.UUID  `json:"account_id" validate:"required" gorm:"type:uuid;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"type:uuid"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"type:uuid"`
}

var userModel core.Model


func Model() core.Model {
	return userModel 
}

func (u *User) GetID() string {
	return u.ID.String()
}

func (u *User) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *User) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *User) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *User) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *User) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *User) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *User) UnsetDeletedBy() {
	u.DeletedBy = nil
}
func (u *User) SetAccountID(id string) {
	u.AccountID, _ = uuid.Parse(id)
}
