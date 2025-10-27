package account

import (
	"time"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/google/uuid"
)

// Account implements Schema
type Account struct {
	ID             uuid.UUID  `json:"id,omitempty" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name           string     `json:"name" validate:"required,min=3,max=100" gorm:"type:varchar(100);not null"`
	Slug           string     `json:"slug" gorm:"type:varchar(50);uniqueIndex;not null"`
	Logo           string     `json:"logo,omitempty" validate:"omitempty,url" gorm:"type:text"`
	Domain         string     `json:"domain" validate:"required,hostname|fqdn" gorm:"type:varchar(255);uniqueIndex;not null"`
	DomainVerified bool       `json:"domain_verified" validate:"boolean" gorm:"default:false;not null"`
	CreatedAt      time.Time  `json:"created_at" gorm:"autoCreateTime:false"`
	CreatedBy      uuid.UUID  `json:"created_by"  gorm:"type:uuid;not null"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" gorm:"autoUpdateTime:false"`
	UpdatedBy      *uuid.UUID `json:"updated_by,omitempty" gorm:"type:uuid"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	DeletedBy      *uuid.UUID `json:"deleted_by,omitempty" gorm:"type:uuid"`
}

func Model() core.Model {
	return core.NewModel("accounts", []string{"name"})
}

func (u *Account) GetID() string {
	return u.ID.String()
}

func (u *Account) SetCreatedAt() {
	u.CreatedAt = time.Now().UTC()
}
func (u *Account) SetUpdatedAt() {
	now := time.Now().UTC()
	u.UpdatedAt = &now
}
func (u *Account) SetDeletedAt() {
	now := time.Now().UTC()
	u.DeletedAt = &now
}
func (u *Account) SetCreatedBy(id string) {
	u.CreatedBy, _ = uuid.Parse(id)
}

func (u *Account) SetUpdatedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.UpdatedBy = &uid
}
func (u *Account) SetDeletedBy(id string) {
	uid, _ := uuid.Parse(id)
	u.DeletedBy = &uid
}
func (u *Account) UnsetDeletedBy() {
	u.DeletedBy = nil
}

func (u *Account) SetAccountID(id string) {
	// not needed
}
