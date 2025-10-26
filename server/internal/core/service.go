package core

import (
	"reflect"

	"github.com/aritradevelops/authinfinity/server/internal/auth"
	"github.com/gofiber/fiber/v2"
)

type Service[S Schema] interface {
	List(*fiber.Ctx, *ListOpts) (*PaginatedResponse[S], error)
	Create(*fiber.Ctx, S) (string, error)
	Update(*fiber.Ctx, string, S) (bool, error)
	View(*fiber.Ctx, string) (S, error)
	Delete(*fiber.Ctx, string) (bool, error)
}

type BaseService[S Schema] struct {
	repository Repository[S]
}

func NewService[S Schema](repository Repository[S]) Service[S] {
	return &BaseService[S]{
		repository: repository,
	}
}

func (s *BaseService[S]) List(c *fiber.Ctx, listOpts *ListOpts) (*PaginatedResponse[S], error) {
	response, err := s.repository.List(listOpts)
	return response, err
}
func (s *BaseService[S]) Create(c *fiber.Ctx, data S) (string, error) {
	authUser, err := auth.GetAuthUser(c)
	if err == nil {
		data.SetCreatedAt()
		data.SetCreatedBy(authUser.ID)
		data.SetAccountID(authUser.AccountID)
	}
	return s.repository.Create(data)
}

func (s *BaseService[S]) Update(c *fiber.Ctx, id string, data S) (bool, error) {
	filter := make(Filter)
	filter["id"] = id
	filter["deleted_at"] = nil
	authUser, err := auth.GetAuthUser(c)
	if err == nil {
		data.SetUpdatedAt()
		data.SetUpdatedBy(authUser.ID)
	}
	return s.repository.Update(filter, data)
}
func (s *BaseService[S]) View(c *fiber.Ctx, id string) (S, error) {
	filter := make(Filter)
	filter["id"] = id
	filter["deleted_at"] = nil
	var data S
	err := s.repository.View(filter, &data)

	if err != nil {
		return data, err
	}
	return data, nil
}
func (s *BaseService[S]) Delete(c *fiber.Ctx, id string) (bool, error) {
	schema := reflect.New(reflect.TypeFor[S]().Elem())
	data := schema.Interface().(S)
	authUser, err := auth.GetAuthUser(c)
	if err == nil {
		data.SetDeletedAt()
		data.SetDeletedBy(authUser.ID)
	}
	filter := make(Filter)
	filter["id"] = id
	filter["deleted_at"] = nil
	return s.repository.Update(filter, data)
}
