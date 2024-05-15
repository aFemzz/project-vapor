package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/config"
)

func ReportOrder() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input range of date report you want")
	fmt.Print("Input Start Date (YYYY-MM-DD):")
	inputStart, _ := reader.ReadString('\n')
	inputStart = strings.TrimSpace(inputStart)
	if inputStart == "" {
		fmt.Println("Please input start date")
		return
	}

	fmt.Print("Input End Date (YYYY-MM-DD):")
	inputEnd, _ := reader.ReadString('\n')
	inputEnd = strings.TrimSpace(inputEnd)
	if inputEnd == "" {
		fmt.Println("Please input end date")
		return
	}

	query := "SELECT g.title, u.username, od.date, g.price FROM users u JOIN orders o ON u.user_id = o.user_id JOIN order_details od ON o.order_id = od.order_id JOIN games g ON od.game_id = g.game_id WHERE od.is_purchased = ? AND od.date BETWEEN ? AND ?"

	rows, err := db.Query(query, true, inputStart, inputEnd)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var totalRevenue float64
	for rows.Next() {
		var title, username, date string
		var price float64
		if err = rows.Scan(&title, &username, &date, &price); err != nil {
			fmt.Println(err)
			return
		}
		totalRevenue += price
		fmt.Println(title, username, date)
	}
	fmt.Printf("%.2f\n", totalRevenue)
}
