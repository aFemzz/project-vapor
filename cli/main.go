package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"vapor/config"
	"vapor/handler"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	db, err := config.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	hd := &handler.Handler{
		DB: db,
	}

	for {
		fmt.Println("Selamat datang di Vapor")
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
				break
			}
			if user.Role == "admin" {
				// role admin
				AdminMenu(user, hd)
			} else {
				// role user
				UserMenu(user, hd)
			}
		case "2":
			fmt.Println("======================================")
			fmt.Println("         REGISTER NEW ACCOUNT ")
			fmt.Println("======================================")
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Insert your username: ")
			username, _ := reader.ReadString('\n')
			username = strings.TrimSpace(username)

			fmt.Print("Insert your email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Insert your password: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)
			err = hd.Register(username, email, password)
			if err != nil {
				fmt.Println(err)
			}
		case "0":
			fmt.Println("Exit program...")
			return
		default:
			fmt.Println("Input is invalid!")
		}
	}
}
