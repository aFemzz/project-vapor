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
	reader := bufio.NewReader(os.Stdin)

	query := "SELECT g.title, g.price, od.is_purchased FROM users u JOIN orders o ON u.user_id = o.user_id JOIN order_details od ON o.order_id = od.order_id JOIN games g ON od.game_id = g.game_id WHERE u.username = ?"

	rows, err := db.Query(query, user.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("==========")
	fmt.Println("   Cart")
	fmt.Println("==========")
	fmt.Println("Game Title       | Price     |")
	for rows.Next() {
		var title string
		var price float64
		var isPurchased bool
		if err = rows.Scan(&title, &price, &isPurchased); err != nil {
			fmt.Println(err)
		}

		if !isPurchased {
			utility.PrintSpace(title, len("Game Title       "))
			utility.PrintSpace(price, len(" Price     "))
			fmt.Println()
		}
	}

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
		fmt.Println("Delete Item in Cart")
		utility.EnterToContinue()
	case "2":
		fmt.Println("Checkout")
		utility.EnterToContinue()
	case "0":
		return
	}
}
