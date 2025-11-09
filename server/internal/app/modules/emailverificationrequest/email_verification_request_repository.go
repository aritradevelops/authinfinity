package emailverificationrequest

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/db"
)

type EmailVerificationRequestRepository struct {
	core.Repository[*EmailVerificationRequest]
}

var emailVerificationRequestRepository *EmailVerificationRequestRepository

func Repository() *EmailVerificationRequestRepository {
	return emailVerificationRequestRepository
}

func (r *EmailVerificationRequestRepository) GetActiveEmailVerificationRequest(accId string, hash string) (*EmailVerificationRequest, error) {
	db := db.Instance().(*db.Postgres)
	emailVerificationRequest := &EmailVerificationRequest{}
	tx := db.Db().Model(&EmailVerificationRequest{})
	err := tx.Where("account_id = ?", accId).
		Where("hash = ?", hash).
		Where("expired_at > NOW()").
		Where("deleted_at IS NULL").
		First(emailVerificationRequest).
		Error
	if err != nil {
		return nil, err
	}
	return emailVerificationRequest, nil
}
