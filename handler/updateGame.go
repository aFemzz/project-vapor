package handler

import (
	"database/sql"
	"fmt"
)

func (s *Handler) GetGameById(gameId int) (string, float64, float64, error) {
	var title string
	var price float64
	var rating float64

	query := "SELECT title, price, rating FROM games WHERE game_id=?"

	err := s.DB.QueryRow(query, gameId).Scan(&title, &price, &rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", 0, 0, fmt.Errorf("game not found")
		} else {
			return "", 0, 0, fmt.Errorf("error while get data")
		}
	}

	return title, price, rating, nil
}

func (s *Handler) UpdateGame(title string, price float64, rating float64, gameId int) error {
	query := "UPDATE games SET title=?, price=?, rating=? WHERE game_id=?"
	_, err := s.DB.Exec(query, title, price, rating, gameId)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while update game")
	}
	
	return nil
}
