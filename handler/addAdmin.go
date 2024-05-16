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

	// Check if the unique column value already exists
	var exists bool
	err = s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists)
	if err != nil {
		return err
	}

	// If the unique column value exists, handle it according to your requirements
	if exists {
		return fmt.Errorf("data with email '%s' already exists", email)
	}

	_, err = s.DB.Exec("INSERT INTO users (username, password, email, role, saldo) VALUES(?,?,?,?,?)", username, hashedPass, email, "admin", 0.0)
	if err != nil {
		return fmt.Errorf("error when registering user: %v", err)
	}

	fmt.Printf("Admin '%v' added succesfuly!\n", username)
	return nil
}
