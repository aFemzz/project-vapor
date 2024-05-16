package handler

import (
	"fmt"
)

func (s *Handler) CheckoutCart(orderId int, totalPrice float64, username string) (float64, error) {
	querySaldo := "SELECT saldo FROM users WHERE username = ?"

	rows, err := s.DB.Query(querySaldo, username)
	if err != nil {
		return 0, fmt.Errorf("error while get data")
	}
	defer rows.Close()

	var saldo float64
	for rows.Next() {
		if err = rows.Scan(&saldo); err != nil {
			return 0, fmt.Errorf("error while scan data")
		}
	}

	if saldo < totalPrice {
		return 0, fmt.Errorf("you have no enough balance for this transaction, please add your funds")
	}

	saldo -= totalPrice

	query := "UPDATE order_details SET is_purchased = 1 WHERE order_id = ? AND is_purchased = ?"

	_, err = s.DB.Exec(query, orderId, 0)
	if err != nil {
		return 0, fmt.Errorf("error while update data")
	}

	query = "UPDATE users SET saldo = ? WHERE username = ?"

	_, err = s.DB.Exec(query, saldo, username)
	if err != nil {
		return 0, fmt.Errorf("error while update data")
	}

	return saldo, nil
}
