package account

import (
	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
)

type AccountService struct {
	core.Service[*Account]
}

func Service() *AccountService {
	var accountRepository = Repository()
	return &AccountService{
		Service: core.NewService(core.Repository[*Account](accountRepository)),
	}
}

func (as *AccountService) Create(c *fiber.Ctx, data *Account) (string, error) {
	data.Slug = slug.Make(data.Name)
	return as.Service.Create(c, data)
}
