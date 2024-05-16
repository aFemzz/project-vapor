package handler

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteGame_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	gameID := 1

	rows := sqlmock.NewRows([]string{"is_deleted"}).
		AddRow(false) // Simulate the game is not deleted

	mock.ExpectQuery("SELECT is_deleted FROM games").
		WithArgs(gameID).
		WillReturnRows(rows)

	mock.ExpectExec("UPDATE games SET is_deleted").
		WithArgs(true, gameID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = s.DeleteGame(gameID)

	assert.NoError(t, err)
}

func TestDeleteGame_GameNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	gameID := 1

	mock.ExpectQuery("SELECT is_deleted FROM games").
		WithArgs(gameID).
		WillReturnError(sql.ErrNoRows)

	err = s.DeleteGame(gameID)

	assert.EqualError(t, err, "game not found")
}

func TestDeleteGame_GameAlreadyDeleted(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing mock database: %s", err)
	}
	defer db.Close()

	s := Handler{DB: db}

	gameID := 1

	rows := sqlmock.NewRows([]string{"is_deleted"}).
		AddRow(true) // Simulate the game is already deleted

	mock.ExpectQuery("SELECT is_deleted FROM games").
		WithArgs(gameID).
		WillReturnRows(rows)

	err = s.DeleteGame(gameID)

	assert.EqualError(t, err, "game already deleted")
}
