package repository

import (
	"database/sql"
	"go-clean-code/contracts/request"
)

type UserRepository struct {
	Db *sql.DB
}

func (r *UserRepository) GetUsers() (*sql.Rows, error) {
	query, err := r.Db.Query(`SELECT * FROM users ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}

	return query, nil
}

func (r *UserRepository) CreateUser(request request.UserRequest) error {
	_, err := r.Db.Exec(`INSERT INTO users (name, age, gender) VALUE (?,?,?)`, request.Name, request.Age, request.Gender)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) UpdateUser(userId int, request request.UserRequest) error {
	_, err := r.Db.Exec(`UPDATE users SET name = ?, age = ?, gender = ? WHERE id = ?`, request.Name, request.Age, request.Gender, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserById(userId int) *sql.Row {
	query := r.Db.QueryRow(`SELECT * FROM users WHERE id = ?`, userId)
	return query
}
