package handler

import (
	"fmt"
	"vapor/config"
	"vapor/utility"
)

func PurchaseGame() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error getting DB:", err)
		return
	}
	defer db.Close()

	queryToDisplayAllGames := `
        SELECT game_id, title, description, price, developer, publisher, rating 
        FROM games 
        WHERE is_deleted IS NULL OR is_deleted = 0;
    `

	rows, err := db.Query(queryToDisplayAllGames)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	fmt.Println("==========")
	fmt.Println("Games List")
	fmt.Println("==========")

	fmt.Println("Game ID	  | Title	 | Game Description    				 | Price     	| Publisher     		| Rating     |")

	for rows.Next() {
		var GameID int
		var Title string
		var Description string
		var Price float64
		var Developer string
		var Publisher string
		var Rating float64

		if err = rows.Scan(&GameID, &Title, &Description, &Price, &Developer, &Publisher, &Rating); err != nil {
			fmt.Println("Error scanning row:", err)

		}

		utility.PrintSpace(GameID, len("Game ID		"))
		utility.PrintSpace(Title, len("Title       		"))
		utility.PrintSpace(Description, len("Game Description       "))
		utility.PrintSpace(Price, len("Price       "))
		utility.PrintSpace(Developer, len("Developer       "))
		utility.PrintSpace(Publisher, len("Publisher       "))
		utility.PrintSpace(Rating, len("Rating       "))
		fmt.Println()

	}
	// Check for errors after the loop
	if err = rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return
	}

	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Print("Input the Game ID to add it to cart : ")
	// 	input, _ := reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	switch input {
	// 	case "1":

	// 	case "2":

	// 	case "3":

	// 	case "4":

	// 	case "7":
	// 		fmt.Println("Logged out ... ")
	// 		return

	// 	}
	// }

}
