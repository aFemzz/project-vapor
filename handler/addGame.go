package handler

import (
	"fmt"
	"vapor/entity"
)

func (s *Handler) AddGame(g entity.Games) error {

	// Check if the unique column value already exists
	var exists bool
	err := s.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM games WHERE title = ?)", g.Title).Scan(&exists)
	if err != nil {
		return err
	}

	// If the unique column value exists, handle it
	if exists {
		return fmt.Errorf("game with title '%s' already exists", g.Title)
	}
	query := "INSERT INTO games (title, description, price, developer, publisher, rating, is_deleted) VALUES(?,?,?,?,?,?,?)"
	_, err = s.DB.Exec(query, g.Title, g.Description, g.Price, g.Developer, g.Publisher, g.Rating, false)
	if err != nil {
		return fmt.Errorf("error when registering new game: %v", err)
	}

	fmt.Println("New game added succesfuly!")
	return nil
}
