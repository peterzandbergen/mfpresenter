package main

import (
	"flag"
	"log"
	"os/exec"
	"time"
)

var playercmd = flag.String("playercmd", "<empty>", "")

func main() {
	flag.Parse()

	log.Printf("playercmd=\"%s\"", *playercmd)

	// Execute the command, sleep for 10 seconds and test if it exited.
	cmd := exec.Command("sleep", "3")
	cmd.Start()
	<-time.After(time.Second * 5)
	cmd.Wait()
	if cmd.ProcessState.Exited() {
		log.Printf("Process has exited.")
	}
}
