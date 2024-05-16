package handler

import (
	"database/sql"
	"fmt"
	"vapor/entity"
)

func (s *Handler) VaporWallet(user entity.User) (float64, error) {

	QueryToGetWallet :=
		`
	SELECT saldo FROM users WHERE user_id = ? 
	`
	var wallet float64
	err := s.DB.QueryRow(QueryToGetWallet, user.User_ID).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("user not found ")
		}
		fmt.Printf("error retrieving wallet balance: %v", err)
		return 0, err
	}

	return wallet, nil
}
