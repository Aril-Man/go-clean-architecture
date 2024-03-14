package command

import (
	"database/sql"
	"go-clean-code/contracts/request"
	"go-clean-code/services"
)

type UserCommand interface {
	CreateUser(request request.UserRequest) error
	UpdateUser(userId int, request request.UserRequest) error
}

type UserCommandImpl struct {
	Db *sql.DB
}

func NewUserCommand(db *sql.DB) (*UserCommandImpl, error) {
	return &UserCommandImpl{
		Db: db,
	}, nil
}

func (c *UserCommandImpl) CreateUser(request request.UserRequest) error {
	service := services.UserService{
		Db: c.Db,
	}

	err := service.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (c *UserCommandImpl) UpdateUser(userId int, request request.UserRequest) error {
	service := services.UserService{
		Db: c.Db,
	}

	err := service.UpdateUser(userId, request)
	if err != nil {
		return err
	}

	return nil
}
