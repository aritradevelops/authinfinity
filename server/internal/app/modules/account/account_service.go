package account 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AccountService struct {
	core.Service[*Account]
}

var accountService *AccountService

func Service() *AccountService {
	return accountService
}
