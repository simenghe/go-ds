package main

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name         string
	Email        string
	Age          uint32
	PasswordHash string
}

type Hub struct {
	Mu      *sync.Mutex
	UserMap map[string]User
	Wg      sync.WaitGroup
}

func (u *User) hash() string {
	return u.Name + "--" + u.Email
}

func (h *Hub) AddUser(user User) error {
	h.Mu.Lock()
	defer h.Mu.Unlock()
	h.UserMap[user.hash()] = user
	return nil
}

func GenerateRandomUser() User {
	email := uuid.NewString() + "@gmail.com"
	password := uuid.NewString()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return User{
		Name:         "John" + strconv.Itoa(rand.Intn(100000)),
		Email:        email,
		Age:          uint32(rand.Intn(100)),
		PasswordHash: string(hash),
	}
}

func main() {
	now := time.Now()
	hub := Hub{&sync.Mutex{}, make(map[string]User, 0), sync.WaitGroup{}}
	userCount := 10

	for i := 0; i < userCount; i++ {
		hub.AddUser(GenerateRandomUser())
	}

	log.Println("Finished", len(hub.UserMap), time.Since(now))
}
