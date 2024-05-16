package handler

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTopSellingPublisher_EmptyResult(t *testing.T) {
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
	publishers, err := handler.TopSellingPublisher()

	fmt.Println(err)
	// Check if there were any errors
	assert.Error(t, err)
	assert.Empty(t, publishers)
	assert.EqualError(t, err, "data game is empty")

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestTopSellingPublisher_NotEmptyResult(t *testing.T) {
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

	rows := sqlmock.NewRows([]string{"publisher", "cnt"}).
		AddRow("EpicGames", 2).
		AddRow("PuzzlePublishers", 2).
		AddRow("RacePublishers", 2).
		AddRow("CD Projekt Red", 1).
		AddRow("ketua rt", 1)

	// Expect the query and define its result
	mock.ExpectQuery("^SELECT.*").WillReturnRows(rows)

	// Call the TopSellingPublisher method
	publishers, err := handler.TopSellingPublisher()

	// Check if there were any errors
	assert.NoError(t, err)

	// Check if the result is as expected
	assert.Equal(t, len(publishers), 5)
	assert.Equal(t, publishers[0].Name, "EpicGames")
	assert.Equal(t, publishers[0].TotalBuy, 2)
	assert.Equal(t, publishers[1].Name, "PuzzlePublishers")
	assert.Equal(t, publishers[1].TotalBuy, 2)
	assert.Equal(t, publishers[2].Name, "RacePublishers")
	assert.Equal(t, publishers[2].TotalBuy, 2)
	assert.Equal(t, publishers[3].Name, "CD Projekt Red")
	assert.Equal(t, publishers[3].TotalBuy, 1)
	assert.Equal(t, publishers[4].Name, "ketua rt")
	assert.Equal(t, publishers[4].TotalBuy, 1)

	// Verify that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
