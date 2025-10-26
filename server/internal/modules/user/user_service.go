package user

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
)

type UserService struct {
	core.Service[*User]
}

func Service() *UserService {
	var userRepository = Repository()
	return &UserService{
		Service: core.NewService(core.Repository[*User](userRepository)),
	}
}
