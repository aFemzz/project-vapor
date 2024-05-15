package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"vapor/config"
)

func AddGame() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error when connecting to db:", err)
		return
	}
	defer db.Close()

	fmt.Println("======================================")
	fmt.Println("             ADD NEW GAME ")
	fmt.Println("======================================")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insert Game Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Insert Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Insert Game Price : ")
	price, _ := reader.ReadString('\n')
	price = strings.TrimSpace(price)
	priceF, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Invalid input!!")
		return
	}

	fmt.Print("Insert Game Developer : ")
	developer, _ := reader.ReadString('\n')
	developer = strings.TrimSpace(developer)

	fmt.Print("Insert Game Publisher : ")
	publisher, _ := reader.ReadString('\n')
	publisher = strings.TrimSpace(publisher)

	fmt.Print("Insert Game Rating : ")
	rating, _ := reader.ReadString('\n')
	rating = strings.TrimSpace(rating)
	ratingF, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		fmt.Println("Invalid input!!")
		return
	}

	_, err = db.Exec("INSERT INTO games (title, description, price, developer, publisher, rating, is_deleted) VALUES(?,?,?,?,?,?,?)", title, description, priceF, developer, publisher, ratingF, false)
	if err != nil {
		fmt.Println("Error when registering new game:", err)
		return
	}

	fmt.Println("New game added succesfuly!")
}
