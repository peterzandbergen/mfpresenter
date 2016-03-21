package main

import (
	"os"
	"os/exec"
)

type Player struct {
	nowPlaying string
	execName   string
	cmd        exec.Cmd
	proc       os.Process
}

// Start starts the player with the given filename. If the player is already
// running, then first kill it.
func (p *Player) Start(filename string) error {
	return nil
}

func (p *Player) Stop() error {
	return p.kill()
}

// kill the running player.
func (p *Player) kill() error {
	return p.proc.Kill()
}

// start the player with the given file.
func (p *Player) start() error {
	return nil
}
