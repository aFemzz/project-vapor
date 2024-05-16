package handler

import (
	"fmt"
	"vapor/entity"
)

func (s *Handler) Cart(user entity.User) ([]entity.Game, float64, int, error) {
	var games []entity.Game
	query := "SELECT od.order_id, g.title, g.price FROM users u JOIN orders o ON u.user_id = o.user_id JOIN order_details od ON o.order_id = od.order_id JOIN games g ON od.game_id = g.game_id WHERE u.username = ? AND od.is_purchased = ?"

	rows, err := s.DB.Query(query, user.Username, 0)
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
