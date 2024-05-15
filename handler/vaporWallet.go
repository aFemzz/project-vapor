package handler

import (
	"database/sql"
	"fmt"
	"vapor/config"
	"vapor/entity"
)

func VaporWallet(user entity.User) {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println("Error getting DB:", err)
		return
	}
	defer db.Close()

	QueryToGetWallet :=
		`
	SELECT saldo FROM users WHERE user_id = ? 
	`
	var wallet float64
	err = db.QueryRow(QueryToGetWallet, user.User_ID).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("user not found")
			return
		}
		fmt.Printf("error retrieving wallet balance: %v", err)
		return
	}
	fmt.Println("====================")
	fmt.Println("Your Balance : ", wallet)
	fmt.Println("====================")

}
