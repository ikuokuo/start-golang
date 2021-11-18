package main

import (
	"fmt"
	"time"
)

func main() {
	duration := time.Second
	timer := time.NewTimer(duration)
	defer timer.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-timer.C:
			fmt.Println("Current time: ", t)
			timer.Reset(duration)
		}
	}
}
