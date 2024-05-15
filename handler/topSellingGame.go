package handler

import (
	"fmt"
	"vapor/config"
	"vapor/utility"
)

func TopSellingGame() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Printf("error when connecting to db:%v\n", err)
		return
	}
	defer db.Close()

	query := `
	SELECT
		g.title,
		COUNT(g.game_id) as cnt
	FROM
		games g
		JOIN order_details od ON g.game_id = od.game_id
	GROUP BY
		g.game_id
	ORDER BY
		COUNT(g.game_id) DESC,
		g.title ASC
	LIMIT 5
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	fmt.Println("=================================================")
	fmt.Println("                 TOP 5 SELLING GAME")
	fmt.Println("=================================================")
	fmt.Println("GAME TITLE               | TOTAL BUY             |")
	for rows.Next() {
		var title string
		var totalBuy int
		if err := rows.Scan(&title, &totalBuy); err != nil {
			fmt.Println(err.Error())
			return
		}

		utility.PrintSpace(title, len("Game Title               "))
		utility.PrintSpace(totalBuy, len(" Total Buy             "))
		fmt.Println()
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}
