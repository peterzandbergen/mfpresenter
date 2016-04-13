// mfpresenter project main.go
// Test command:
//      ./mfpresent --player-exec "vlc --loop" --cache-dir /home/peza/Desktop/CacheDir --check-dir /home/peza/Desktop/CheckDir
package main

import (
	"fmt"
	"io"
	// "log"
	"os"
	"path/filepath"
	"time"

	"github.com/peterzandbergen/mfpresenter"
	"github.com/peterzandbergen/mfpresenter/config"
	"github.com/peterzandbergen/mfpresenter/player"
	"github.com/peterzandbergen/mfpresenter/scanner"

	"github.com/fsnotify/fsnotify"
)

var log = mfpresenter.Logger()

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	written, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

func scanAndCopyFile(conf *config.Config) error {
	// Try to find a file in the check dir.
	f, err := scanner.FindNewest(conf.CheckDir, conf.MediaFileExtensions)
	if err != nil {
		return err
	}
	// Construct the target name.
	to := filepath.Join(conf.CacheDir, filepath.Base(f))
	// Copy the file to the cache dir.
	if _, err := CopyFile(to, f); err != nil {
		return err
	}
	return nil
}

func playableFile(dir string) (string, error) {
	// scan for a file in the cache dir.
	f, err := scanner.FindNewest(dir, nil)
	if err != nil {
		return "", err
	}
	return f, nil
}

// Run starts the monitor loop for the USB directory, default is taken from
// the configuration settings.
func run(conf *config.Config) error {
	// Create the player.
	p, err := player.New(conf.PlayerExec)
	if err != nil {
		return err
	}
	log.Println("Player created.")

	// Start the notifier.
	n, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer n.Close()
	err = n.Add(conf.CheckDir)
	if err != nil {
		return err
	}

	var newFile = true
	for {
		if newFile {
			p.Stop()
			scanAndCopyFile(conf)
			f, err := playableFile(conf.CacheDir)
			if err != nil {
				log.Printf("Error finding the playable file: %s", err.Error())
			} else {
				log.Printf("New file found, restarting the player with: %s", f)
				p.Start(f)
			}
		}

		select {
		case evt := <-n.Events:
			log.Printf("fsnotify event: %s", evt.String())
			newFile = evt.Op == fsnotify.Create
			// Sleep for 5 seconds.
			<-time.After(5 * time.Second)
		}
	}
	p.Stop()
	return nil
}

func initConfig() (*config.Config, error) {
	// Process the settings.
	conf := config.NewConfig()
	conf.InitConfig()

	// Check the environment.
	if err := conf.EnvValid(); err != nil {
		return nil, fmt.Errorf("Init config failed: %s", err.Error())
	}
	return conf, nil
}

func main() {
	conf, err := initConfig()
	if err != nil {
		// Log the error and abort.
		log.Fatalf("initConfig failed:%s", err.Error())
		os.Exit(1)
	}

	// Start the processes.
	if err := run(conf); err != nil {
		log.Fatalf("Error starting the run process: %s", err.Error())
		os.Exit(2)
	}
}
