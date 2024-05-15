package handler

import (
	"database/sql"
	"fmt"
)

func CheckoutCart(db *sql.DB, orderId int, totalPrice float64, username string) {
	querySaldo := "SELECT saldo FROM users WHERE username = ?"

	rows, err := db.Query(querySaldo, username)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var saldo float64
	for rows.Next() {
		if err = rows.Scan(&saldo); err != nil {
			fmt.Println(err)
			return
		}
	}

	if saldo < totalPrice {
		fmt.Println("You have no enough balance for this transaction, please add your funds")
		return
	}
	saldo -= totalPrice

	query := "UPDATE order_details SET is_purchased = 1 WHERE order_id = ? AND is_purchased = ?"

	_, err = db.Exec(query, orderId, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	query = "UPDATE users SET saldo = ? WHERE username = ?"

	_, err = db.Exec(query, saldo, username)

	fmt.Println("Order has been purchased")
	fmt.Printf("Your current saldo is $%.2f\n", saldo)
}
