package password 

import (
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
)

type PasswordService struct {
	core.Service[*Password]
}


func Service() *PasswordService {
	var passwordRepository = Repository()
	return &PasswordService{
		Service: core.NewService(core.Repository[*Password](passwordRepository)),
	}
}
