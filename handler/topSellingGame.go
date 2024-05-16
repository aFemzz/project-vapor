package handler

import (
	"fmt"
	"vapor/entity"
)

func (h *Handler) TopSellingGame() ([]entity.TopGame, error) {
	query := `
	SELECT
		g.title,
		COUNT(g.game_id) as cnt
	FROM
		games g
		JOIN order_details od ON g.game_id = od.game_id
	GROUP BY
		g.game_id
	ORDER BY
		COUNT(g.game_id) DESC,
		g.title ASC
	LIMIT 5
	`

	listTopGame := []entity.TopGame{}
	rows, err := h.DB.Query(query)
	if err != nil {
		return listTopGame, err
	}
	defer rows.Close()

	for rows.Next() {

		var g entity.TopGame
		if err := rows.Scan(&g.Name, &g.TotalBuy); err != nil {
			return listTopGame, err
		}
		listTopGame = append(listTopGame, g)

	}
	if len(listTopGame) == 0 {
		return listTopGame, fmt.Errorf("empty list")
	}
	return listTopGame, nil
}
