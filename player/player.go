package player

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
)

// OMX player.
type Player struct {
	exe  string
	args []string
	cmd  *exec.Cmd // if playing then not nil
}

func splitCommandLine(cmdline string) (exe string, args []string, err error) {
	parts := strings.Split(cmdline, " ")
	if len(parts) < 1 {
		return "", nil, errors.New("Illegal player command")
	}
	exe = parts[0]
	if len(parts[1:]) > 0 {
		args = append(args, parts[1:]...)
	}
	return exe, args, nil
}

//func (p *Player) newCommand(exe string, args ...string) {

//}

// NewPlayer returns a new player or an error if the exec does not exist.
// playercmd is a string separated by spaces. It is the commmand that
// will have the name of the file to be played appended.
func New(playercmd string) (*Player, error) {
	var p = &Player{}
	// Extract the executable name and the parameters.
	if exe, args, err := splitCommandLine(playercmd); err != nil {
		return nil, err
	} else {
		p.exe = exe
		p.args = args // If nowPlaying has len() > 0 then the player is playing.
	}

	// Test if the executable exists and can execute.
	//	if fi, err := os.Stat(c.Path); err != nil || fi.IsDir() || (fi.Mode().Perm()&0100) == 0 {
	//		return nil, err
	//	}
	log.Printf("created new player from command: %s\n%s %s", playercmd, p.exe, p.args)
	return p, nil
}

// Start starts the player with the given filename. If the player is already
// running, then stop it first.
func (p *Player) Start(filename string) error {
	// Stop if we're currently playing a file.
	if p.cmd != nil {
		p.stop()
	}
	// Test if the filename is given.
	if len(filename) == 0 {
		return errors.New("No filename given.")
	}
	// Check if the file exists and is readable.
	fi, err := os.Stat(filename)
	if err != nil || (fi.Mode().Perm()&0400 == 0) {
		return errors.New("Player: file to be played not found.")
	}
	return p.start(filename)
}

func (p *Player) Stop() error {
	if p.cmd == nil {
		return errors.New("Was not playing.")
	}
	return p.stop()
}

// stop the running player.
func (p *Player) stop() error {
	p.cmd.Process.Kill()
	p.cmd.Wait()
	// Remove the file argument.
	p.cmd = nil
	return nil
}

// start the player with the given file.
func (p *Player) start(filename string) error {
	// Create the command.
	p.cmd = exec.Command(p.exe, append(p.args, filename)...)
	err := p.cmd.Start()
	if err != nil {
		p.cmd = nil
		return err
	}
	return nil
}
