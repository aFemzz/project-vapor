package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
			err := hd.PurchaseGame(user)

			if err != nil {
				fmt.Println(err)
				break
			}
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
			topGame, err := hd.TopSellingGame()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("=================================================")
			fmt.Println("                 TOP 5 SELLING GAME")
			fmt.Println("=================================================")
			fmt.Println("GAME TITLE               | TOTAL BUY             |")
			for _, v := range topGame {
				utility.PrintSpace(v.Name, len("Game Title               "))
				utility.PrintSpace(v.TotalBuy, len(" Total Buy             "))
				fmt.Println()
			}
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
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
			fmt.Println("======================================")
			fmt.Println("             ADD FUNDS ")
			fmt.Println("======================================")
			reader := bufio.NewReader(os.Stdin)
			var funds float64
			var err error
			for {

				fmt.Print("How much you want to top up? $ ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				funds, err = strconv.ParseFloat(input, 64)
				if err != nil || funds <= 0 {
					fmt.Println("Please input valid value")
					continue
				}
				break
			}
			err = hd.AddFunds(user, funds)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		case "7":
			fmt.Println("Logged out ... ")
			return
		case "0":
			fmt.Println("Exit program ... ")
			os.Exit(1)
		}
	}
}
