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

func DeleteGame() {
	reader := bufio.NewReader(os.Stdin)
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Print("Input Game Id to delete:")
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

	var is_deleted bool

	query := "SELECT is_deleted FROM games WHERE game_id=?"

	err = db.QueryRow(query, gameId).Scan(&is_deleted)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Game not found")
		} else {
			fmt.Println(err)
		}
		return
	}

	if is_deleted {
		fmt.Println("Game already deleted")
		return
	}
	is_deleted = true

	query = "UPDATE games SET is_deleted=? WHERE game_id=?"
	_, err = db.Exec(query, is_deleted, gameId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Game successfully deleted")
}
