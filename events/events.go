package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type EventData struct {
	Database *sqlx.DB
}

type EventPublisher struct {
	Rate     time.Duration
	Messages chan string
}

func CreateEventPublisher(rate time.Duration) *EventPublisher {
	return &EventPublisher{Rate: rate, Messages: make(chan string)}
}

func CreateEventData() *EventData {
	db, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		`CREATE TABLE users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		uuid TEXT NOT NULL);
	`)

	if err != nil {
		panic(err)
	}

	return &EventData{Database: db}
}

func StartDatabase() {

}

func (e *EventPublisher) Poll() <-chan string {
	go func() {
		for {
			e.Messages <- uuid.NewString()
			time.Sleep(e.Rate)
		}
	}()
	return e.Messages
}
