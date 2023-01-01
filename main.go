package main

import (
	"fmt"
	"go-ds/events"
	"time"
)

func main() {
	publisher := events.CreateEventPublisher(time.Second)
	events.CreateEventData()
	for str := range publisher.Poll() {
		fmt.Println(str)
	}
}
