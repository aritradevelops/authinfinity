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
}

var accountService *AccountService

func Service() *AccountService {
	return accountService
}

func (as *AccountService) GetAccountFromReq(c *fiber.Ctx) (*Account, error) {
	host := string(c.Request().URI().Host())
	conf := config.Instance()

	fmt.Println("host", host, "config", conf.Env.BaseDomain)

	var account = new(Account)

	// if err != nil {
	// 	return nil, core.NewInternalServerError(c)
	// }

	if strings.HasSuffix(host, "."+conf.Env.BaseDomain) {
		slug := strings.TrimSuffix(host, "."+conf.Env.BaseDomain)
		err := Repository().View(core.Filter{"slug": slug}, &account)
		if err != nil {
			return nil, err
		}
	} else {
		err := Repository().View(core.Filter{"domain": host}, &account)
		if err != nil {
			return nil, err
		}
	}

	return account, nil
}
