package user

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/google/uuid"
)

// User implements Schema
type User struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string     `json:"name" validate:"required,min=3,max=100" gorm:"type:varchar(100);not null"`
	Email     string     `json:"email" validate:"required,email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Dp        *string    `json:"dp,omitempty" validate:"omitempty,url" gorm:"type:text"`
	AccountID uuid.UUID  `json:"account_id" validate:"required" gorm:"type:uuid;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"type:uuid;not null"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTim:false"`
	UpdatedBy *uuid.UUID `json:"updated_by,omitempty" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	DeletedBy *uuid.UUID `json:"deleted_by,omitempty" gorm:"type:uuid"`
}

func Model() core.Model {
	return core.NewModel("users", []string{"name"})
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
