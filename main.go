package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("mpg123", "./beep.mp3")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
