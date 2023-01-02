package main

import (
	"fmt"
	"go-ds/events"
	"runtime"
	"time"

	"github.com/google/uuid"
)

func main() {
	// publisher := events.CreateEventPublisher(time.Second)
	evtData := events.CreateEventData(time.Microsecond)
	evtData.StreamRows()
	// evtData.StartDatabase()
	limiter := make(chan bool, 8)
	for user := range evtData.StreamRows() {
		// fmt.Println(user)
		limiter <- true
		go func(user events.User) {
			fmt.Printf("Editing: %+v with %+v goroutines active\n", user.Name, runtime.NumGoroutine())
			time.Sleep(10 * time.Second)
			<-limiter
			user.Edit("Cucarach", uuid.NewString())
		}(user)
	}
	// for str := range publisher.Poll() {
	// 	fmt.Println(str)
	// }
}
