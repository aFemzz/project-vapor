package handler

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"vapor/config"
	"vapor/entity"

	"golang.org/x/crypto/bcrypt"
)

func Register() {
	var u entity.User
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error when connecting to db:", err)
		return
	}
	defer db.Close()

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

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while encrypting password")
		return
	}

	_, err = db.Exec("INSERT INTO users (username, password, email, role, saldo) VALUES(?,?,?,?,?)", username, hashedPass, email, "user", 0.0)
	if err != nil {
		fmt.Println("Error when registering user:", err)
		return
	}

	fmt.Printf("User '%v' added succesfuly!\n", username)
	fmt.Println("Your funds is zero, do you want to add funds? (y/n)")
	for {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "y":
			// get the user data
			err = db.QueryRow("SELECT user_id, username, password, role, saldo FROM users WHERE email = ?", email).Scan(&u.User_ID, &u.Username, &u.Password, &u.Role, &u.Saldo)
			switch {
			case err == sql.ErrNoRows:
				fmt.Println("password or user doesn't match")
			case err != nil:
				fmt.Printf("error: %v\n", err.Error())
			}
			AddFunds(u)
			return
		case "n":
			return
		default:
			fmt.Println("invalid input")
		}
	}
}
