package main

import (
	"fmt"
	"go-ds/events"
	"time"
)

func main() {
	publisher := events.CreateEventPublisher(time.Second)
	for str := range publisher.Poll() {
		fmt.Println(str)
	}
}
