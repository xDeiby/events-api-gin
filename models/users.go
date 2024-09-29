package models

import (
	"errors"

	"example.com/gin-project/db"
	"example.com/gin-project/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password)
		VALUES (?, ?)
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	passwordHash, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := statement.Exec(u.Email, passwordHash)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT id, password FROM users
		WHERE email = ?
	`

	row := db.DB.QueryRow(query, u.Email)

	var passwordDb string
	err := row.Scan(&u.ID, &passwordDb)

	if err != nil {
		return err
	}

	isValidPassword := utils.ComparePasswords(u.Password, passwordDb)

	if !isValidPassword {
		return errors.New("invalid password")
	}

	return nil
} 