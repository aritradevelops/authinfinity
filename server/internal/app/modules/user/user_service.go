package user 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type UserService struct {
	core.Service[*User]
}

var userService *UserService

func Service() *UserService {
	return userService
}
