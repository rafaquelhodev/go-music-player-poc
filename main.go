package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func startCron(done chan os.Signal, bpm int) {
	ticker := time.NewTicker(time.Duration(int(math.Round(float64(60*1000/bpm)))) * time.Millisecond)

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
				go func() {
					cmd := exec.Command("mpg123", sound)
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}()

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

	bpm := flag.Int("bpm", 60, "the BPM value")

	flag.Parse()

	go startCron(done, *bpm)

	fmt.Println("Blocking, press ctrl+c to continue...")
	<-done // Will block here until user hits ctrl+c
}
