package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"vapor/config"
	"vapor/entity"
)

func AddFunds(u entity.User) {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error when connecting to db:", err)
		return
	}
	defer db.Close()

	fmt.Println("======================================")
	fmt.Println("             ADD FUNDS ")
	fmt.Println("======================================")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insert your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Insert your email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
}
