package services

import (
	"database/sql"
	"go-clean-code/contracts/request"
	"go-clean-code/repository"
)

type UserService struct {
	Db *sql.DB
}

func (s *UserService) GetUsers() (*sql.Rows, error) {
	repo := repository.UserRepository{
		Db: s.Db,
	}

	rows, err := repo.GetUsers()
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *UserService) CreateUser(request request.UserCreateRequest) error {
	repo := repository.UserRepository{
		Db: s.Db,
	}

	err := repo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}
