package models

import (
	"errors"

	"example.com/events/db"
	"example.com/events/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}

func (u *User) ValidLogin() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retriverPasswrod string
	err := row.Scan(&u.ID, &retriverPasswrod)
	if err != nil {
		return errors.New("cerdentials invalid")
	}

	passwordIsVaild := utils.ComparePassword(u.Password, retriverPasswrod)
	if !passwordIsVaild {
		return errors.New("cerdentials invalid")
	}
	return nil
}
