package main

import (
	"errors"
	"os"
	"os/exec"
)

// OMX player.
type Player struct {
	nowPlaying string
	execName   string
	cmd        exec.Cmd
	proc       os.Process
}

// NewPlayer returns a new player or an error if the exec does not exist.
func NewPlayer(ename string) (*Player, error) {
	// Test if the executable exists.
	// Check if the player exists.
	fi, err := os.Stat(ename)
	if err != nil || fi.IsDir() || (fi.Mode().Perm()&0x100) == 0 {
		// File cannot be executed.
		return nil, err
	}
	return &Player{
		execName: ename,
	}, nil
}

// Start starts the player with the given filename. If the player is already
// running, then first kill it.
func (p *Player) Start(filename string) error {
	if len(p.nowPlaying) > 0 {
		p.kill()
		p.nowPlaying = ""
	}
	p.nowPlaying = filename
	return p.start()
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
	// Check if the file exists and is readable.
	fi, err := os.Stat(p.nowPlaying)
	if err != nil || (fi.Mode().Perm()&0400 == 0) {
		return errors.New("Player: file to be played not found.")
	}
	// Build the command.

	return nil
}
