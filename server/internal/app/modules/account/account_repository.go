package account 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AccountRepository struct {
	core.Repository[*Account]
}

var accountRepository *AccountRepository 

func Repository() *AccountRepository {
	return accountRepository 
}
	
