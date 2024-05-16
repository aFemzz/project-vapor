package handler

import (
	"fmt"
)

func (s *Handler) CheckoutCart(orderId int, totalPrice float64, userId int) (float64, error) {
	querySaldo := "SELECT saldo FROM users WHERE user_id = ?"

	rows, err := s.DB.Query(querySaldo, userId)
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

	query = "UPDATE users SET saldo = ? WHERE user_id = ?"

	_, err = s.DB.Exec(query, saldo, userId)
	if err != nil {
		return 0, fmt.Errorf("error while update data")
	}

	return saldo, nil
}
