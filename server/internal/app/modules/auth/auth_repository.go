package auth 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type AuthRepository struct {
	core.Repository[*Auth]
}

var authRepository *AuthRepository 

func Repository() *AuthRepository {
	return authRepository 
}
	
