package model

import (
	"github.com/car-api/db"
	"github.com/car-api/utils"
)

type Users struct {
	Name     string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := `INSERT INTO USERS(name, email, password) VALUES(?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Name, u.Email, hashPass)

	return err
}
