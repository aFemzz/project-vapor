package handler

import (
	"fmt"
)

func (s *Handler) DeleteItemInCart(orderID int, gameTitle string) error {

	query := "SELECT od.order_detail_id FROM games g JOIN order_details od ON g.game_id = od.game_id WHERE g.title = ? AND od.order_id = ?"

	rows, err := s.DB.Query(query, gameTitle, orderID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while get data")
	}
	defer rows.Close()

	if !rows.Next() {
		return fmt.Errorf("there is no such item in your cart")
	}

	var id int
	if err = rows.Scan(&id); err != nil {
		fmt.Println(err)
	}

	query = "DELETE FROM order_details WHERE order_detail_id = ?"

	_, err = s.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error while delete data")
	}

	return nil
}
