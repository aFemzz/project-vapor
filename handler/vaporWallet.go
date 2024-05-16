package handler

import (
	"database/sql"
	"fmt"
	"vapor/entity"
)

func (s *Handler) VaporWallet(user entity.User) (error, float64) {

	QueryToGetWallet :=
		`
	SELECT saldo FROM users WHERE user_id = ? 
	`
	var wallet float64
	err := s.DB.QueryRow(QueryToGetWallet, user.User_ID).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found "), 0
		}
		fmt.Printf("error retrieving wallet balance: %v", err)
		return err, 0
	}

	return nil, wallet
}
