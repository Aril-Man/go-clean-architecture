package query

import (
	"database/sql"
	"go-clean-code/contracts/response"
	"go-clean-code/services"
)

type UserQuery interface {
	GetUsers() ([]response.UserResponse, error)
	GetUserById(userId int) response.UserResponse
}

type UserQueryImpl struct {
	Db *sql.DB
}

func NewUserQuery(db *sql.DB) (*UserQueryImpl, error) {
	return &UserQueryImpl{
		Db: db,
	}, nil
}

func (q *UserQueryImpl) GetUsers() ([]response.UserResponse, error) {

	var usersResponse []response.UserResponse

	service := services.UserService{
		Db: q.Db,
	}

	rows, err := service.GetUsers()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var model response.UserResponse
		if err := rows.Scan(
			&model.Id,
			&model.Name,
			&model.Age,
			&model.Gender,
		); err != nil {
			return nil, err
		}

		usersResponse = append(usersResponse, model)
	}

	return usersResponse, nil
}

func (q *UserQueryImpl) GetUserById(userId int) response.UserResponse {
	var userReponse response.UserResponse

	service := services.UserService{
		Db: q.Db,
	}

	row := service.GetUserById(userId)

	if err := row.Scan(
		&userReponse.Id,
		&userReponse.Name,
		&userReponse.Age,
		&userReponse.Gender,
	); err != nil {
		return userReponse
	}

	return userReponse
}
