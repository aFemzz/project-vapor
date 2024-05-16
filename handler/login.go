package handler

import (
	"database/sql"
	"fmt"
	"vapor/entity"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Login(email, password string) (entity.User, error) {
	var u entity.User

	err := h.DB.QueryRow("SELECT user_id, username, password, role, saldo FROM users WHERE email = ?", email).Scan(&u.User_ID, &u.Username, &u.Password, &u.Role, &u.Saldo)
	switch {
	case err == sql.ErrNoRows:
		// user does not exist
		return u, fmt.Errorf("password or user does not match")
	case err != nil:
		return u, fmt.Errorf("error: %v", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		// password doesn't match
		return u, fmt.Errorf("password or user doesn't match")
	}

	u.Email = email

	return u, nil
}
