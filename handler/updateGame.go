package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vapor/config"
)

func UpdateGame() {
	reader := bufio.NewReader(os.Stdin)
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Print("Input Game Id to edit:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		fmt.Println("You have to enter a game id")
		return
	}
	gameId, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	var title string
	var price float64
	var rating float64

	query := "SELECT title, price, rating FROM games WHERE game_id=?"

	err = db.QueryRow(query, gameId).Scan(&title, &price, &rating)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Game not found")
		} else {
			fmt.Println(err)
		}
		return
	}

	fmt.Print("Input New Game Title (Press enter to skip):")
	inputTitle, _ := reader.ReadString('\n')
	newTitle := strings.TrimSpace(inputTitle)
	if newTitle == "" {
		newTitle = title
	}

	fmt.Print("Input New Game Price (Press enter to skip):")
	inputPrice, _ := reader.ReadString('\n')
	inputPrice = strings.TrimSpace(inputPrice)
	var newPrice float64
	if inputPrice == "" {
		newPrice = price
	} else {
		newPrice, err = strconv.ParseFloat(inputPrice, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Print("Input New Game Rating (Press enter to skip):")
	inputRating, _ := reader.ReadString('\n')
	inputRating = strings.TrimSpace(inputRating)
	var newRating float64
	if inputRating == "" {
		newRating = rating
	} else {
		newRating, err = strconv.ParseFloat(inputRating, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	query = "UPDATE games SET title=?, price=?, rating=? WHERE game_id=?"
	_, err = db.Exec(query, newTitle, newPrice, newRating, gameId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Game successfully updated")
}
