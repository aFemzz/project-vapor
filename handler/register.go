package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/config"

	"golang.org/x/crypto/bcrypt"
)

func Register() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error when connecting to db:", err)
		return
	}
	defer db.Close()

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
	fmt.Println()
}
