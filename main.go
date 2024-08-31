package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/rafaquelhodev/go-sound-player/internal/drawer"
	"github.com/rafaquelhodev/go-sound-player/internal/options"
)

func startCron(done chan os.Signal, opts *options.Options) {
	periodMilliseconds := 60 * 1000 / *opts.Bpm / *opts.Subdivisions
	ticker := time.NewTicker(time.Duration(periodMilliseconds) * time.Millisecond)

	drw := drawer.NewDrawer(opts)

	// TODO: centralize logic
	beats := *opts.Beats + 1 + *opts.Beats*(*opts.Subdivisions-1)

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
				go func(count int) {
					drw.Draw(count)
					cmd := exec.Command("mpg123", sound)
					if err := cmd.Run(); err != nil {
						log.Fatal(err)
					}
				}(count)

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

	opts := options.ReadOptions()

	go startCron(done, opts)

	<-done // Will block here until user hits ctrl+c
}
