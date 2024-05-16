package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/config"
	"vapor/entity"
	"vapor/utility"
)

func Cart(user entity.User) {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)

	query := "SELECT od.order_id, g.title, g.price, od.is_purchased FROM users u JOIN orders o ON u.user_id = o.user_id JOIN order_details od ON o.order_id = od.order_id JOIN games g ON od.game_id = g.game_id WHERE u.username = ? AND od.is_purchased = ?"

	rows, err := db.Query(query, user.Username, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var orderId int
	var totalPrice float64
	data := false

	fmt.Println("=================================================")
	fmt.Println("                      CART")
	fmt.Println("=================================================")
	fmt.Println("GAME TITLE           | PRICE             |")

	for rows.Next() {
		data = true
		var title string
		var price float64
		var isPurchased bool
		if err = rows.Scan(&orderId, &title, &price, &isPurchased); err != nil {
			fmt.Println(err)
		}

		totalPrice += price
		if !isPurchased {
			utility.PrintSpace(title, len("GAME TITLE           "))
			utility.PrintSpace(price, len(" PRICE             "))
			fmt.Println()
		}
	}

	if !data {
		fmt.Println("No item in your carts")
		fmt.Println()
		utility.EnterToContinue()
		return
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println()
	fmt.Println("Choose menu:")
	fmt.Println("1. Delete Item in Cart")
	fmt.Println("2. Checkout")
	fmt.Println("0. Main Menu")
	fmt.Println()

	fmt.Print("Input the number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		DeleteItemInCart(db, orderId)
		utility.EnterToContinue()
	case "2":
		CheckoutCart(db, orderId, totalPrice, user.Username)
		utility.EnterToContinue()
	case "0":
		return
	}
}
