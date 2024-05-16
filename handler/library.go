package handler

import (
	"fmt"
	"vapor/config"
	"vapor/entity"
	"vapor/utility"
)

func Library(user entity.User) {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("=================================================")
	fmt.Println("                    LIBRARY")
	fmt.Println("=================================================")

	query := "SELECT g.title FROM games g JOIN order_details od ON g.game_id = od.game_id JOIN orders o ON od.order_id = o.order_id WHERE o.user_id = ? AND od.is_purchased = ?"

	rows, err := db.Query(query, user.User_ID, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	index := 1
	data := false
	for rows.Next() {
		data = true
		var title string

		if err = rows.Scan(&title); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%d. %s\n", index, title)
		index += 1
	}

	if !data {
		fmt.Println("No game in your library")
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println()
	utility.EnterToContinue()
}
