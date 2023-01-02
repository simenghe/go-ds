package events

import (
	"time"

	"github.com/google/uuid"
	"github.com/goombaio/namegenerator"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type EventData struct {
	Database *sqlx.DB
	Rate     time.Duration
}

type EventPublisher struct {
	Rate     time.Duration
	Messages chan string
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	UUID string `db:"uuid"`
}

func (u *User) Edit(name string, uuid string) {
	u.Name = name
	u.UUID = uuid
}

func CreateEventPublisher(rate time.Duration) *EventPublisher {
	return &EventPublisher{Rate: rate, Messages: make(chan string)}
}

func CreateEventData(rate time.Duration) *EventData {
	db, err := sqlx.Connect("sqlite3", "database.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		uuid TEXT NOT NULL);
	`)

	if err != nil {
		panic(err)
	}

	return &EventData{Database: db, Rate: rate}
}

func (ed *EventData) BuilDatabaseMock() {
	tx := ed.Database.MustBegin()
	gen := namegenerator.NewNameGenerator(time.Now().Unix())
	for i := 0; i < 10000000; i++ {
		tx.MustExec("INSERT INTO users (name, uuid) VALUES ($1, $2)", gen.Generate(), uuid.NewString())
	}
	tx.Commit()
}

func (ed *EventData) StreamRows() <-chan User {
	users := make(chan User)
	rows, err := ed.Database.Queryx("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	go func() {
		for rows.Next() {
			var curUser User
			err = rows.StructScan(&curUser)
			users <- curUser
			time.Sleep(ed.Rate)
		}

	}()
	if err != nil {
		panic(err)
	}
	return users
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
