package main

import (
	"net/http"
	"log"
	"CarStatus/handlers"
	"fmt"
	"time"
)

func main() {

	// Start Job
	go runEveryMinute(time.Minute,  handlers.PingCarAndSaveStatus)

	// Routes
	http.HandleFunc("/car/overview", handlers.GetAllCarDetails)
	http.HandleFunc("/car", handlers.SaveCarStatus)

	// Start server
	fmt.Printf("Starting server ...\n")
	if err := http.ListenAndServe(":8321", nil); err != nil {
		log.Fatal(err)
	}
}

func runEveryMinute(d time.Duration, f func(time.Time)) {
	fmt.Printf("Starting job to ping every minute ...\n")
	ticker := time.Tick(d)
	for {
		select {
		case x := <- ticker:
			f(x)
		}
	}
}