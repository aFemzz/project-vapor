package handler

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"vapor/entity"
)

func TestGetGameById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	gameID := 1
	title := "Test Game"
	price := 10.0
	rating := 4.5

	rows := sqlmock.NewRows([]string{"title", "price", "rating"}).
		AddRow(title, price, rating)

	mock.ExpectQuery("SELECT title, price, rating FROM games").
		WithArgs(gameID).
		WillReturnRows(rows)

	expectedGame := entity.Games{GameID: gameID, Title: title, Price: price, Rating: rating}
	game, err := s.GetGameById(gameID)

	assert.NoError(t, err)
	assert.Equal(t, expectedGame, game)
}

func TestGetGameById_GameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	gameID := 1

	mock.ExpectQuery("SELECT title, price, rating FROM games").
		WithArgs(gameID).
		WillReturnError(sql.ErrNoRows)

	_, err = s.GetGameById(gameID)

	assert.EqualError(t, err, "game not found")
}

func TestUpdateGame_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	game := entity.Games{GameID: 1, Title: "Updated Game", Price: 20.0, Rating: 4.0}

	mock.ExpectExec("UPDATE games SET").
		WithArgs(game.Title, game.Price, game.Rating, game.GameID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = s.UpdateGame(game)

	assert.NoError(t, err)
}

func TestUpdateGame_ErrorExec(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	game := entity.Games{GameID: 1, Title: "Updated Game", Price: 20.0, Rating: 4.0}

	mock.ExpectExec("UPDATE games SET").
		WithArgs(game.Title, game.Price, game.Rating, game.GameID).
		WillReturnError(errors.New("database error"))

	err = s.UpdateGame(game)

	assert.EqualError(t, err, "error while update game")
}
