package handler

import (
	"fmt"
	"vapor/entity"
)

func (s *Handler) Cart(user entity.User) ([]entity.Game, float64, int, error) {
	var games []entity.Game
	query := "SELECT od.order_id, g.title, g.price FROM games g JOIN order_details od ON g.game_id = od.game_id JOIN orders o ON od.order_id = o.order_id WHERE o.user_id = ? AND od.is_purchased = ?"

	rows, err := s.DB.Query(query, user.User_ID, 0)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("error while get data")
	}
	defer rows.Close()

	var orderId int
	var totalPrice float64

	for rows.Next() {
		var title string
		var price float64
		if err = rows.Scan(&orderId, &title, &price); err != nil {
			return nil, 0, 0, fmt.Errorf("error while scan data")
		}

		totalPrice += price
		games = append(games, entity.Game{Title: title, Price: price})
	}

	return games, totalPrice, orderId, nil
}
