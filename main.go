package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func startCron(done chan os.Signal) {
	ticker := time.NewTicker(time.Duration(1) * time.Millisecond)

	beats := 4

	count := 0

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				count += 1
				sound := "./metronome.mp3"
				if count == 1 {
					sound = "./beep.mp3"
				}
				cmd := exec.Command("mpg123", sound)
				if err := cmd.Run(); err != nil {
					log.Fatal(err)
				}
				if count == beats {
					count = 0
				}
			}
		}
	}()
}

func main() {
	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go startCron(done)

	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done // Will block here until user hits ctrl+c
}
