package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/entity"
	"vapor/handler"
)

func AdminMenu(admin entity.User) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Welcome Admin,", admin.Username)
		fmt.Println()
		fmt.Println("Choose menu:")
		fmt.Println("1. Add New Game")
		fmt.Println("2. Update Game")
		fmt.Println("3. Delete Game")
		fmt.Println("4. Report Order")
		fmt.Println("5. User Report")
		fmt.Println("6. Top Selling Publisher")
		fmt.Println("7. Add Admin")
		fmt.Println("8. Log out as Admin")
		fmt.Println("0. Exit program")
		fmt.Println()

		fmt.Print("Input the number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			handler.AddGame()
		case "2":
			handler.UpdateGame()
		case "3":
			handler.DeleteGame()
		case "4":
			handler.ReportOrder()
		case "5":
			handler.UserReport()
		case "6":
			handler.TopSellingPublisher()
		case "7":
			handler.AddAdmin()
		case "8":
			fmt.Println("Logged out as admin ... ")
			return
		case "0":
			fmt.Println("Exit program ... ")
			os.Exit(1)
		}
	}
}
