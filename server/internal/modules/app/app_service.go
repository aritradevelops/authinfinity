package app

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/aritradevelops/authinfinity/server/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AppService struct {
	core.Service[*App]
}

func Service() *AppService {
	var appRepository = Repository()
	return &AppService{
		Service: core.NewService(core.Repository[*App](appRepository)),
	}
}

func (as *AppService) Create(c *fiber.Ctx, data *App) (string, error) {
	data.ClientID = uuid.New().String()
	secretBytes := make([]byte, 64)
	_, err := rand.Read(secretBytes)
	if err != nil {
		return "", err
	}
	data.ClientSecret = hex.EncodeToString(secretBytes)
	return as.Service.Create(c, data)
}
