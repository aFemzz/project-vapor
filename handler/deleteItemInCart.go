package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func DeleteItemInCart(db *sql.DB, orderID int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter the game title you want to delete:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")

	query := "SELECT od.order_detail_id FROM games g JOIN order_details od ON g.game_id = od.game_id WHERE g.title = ? AND od.order_id = ?"

	rows, err := db.Query(query, input, orderID)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("There is no such item in your cart")
		return
	}

	var id int
	if err = rows.Scan(&id); err != nil {
		fmt.Println(err)
	}

	query = "DELETE FROM order_details WHERE order_detail_id = ?"

	_, err = db.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Item deleted successfully")
}
