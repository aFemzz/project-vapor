package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"vapor/entity"
	"vapor/utility"
)

func (s *Handler) ListGame() ([]entity.Games, error) {
	var games []entity.Games

	// write your code here
	queryToDisplayAllGames := `
	SELECT game_id, title, description, price, developer, publisher, rating 
	FROM games 
	WHERE is_deleted IS NULL OR is_deleted = 0;
`

	rows, err := s.DB.Query(queryToDisplayAllGames)
	if err != nil {

		return games, err
	}
	defer rows.Close()

	for rows.Next() {
		var game entity.Games
		if err := rows.Scan(&game.GameID, &game.Title, &game.Description, &game.Price, &game.Developer, &game.Publisher, &game.Rating); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		games = append(games, game)
	}
	if err := rows.Err(); err != nil {
		// fmt.Println("Error during rows iteration:", err)
		return games, err
	}

	return games, nil
}

func (s *Handler) PurchaseGame(user entity.User) error {

	var scanOrderId int
	var scanExists bool
	queryOrderIdExists :=
		`
	SELECT EXISTS ( SELECT 1 FROM orders WHERE user_id = ? )
	`

	err := s.DB.QueryRow(queryOrderIdExists, user.User_ID).Scan(&scanExists)
	if err != nil {

		return err
	}

	if !scanExists {
		queryToCreateOrderId :=
			`
		INSERT INTO orders ( user_id ) VALUES (?)
		`
		_, err := s.DB.Exec(queryToCreateOrderId, user.User_ID)
		if err != nil {

			return err
		}
		fmt.Println()
	}
	//
	queryToScanOrderId :=
		`
		SELECT order_id FROM orders
		WHERE user_id = ?
		`

	err = s.DB.QueryRow(queryToScanOrderId, user.User_ID).Scan(&scanOrderId)
	if err != nil {

		return err
	}

	games, _ := s.ListGame()
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
		var selectedGame entity.Games
		var gameIsAda bool = false
		for idx, game := range games {
			if game.GameID == InputID {
				selectedGame = games[idx]
				gameIsAda = true
				break
			}
		}

		if !gameIsAda {
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
		err = s.DB.QueryRow(checkQuery, scanOrderId, selectedGame.GameID).Scan(&isOrdered)
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

			err = s.DB.QueryRow(checkQuery, scanOrderId, selectedGame.GameID).Scan(&isPurchased)
			if err != nil {
				fmt.Println("Error checking cart status:", err)
				continue
			}

			// fmt.Println("abis query isPurchased : ", isPurchased)
			if isPurchased {
				fmt.Println("The game is already in the library.") // return this as an error params, for the mock test to check
				continue
			}
			if !isPurchased {

				fmt.Println("The game already in the cart")
				continue
			}
		}

		// Check if the game is already in the cart (order_details with is_purchased = 0)

		// Insert into order_details
		err = s.InsertToCart(scanOrderId, selectedGame.GameID)
		if err != nil {
			fmt.Println("Error inserting query", err)
		}

		fmt.Println("The game has been added to your cart.") // return this as success params, for the mock test to check
		fmt.Println()

	}
	return nil
}

func (s *Handler) InsertToCart(scanOrderId, GameID int) error {
	insertDetailQuery :=
		`INSERT INTO order_details (order_id, game_id, is_purchased, date) 
	VALUES ( ?, ?, ?, ?)`
	_, err := s.DB.Exec(insertDetailQuery, scanOrderId, GameID, false, time.Now().Format("2006-01-02"))
	if err != nil {
		return fmt.Errorf("error inserting order details: %v", err)

	}
	return nil
}
