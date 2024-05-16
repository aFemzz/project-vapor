package handler

import (
	"database/sql"
	"fmt"
)

func (s *Handler) DeleteGame(gameId int) error {
	var is_deleted bool

	query := "SELECT is_deleted FROM games WHERE game_id=?"

	err := s.DB.QueryRow(query, gameId).Scan(&is_deleted)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("game not found")
		} else {
			return fmt.Errorf("error while scan data")
		}
	}

	if is_deleted {
		return fmt.Errorf("game already deleted")
	}
	is_deleted = true

	query = "UPDATE games SET is_deleted=? WHERE game_id=?"
	_, err = s.DB.Exec(query, is_deleted, gameId)
	if err != nil {
		return fmt.Errorf("error while delete game")
	}

	return nil
}
