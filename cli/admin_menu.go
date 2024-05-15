package main

import (
	"bufio"
	"fmt"
	"os"
	"vapor/entity"
)

func adminMenu(u entity.User) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Selamat datang ADMIN, ", u.Username)
	fmt.Println()
	fmt.Println("Silahkan pilih menu:")

}
