package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/handler"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Selamat datang di game Vapor")
		fmt.Println()
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Println()
		fmt.Print("Masukkan nomor (1/2): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			user, err := handler.Login()
			if err != nil {
				fmt.Println("Error when log in: ", err)
			}
			if user.Role == "admin" {
				// role admin
				adminMenu(user)
			} else {
				// role user
				userMenu(user)
			}
		case "2":
			handler.Register()
		case "3":
			fmt.Println("Exit program...")
		default:
			fmt.Println("Input is invalid!")
		}
	}
}
