// mfpresenter project main.go
package main

import (
	"flag"
	"os"
	"strings"

	_ "github.com/peterzandbergen/mfpresenter/watcher"
)

const (
	// EnvCheckDir is the name of the env variables for the checked directory.
	EnvCheckDir = "MFP_CHECK_DIR"
	// EnvPlayerExec is the name of the env variable for the media player executable.
	EnvPlayerExec = "MFP_PLAYER_EXEC"
	// EnvExtensions is the name of the env variable that can be used to specify the file extensions
	// that the player supports. Must be specified using a colon as separator.
	EnvExtensions = "MFP_EXTENSIONS"

	// DefaultCheckDir is the default directory that is being monitored for new directories (not files).
	DefaultCheckDir = "/media"
	// DefaultPlayerExec is name of the default player executable.
	DefaultPlayerExec = "omxplayer"
	// DefaultExtensions contains the list of default file extensions, separated by a ":".
	DefaultExtensions = "mp4"
)

var (
	// FlagCheckDir is the command line option for the directory to be checked.
	FlagCheckDir = flag.String("check-dir", "", "Directory being checked for new directories and files beneath, default is /media/<userid>")
	// FlagPlayerExec is the command line option for the player exec name.
	FlagPlayerExec = flag.String("player-exec", "", "Player executable, default is omxplayer.")
	// FlagExtensions is the command line option for the valid file extensions.
	FlagExtensions = flag.String("default-extensions", "", "The valid file extensions separated by a colon. Default is mp4.")
)

// Config contains the configuration settings.
type Config struct {
	// The directory that is being checked for new directories.
	CheckDir            string
	PlayerExec          string
	MediaFileExtensions []string
}

var defaultConfig Config

// Initialize a default config.
func newConfig() *Config {
	return &Config{
		CheckDir:            DefaultCheckDir,
		PlayerExec:          DefaultPlayerExec,
		MediaFileExtensions: strings.Split(DefaultExtensions, ":"),
	}
}

func (c *Config) initConfigFromEnv() {
	if v, b := os.LookupEnv(EnvCheckDir); b {
		c.CheckDir = v
	}
	if v, b := os.LookupEnv(EnvPlayerExec); b {
		c.PlayerExec = v
	}
	if v, b := os.LookupEnv(EnvExtensions); b {
		exts := strings.Split(v, ":")
		if len(exts) > 0 {
			c.MediaFileExtensions = exts
		}
	}
}

func (c *Config) initConfigFromOptions() {
	flag.Parse()
	if len(*FlagCheckDir) > 0 {
		c.CheckDir = *FlagCheckDir
	}
	if len(*FlagPlayerExec) > 0 {
		c.PlayerExec = *FlagPlayerExec
	}
	if len(*FlagExtensions) > 0 {
		exts := strings.Split(*FlagExtensions, ":")
		if len(exts) > 0 {
			c.MediaFileExtensions = exts
		}
	}
}

// Run starts the monitor loop for the USB directory, default is /media.
func Run() {
	// TODO Fix this function.
}

func main() {
	// Process the settings.
	// Check the environment.
	// Start the processes.

	// Test
	// d := watcher.Dir
}
