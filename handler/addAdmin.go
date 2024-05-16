package handler

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	DB *sql.DB
}

func (s *Handler) AddAdmin(username, password, email string) error {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error while encrypting password")
	}

	_, err = s.DB.Exec("INSERT INTO users (username, password, email, role, saldo) VALUES(?,?,?,?,?)", username, hashedPass, email, "admin", 0.0)
	if err != nil {
		return fmt.Errorf("error when registering user: %v", err)
	}

	fmt.Printf("User '%v' added succesfuly!\n", username)
	return nil
}
