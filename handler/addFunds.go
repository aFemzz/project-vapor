package handler

import (
	"fmt"
	"vapor/entity"
)

func (h *Handler) AddFunds(u entity.User, funds float64) error {

	_, err := h.DB.Exec("UPDATE users SET saldo = saldo + ? WHERE user_id = ?", funds, u.User_ID)
	if err != nil {
		return fmt.Errorf("error when updating funds:%v", err)

	}

	fmt.Println("Add fund completed!!")
	fmt.Println()
	return nil
}
