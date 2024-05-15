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

func Login() (entity.User, error) {
	var u entity.User

	db, err := config.GetDB()
	if err != nil {
		return u, fmt.Errorf("error when connecting to db:%v", err)
	}
	defer db.Close()

	fmt.Println("======================================")
	fmt.Println("             	 LOGIN ")
	fmt.Println("======================================")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Insert your email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Insert your password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	err = db.QueryRow("SELECT user_id, username, password, role, saldo FROM users WHERE email = ?", email).Scan(&u.User_ID, &u.Username, &u.Password, &u.Role, &u.Saldo)
	switch {
	case err == sql.ErrNoRows:
		return u, fmt.Errorf("password or user doesn't match")
	case err != nil:
		return u, fmt.Errorf("error: %v", err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return u, fmt.Errorf("password or user doesn't match")
	}

	u.Email = email

	return u, nil
}
