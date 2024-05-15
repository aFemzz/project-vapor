package handler

import (
	"database/sql"
	"fmt"
)

func CheckoutCart(db *sql.DB, orderId int) {
	query := "UPDATE order_details SET is_purchased = 1 WHERE order_id = ? AND is_purchased = ?"

	_, err := db.Exec(query, orderId, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Order has been purchased")
}
