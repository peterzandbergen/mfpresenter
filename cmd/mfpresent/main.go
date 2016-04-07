// mfpresenter project main.go
package main

import (
	// "fmt"
	"os"

	"github.com/peterzandbergen/mfpresenter/config"
)

// Run starts the monitor loop for the USB directory, default is taken from
// the configuration settings.
func Run() error {
	// TODO: Implement this function.
	// Start the player if there is a file in the cache.
	for {
		// Wait for changes in the directory.
	}
}

func main() {
	// Process the settings.
	conf := config.NewConfig()
	conf.InitConfig()

	// Check the environment.
	if err := conf.EnvValid(); err != nil {
		// Log the error and abort.
		os.Exit(1)
	}
	// Start the processes.
	if err := Run(); err != nil {
		// Log the error.
	}
}
