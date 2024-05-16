package handler

import (
	"fmt"
	"vapor/entity"
)

func (s *Handler) Library(user entity.User) ([]string, bool, error) {
	var data []string
	isNotEmpty := false

	query := "SELECT g.title FROM games g JOIN order_details od ON g.game_id = od.game_id JOIN orders o ON od.order_id = o.order_id WHERE o.user_id = ? AND od.is_purchased = ?"

	rows, err := s.DB.Query(query, user.User_ID, true)
	if err != nil {
		return nil, false, fmt.Errorf("error while get data")
	}
	defer rows.Close()

	for rows.Next() {
		isNotEmpty = true
		var title string

		if err = rows.Scan(&title); err != nil {
			return nil, isNotEmpty, fmt.Errorf("error while scan data")
		}
		data = append(data, title)
	}

	return data, isNotEmpty, nil
}
