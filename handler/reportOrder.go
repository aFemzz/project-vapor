package handler

import (
	"fmt"
)

type Report struct {
	GameTitle string
	Username  string
	Date      string
}

func (s *Handler) ReportOrder(inputStart string, inputEnd string) ([]Report, float64, error) {
	var reports []Report
	query := "SELECT g.title, u.username, od.date, g.price FROM users u JOIN orders o ON u.user_id = o.user_id JOIN order_details od ON o.order_id = od.order_id JOIN games g ON od.game_id = g.game_id WHERE od.is_purchased = ? AND od.date BETWEEN ? AND ?"

	rows, err := s.DB.Query(query, true, inputStart, inputEnd)
	if err != nil {
		return []Report{}, 0, fmt.Errorf("error while get data")
	}
	defer rows.Close()

	var totalRevenue float64
	for rows.Next() {
		var title, username, date string
		var price float64
		if err = rows.Scan(&title, &username, &date, &price); err != nil {
			return []Report{}, 0, fmt.Errorf("error while scan data")
		}
		totalRevenue += price
		reports = append(reports, Report{title, username, date})
	}

	return reports, totalRevenue, nil
}
