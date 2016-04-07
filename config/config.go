package config

import (
	"flag"
	// "fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	// EnvCheckDir is the name of the env variables for the checked directory.
	EnvCheckDir = "MFP_CHECK_DIR"
	// EnvCacheDir
	EnvCacheDir = "MFP_CACHE_DIR"
	// EnvPlayerExec is the name of the env variable for the media player executable.
	EnvPlayerExec = "MFP_PLAYER_EXEC"
	// EnvExtensions is the name of the env variable that can be used to specify the file extensions
	// that the player supports. Must be specified using a colon as separator.
	EnvExtensions = "MFP_EXTENSIONS"
	// DefaultCheckDir is the default directory that is being monitored for new directories (not files).
	DefaultCheckDir = "/media"
	// DefaultCacheDir
	DefaultCacheDir = "/var/lib/mfpresent"
	// DefaultPlayerExec is name of the default player executable.
	DefaultPlayerExec = "omxplayer"
	// DefaultExtensions contains the list of default file extensions, separated by a ":".
	DefaultExtensions = "mp4"
)

var (
	// FlagCheckDir is the command line option for the directory to be checked.
	FlagCheckDir = flag.String("check-dir", "", "Directory being checked for new directories and files beneath, default is /media/<userid>")
	// FlagCacheDir
	FlagCacheDir = flag.String("cache-dir", "", "Cache directory, defaults to /var/lib/mfpresent")
	// FlagPlayerExec is the command line option for the player exec name.
	FlagPlayerExec = flag.String("player-exec", "", "Player command, default is omxplayer.")
	// FlagExtensions is the command line option for the valid extensions.
	FlagExtensions = flag.String("default-extensions", "", "The valid file extensions separated by a colon. Default is mp4.")
)

// Config contains the configuration settings.
type Config struct {
	// The directory that is being checked for new directories.
	CheckDir            string
	CacheDir            string
	PlayerExec          string
	MediaFileExtensions []string
}

var defaultConfig Config

// Initialize a default config.
func NewConfig() *Config {
	return &Config{
		CheckDir:            DefaultCheckDir,
		CacheDir:            DefaultCacheDir,
		PlayerExec:          DefaultPlayerExec,
		MediaFileExtensions: strings.Split(DefaultExtensions, ":"),
	}
}

func (c *Config) InitConfig() {
	c.initConfigFromEnv()
	c.initConfigFromOptions()
}

func (c *Config) initConfigFromEnv() {
	if v, b := os.LookupEnv(EnvCheckDir); b {
		c.CheckDir = v
	}
	if v, b := os.LookupEnv(EnvCacheDir); b {
		c.CacheDir = v
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
	if len(*FlagCacheDir) > 0 {
		c.CacheDir = *FlagCacheDir
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

type EnvError struct {
	Err    string
	Errors []string
}

func (e *EnvError) Error() string {
	//	if e == nil {
	//		return "error was nil"
	//	}
	// Return a nice error message.
	return e.Err // + e.Errors
}

// EnvValid checks if the settings are valid (exec is present, cache dir can be
// created.
func (c *Config) EnvValid() error {
	retError := &EnvError{}

	// Check if the player exists.
	parts := strings.Split(c.PlayerExec, " ")
	epath, err := exec.LookPath(parts[0])
	if err != nil {
		return err
	}
	fi, err := os.Stat(epath)
	if err != nil || fi.IsDir() || (fi.Mode().Perm()&0500) == 0 {
		// File cannot be executed.
		retError.Err = "player not present or executable"
		retError.Errors = append(retError.Errors, retError.Err)
		return retError
	}

	// Check if the monitor directory exists.
	fi, err = os.Stat(c.CheckDir)
	if err != nil || !fi.IsDir() || (fi.Mode().Perm()&0400) == 0 {
		retError.Err = "check dir not present or not readable"
		retError.Errors = append(retError.Errors, retError.Err)
		return retError
	}
	// Check if the cache directory exists.
	fi, err = os.Stat(c.CacheDir)
	if err != nil || !fi.IsDir() || (fi.Mode().Perm()&0600) == 0 {
		retError.Err = "cache dir not present or not rw"
		retError.Errors = append(retError.Errors, retError.Err)
		return retError
	}
	if len(retError.Errors) == 0 {
		return nil
	}
	return retError
}
