package handler

import (
	"fmt"
	"vapor/config"
	"vapor/utility"
)

func TopSellingPublisher() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Printf("error when connecting to db:%v\n", err)
		return
	}
	defer db.Close()

	query := `
	SELECT
		g.publisher,
		COUNT(g.publisher) as cnt
	FROM
		games g
		JOIN order_details od ON g.game_id = od.game_id
	GROUP BY
		g.publisher
	ORDER BY
		COUNT(g.game_id) DESC,
		g.publisher ASC
	LIMIT 5
	`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	fmt.Println("=================================================")
	fmt.Println("                 TOP 5 GAME PUBLISHER")
	fmt.Println("=================================================")
	fmt.Println("GAME PUBLISHER           | TOTAL BUY             |")
	for rows.Next() {
		var publisher string
		var totalBuy int
		if err := rows.Scan(&publisher, &totalBuy); err != nil {
			fmt.Println(err.Error())
			return
		}

		utility.PrintSpace(publisher, len("Game PUBLISHER           "))
		utility.PrintSpace(totalBuy, len(" Total Buy             "))
		fmt.Println()
	}
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
}
