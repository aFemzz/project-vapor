package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/entity"
	"vapor/handler"
)

func UserMenu(user entity.User) {
	example := entity.User{User_ID: 1, Username: "johndoe"}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Welcome to Vapor , ", example.Username)
		fmt.Println()
		fmt.Println("Choose menu:")
		// fmt.Println("1. List Game")
		fmt.Println("1. Purchase Game")
		fmt.Println("2. Cart")
		fmt.Println("3. Library")
		fmt.Println("4. Top Selling Game")
		fmt.Println("5. Vapor Wallet")
		fmt.Println("6. Add Funds")
		fmt.Println("7. Log out")
		fmt.Println("0. Exit program")
		fmt.Println()

		fmt.Print("Input the number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			handler.PurchaseGame()
		case "2":
			handler.Cart(example)
		case "3":
			handler.Library()
		case "4":
			handler.TopSellingGame()
		case "5":
			handler.VaporWallet()
		case "6":
			handler.AddFunds()
		case "7":
			fmt.Println("Logged out ... ")
			return
		case "0":
			fmt.Println("Exit program ... ")
			os.Exit(1)
		}
	}
}
