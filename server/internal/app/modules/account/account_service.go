package account

import (
	"fmt"
	"strings"

	"github.com/aritradevelops/authinfinity/server/internal/pkg/config"
	"github.com/aritradevelops/authinfinity/server/internal/pkg/core"
	"github.com/gofiber/fiber/v2"
)

type AccountService struct {
	core.Service[*Account]
	repository AccountRepository
}

func Service() *AccountService {
	var accountRepository = Repository()
	return &AccountService{
		Service: core.NewService(core.Repository[*Account](accountRepository)),
	}
}

func (as *AccountService) GetAccountFromReq(c *fiber.Ctx) (*Account, error) {
	host := string(c.Request().URI().Host())
	fmt.Println("host", host)

	var account = new(Account)
	conf, err := config.Load()

	if err != nil {
		return nil, core.NewInternalServerError(c)
	}

	if strings.HasSuffix(host, conf.Env.BaseDomain) {
		slug := strings.TrimSuffix(host, "."+conf.Env.BaseDomain)
		err := as.repository.View(core.Filter{"slug": slug}, &account)
		if err != nil {
			return nil, err
		}
	} else {
		err := as.repository.View(core.Filter{"domain": host}, &account)
		if err != nil {
			return nil, err
		}
	}

	return account, nil
}
