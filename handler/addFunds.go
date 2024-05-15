package handler

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var funds float64
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

	_, err = db.Exec("UPDATE users SET saldo = saldo + ? WHERE user_id = ?", funds, u.User_ID)
	if err != nil {
		fmt.Println("Error when updating saldo:", err)
		return
	}

	fmt.Println("Add fund completed!!")
	fmt.Println()
}
