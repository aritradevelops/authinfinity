package app

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// App implements Schema
type App struct {
	ID                     uuid.UUID      `json:"id,omitempty" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name                   string         `json:"name" validate:"required,min=3,max=100" gorm:"type:varchar(100);not null"`
	Description            string         `json:"description,omitempty" validate:"omitempty,max=500" gorm:"type:text"`
	LandingUrl             string         `json:"landing_url" validate:"required,url" gorm:"type:text;not null"`
	Logo                   *string        `json:"logo,omitempty" validate:"omitempty,url" gorm:"type:text"`
	Branding               *string        `json:"branding,omitempty" validate:"omitempty" gorm:"type:text"`
	ClientID               string         `json:"client_id" gorm:"type:varchar(64);uniqueIndex;not null"`
	ClientSecret           string         `json:"client_secret" gorm:"type:varchar(255);not null"`
	RedirectUris           pq.StringArray `json:"redirect_uris" validate:"required,dive,url,min=1" gorm:"type:text[]"`
	JwtAlgo                string         `json:"jwt_algo" validate:"required,oneof=HS256 HS384 HS512 RS256 RS384 RS512" gorm:"type:varchar(10);not null"`
	JwtSecret              string         `json:"jwt_secret" validate:"required,min=32" gorm:"type:varchar(255);not null"`
	JwtLifetime            string         `json:"jwt_lifetime" validate:"required" gorm:"type:varchar(20);not null"`
	RefreshTokenLifetime   string         `json:"refresh_token_lifetime" validate:"required" gorm:"type:varchar(20);not null"`
	PermanentCallback      string         `json:"permanent_callback" validate:"required,url" gorm:"type:text;not null"`
	PermanentErrorCallback string         `json:"permanent_error_callback" validate:"required,url" gorm:"type:text;not null"`
	AccountID              uuid.UUID      `json:"account_id" gorm:"type:uuid;not null"`
	CreatedAt              time.Time      `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy              uuid.UUID      `json:"created_by" gorm:"type:uuid;not null"`
	UpdatedAt              *time.Time     `json:"updated_at,omitempty" gorm:"autoUpdateTime:false"`
	UpdatedBy              *uuid.UUID     `json:"updated_by,omitempty" gorm:"type:uuid"`
	DeletedAt              *time.Time     `json:"deleted_at,omitempty" gorm:"index"`
	DeletedBy              *uuid.UUID     `json:"deleted_by,omitempty" gorm:"type:uuid"`
}

func Model() core.Model {
	return core.NewModel("apps", []string{"name"})
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
