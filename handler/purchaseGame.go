package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vapor/config"
	"vapor/entity"
	"vapor/utility"
)

func PurchaseGame(user entity.User) {
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
	// var GameID int
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

	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Input the Game ID to add it to cart  ")
		fmt.Print("(type done to return to menu) :")
		InputIDStr, _ := reader.ReadString('\n')
		InputIDStr = strings.TrimSpace(InputIDStr)
		_, err := strconv.Atoi(InputIDStr)

		if InputIDStr == "done" {
			fmt.Println("Returning to menu...")
			break // Exit the loop if "done" is entered
		}
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		// fmt.Printf("%s\n", Title)

		fmt.Println(user.User_ID)

		QueryToOrders, err := db.Prepare("INSERT INTO orders (user_id) VALUES (?)")
		if err != nil {
			fmt.Println("Error preparing statement:", err)
			return
		}
		defer QueryToOrders.Close()

		_, err = QueryToOrders.Exec(user.User_ID)
		if err != nil {
			fmt.Println("Error executing statement:", err)
			return
		}

	}

}
