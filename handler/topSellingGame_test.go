package handler

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTopSellingGame_EmptyResult(t *testing.T) {
	// Create a new mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new handler with the mock DB
	handler := &Handler{
		DB: mockDB,
	}

	// Expect the query to be executed and return an empty result set
	mock.ExpectQuery("^SELECT.*").WillReturnRows(sqlmock.NewRows([]string{}))

	// Call the TopSellingPublisher method
	topGame, err := handler.TopSellingGame()

	fmt.Println(err)
	// Check if there were any errors
	assert.Error(t, err)
	assert.Empty(t, topGame)
	assert.EqualError(t, err, "empty list")

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTopSellingGame_NotEmptyResult(t *testing.T) {
	// Create a new mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating mock database: %v", err)
	}
	defer mockDB.Close()

	// Create a new handler with the mock DB
	handler := &Handler{
		DB: mockDB,
	}

	rows := sqlmock.NewRows([]string{"title", "cnt"}).
		AddRow("Epic Adventure", 2).
		AddRow("Puzzle Quest", 2).
		AddRow("Racing Pro", 2).
		AddRow("Cyberpunk 2077", 1).
		AddRow("The Witcher 3", 1)

	// Expect the query and define its result
	mock.ExpectQuery("^SELECT.*").WillReturnRows(rows)

	// Call the TopSellingGame method
	topGame, err := handler.TopSellingGame()

	// Check if there were any errors
	assert.NoError(t, err)

	// Check if the result is as expected
	assert.Equal(t, len(topGame), 5)
	assert.Equal(t, topGame[0].Name, "Epic Adventure")
	assert.Equal(t, topGame[0].TotalBuy, 2)
	assert.Equal(t, topGame[1].Name, "Puzzle Quest")
	assert.Equal(t, topGame[1].TotalBuy, 2)
	assert.Equal(t, topGame[2].Name, "Racing Pro")
	assert.Equal(t, topGame[2].TotalBuy, 2)
	assert.Equal(t, topGame[3].Name, "Cyberpunk 2077")
	assert.Equal(t, topGame[3].TotalBuy, 1)
	assert.Equal(t, topGame[4].Name, "The Witcher 3")
	assert.Equal(t, topGame[4].TotalBuy, 1)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
