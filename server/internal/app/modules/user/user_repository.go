package user 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type UserRepository struct {
	core.Repository[*User]
}

var userRepository *UserRepository 

func Repository() *UserRepository {
	return userRepository 
}
	
