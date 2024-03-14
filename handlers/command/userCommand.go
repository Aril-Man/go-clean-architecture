package command

import (
	"database/sql"
	"go-clean-code/contracts/request"
	"go-clean-code/services"
)

type UserCommand interface {
	CreateUser(request request.UserCreateRequest) error
}

type UserCommandImpl struct {
	Db *sql.DB
}

func NewUserCommand(db *sql.DB) (*UserCommandImpl, error) {
	return &UserCommandImpl{
		Db: db,
	}, nil
}

func (c *UserCommandImpl) CreateUser(request request.UserCreateRequest) error {
	service := services.UserService{
		Db: c.Db,
	}

	err := service.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}
