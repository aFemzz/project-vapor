package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

	var scanOrderId int
	var scanExists bool
	queryOrderIdExists :=
		`
	SELECT EXISTS ( SELECT 1 FROM orders WHERE user_id = ? )
	`

	err = db.QueryRow(queryOrderIdExists, user.User_ID).Scan(&scanExists)
	if err != nil {
		fmt.Println("Error checking cart status:", err)
		return
	}

	if !scanExists {
		queryToCreateOrderId :=
			`
		INSERT INTO orders ( user_id ) VALUES (?)
		`
		_, err := db.Exec(queryToCreateOrderId, user.User_ID)
		if err != nil {
			fmt.Println("Error checking cart status:", err)
			return
		}
		fmt.Println()
	}

	queryToScanOrderId :=
		`
		SELECT order_id FROM orders
		WHERE user_id = ?
		`

	err = db.QueryRow(queryToScanOrderId, user.User_ID).Scan(&scanOrderId)
	if err != nil {
		fmt.Println("Error checking cart status:", err)
		return
	}

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

	var games []entity.Game
	for rows.Next() {
		var game entity.Game
		if err := rows.Scan(&game.GameID, &game.Title, &game.Description, &game.Price, &game.Developer, &game.Publisher, &game.Rating); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
		return
	}

	// Display all games
	fmt.Println("==========")
	fmt.Println("Games List")
	fmt.Println("==========")
	fmt.Println("Game ID    | Title         | Game Description         | Price     | Developer          | Publisher          | Rating    |")
	for _, game := range games {
		utility.PrintSpace(game.GameID, len("Game ID    "))
		utility.PrintSpace(game.Title, len("Title         "))
		utility.PrintSpace(game.Description, len("Game Description         "))
		utility.PrintSpace(game.Price, len("Price     "))
		utility.PrintSpace(game.Developer, len("Developer          "))
		utility.PrintSpace(game.Publisher, len("Publisher          "))
		utility.PrintSpace(game.Rating, len("Rating     "))
		fmt.Println()
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input the Game ID to add it to cart")
		fmt.Print("(type done to return to menu) :")
		InputIDStr, _ := reader.ReadString('\n')
		InputIDStr = strings.TrimSpace(InputIDStr)
		if InputIDStr == "done" {
			fmt.Println("Returning to menu...")
			break // Exit the loop if "done" is entered
		}
		InputID, err := strconv.Atoi(InputIDStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		// Find the game with the specified ID
		var selectedGame *entity.Game
		for _, game := range games {
			if game.GameID == InputID {
				selectedGame = &game
				break
			}
		}

		if selectedGame == nil {
			fmt.Println("Game ID not found. Please enter a valid Game ID.")
			continue
		}

		fmt.Printf("Game ID: %d\n", selectedGame.GameID)
		fmt.Printf("Title: %s\n", selectedGame.Title)
		fmt.Println()

		var isOrdered bool
		checkQuery :=
			`SELECT EXISTS(
            SELECT 1 FROM order_details 
            WHERE order_id = ? AND game_id = ?
        )`
		err = db.QueryRow(checkQuery, scanOrderId, selectedGame.GameID).Scan(&isOrdered)
		if err != nil {
			fmt.Println("Error checking cart status:", err)
			continue
		}

		if isOrdered {
			var isPurchased bool
			// fmt.Println("sebelum query isPurchased : ", isPurchased)
			checkQuery :=
				`SELECT is_purchased FROM order_details 
			
			WHERE order_id = ? AND game_id = ?
			`

			err = db.QueryRow(checkQuery, scanOrderId, selectedGame.GameID).Scan(&isPurchased)
			if err != nil {
				fmt.Println("Error checking cart status:", err)
				continue
			}

			// fmt.Println("abis query isPurchased : ", isPurchased)
			if isPurchased {
				fmt.Println("The game is already in the library.")
				continue
			}
			if !isPurchased {

				fmt.Println("The game already in the cart")
				continue
			}
		}

		// Check if the game is already in the cart (order_details with is_purchased = 0)

		// Insert into order_details
		insertDetailQuery :=
			`INSERT INTO order_details (order_id, game_id, is_purchased, date) 
			VALUES ( ?, ?, ?, ?)`
		_, err = db.Exec(insertDetailQuery, scanOrderId, selectedGame.GameID, false, time.Now().Format("2006-01-02"))
		if err != nil {
			fmt.Println("Error inserting order details:", err)
			continue
		}

		fmt.Println("The game has been added to your cart.")
		fmt.Println()
	}
}
