package password 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type PasswordRepository struct {
	core.Repository[*Password]
}

var passwordRepository *PasswordRepository 

func Repository() *PasswordRepository {
	return passwordRepository 
}
	
