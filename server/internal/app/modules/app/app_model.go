package app

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// implements Schema
type App struct {
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

var appModel core.Model


func Model() core.Model {
	return appModel 
}

func (u *App) GetID() string {
	return u.ID.String()
}

func (u *App) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *App) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *App) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *App) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}
func (u *App) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *App) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *App) UnsetDeletedBy() {
	u.DeletedBy = nil
}
func (u *App) SetAccountID(id string) {
	u.AccountID, _ = uuid.Parse(id)
}
