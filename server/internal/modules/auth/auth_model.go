package auth

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/google/uuid"
)

// implements Schema
type Auth struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string     `json:"name" validate:"required,min=3"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy uuid.UUID  `json:"created_by" gorm:"type:uuid"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"autoUpdateTime:false"`
	UpdatedBy *uuid.UUID `json:"updated_by" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"type:uuid"`
}

func Model() core.Model {
	return core.NewModel("auths", []string{"name"})
}

func (u *Auth) GetID() string {
	return u.ID.String()
}

func (u *Auth) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *Auth) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *Auth) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *Auth) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *Auth) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *Auth) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *Auth) UnsetDeletedBy() {
	u.DeletedBy = nil
}
