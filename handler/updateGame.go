package handler

import (
	"database/sql"
	"fmt"
	"vapor/entity"
)

func (s *Handler) GetGameById(gameId int) (entity.Games, error) {
	var game entity.Games
	game.GameID = gameId

	query := "SELECT title, price, rating FROM games WHERE game_id=?"

	err := s.DB.QueryRow(query, game.GameID).Scan(&game.Title, &game.Price, &game.Rating)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Games{}, fmt.Errorf("game not found")
		} else {
			return entity.Games{}, fmt.Errorf("error while get data")
		}
	}

	return game, nil
}

func (s *Handler) UpdateGame(game entity.Games) error {
	query := "UPDATE games SET title=?, price=?, rating=? WHERE game_id=?"
	_, err := s.DB.Exec(query, game.Title, game.Price, game.Rating, game.GameID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while update game")
	}

	return nil
}
