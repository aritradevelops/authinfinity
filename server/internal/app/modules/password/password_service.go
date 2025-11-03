package password 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type PasswordService struct {
	core.Service[*Password]
}

var passwordService *PasswordService

func Service() *PasswordService {
	return passwordService
}
