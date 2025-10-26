package password

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// implements Schema
type Password struct {
	ID        uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4()"`
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

func Model() core.Model {
	return core.NewModel("passwords", []string{"name"})
}

func (u *Password) GetID() string {
	return u.ID.String()
}

func (u *Password) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *Password) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *Password) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *Password) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *Password) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *Password) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *Password) UnsetDeletedBy() {
	u.DeletedBy = nil
}
func (u *Password) SetAccountID(id string) {
	u.AccountID, _ = uuid.Parse(id)
}
