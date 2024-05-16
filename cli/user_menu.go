package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/entity"
	"vapor/handler"
	"vapor/utility"
)

func UserMenu(user entity.User, hd *handler.Handler) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Welcome to Vapor, %v\n", user.Username)
		fmt.Println()
		fmt.Println("Choose menu:")
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
			handler.PurchaseGame(user)
		case "2":
			fmt.Println("=================================================")
			fmt.Println("                      CART")
			fmt.Println("=================================================")
			fmt.Println("GAME TITLE           | PRICE             |")

			data, totalPrice, orderId, err := hd.Cart(user)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			if len(data) == 0 {
				fmt.Println("No item in your carts")
				fmt.Println()
				utility.EnterToContinue()
				break
			}

			for _, item := range data {
				utility.PrintSpace(item.Title, len("GAME TITLE           "))
				utility.PrintSpace(item.Price, len(" PRICE             "))
				fmt.Println()
			}
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			fmt.Println()
			fmt.Println("Choose menu:")
			fmt.Println("1. Delete Item in Cart")
			fmt.Println("2. Checkout")
			fmt.Println("0. Main Menu")
			fmt.Println()

			fmt.Print("Input the number: ")
			inputCartMenu, _ := reader.ReadString('\n')
			inputCartMenu = strings.TrimSpace(inputCartMenu)

			switch inputCartMenu {
			case "1":
				fmt.Print("Please enter the game title you want to delete:")
				inputDeleteItem, _ := reader.ReadString('\n')
				inputDeleteItem = strings.TrimSpace(inputDeleteItem)
				err = hd.DeleteItemInCart(orderId, inputDeleteItem)
				fmt.Println("Item deleted successfully")
				utility.EnterToContinue()
			case "2":
				saldo, err := hd.CheckoutCart(orderId, totalPrice, user.Username)
				if err != nil {
					fmt.Println("Error:", err)
					break
				}
				fmt.Println("Order has been purchased")
				fmt.Printf("Your current saldo is $%.2f\n", saldo)
				utility.EnterToContinue()
			case "0":
				break
			default:
				fmt.Println("Invalid input")
				break
			}
		case "3":
			fmt.Println("=================================================")
			fmt.Println("                    LIBRARY")
			fmt.Println("=================================================")
			data, err := hd.Library(user)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			if len(data) == 0 {
				fmt.Println("No game in your library")
			} else {
				for index, title := range data {
					fmt.Printf("%d. %s\n", index+1, title)
				}
			}
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
			fmt.Println()
			utility.EnterToContinue()
		case "4":
			handler.TopSellingGame()
		case "5":
			err, wallet := hd.VaporWallet(user)

			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("====================")
			fmt.Printf("Your Balance : $%.2f\n", wallet)
			fmt.Println("====================")

		case "6":
			handler.AddFunds(user)
		case "7":
			fmt.Println("Logged out ... ")
			return
		case "0":
			fmt.Println("Exit program ... ")
			os.Exit(1)
		}
	}
}
