package handler

import (
	"fmt"
	"vapor/entity"
)

func (s *Handler) TopSellingPublisher() ([]entity.Publisher, error) {
	query := `
	SELECT
		g.publisher,
		COUNT(g.publisher) as cnt
	FROM
		games g
		JOIN order_details od ON g.game_id = od.game_id
	GROUP BY
		g.publisher
	ORDER BY
		COUNT(g.game_id) DESC,
		g.publisher ASC
	LIMIT 5
	`

	var publisher []entity.Publisher

	rows, err := s.DB.Query(query)
	if err != nil {
		return publisher, err
	}
	defer rows.Close()

	for rows.Next() {
		var p entity.Publisher
		if err := rows.Scan(&p.Name, &p.TotalBuy); err != nil {
			return publisher, err
		}
		publisher = append(publisher, p)
	}
	if len(publisher) == 0 {
		return publisher, fmt.Errorf("data game is empty")
	}
	return publisher, nil
}
