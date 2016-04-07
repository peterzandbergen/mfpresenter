// mfpresenter project main.go
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/peterzandbergen/mfpresenter/config"
	"github.com/peterzandbergen/mfpresenter/player"
	"github.com/peterzandbergen/mfpresenter/scanner"

	"github.com/fsnotify/fsnotify"
)

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
	p, err := player.NewPlayer(conf.PlayerExec)
	if err != nil {
		return err
	}
	log.Println("Player created.")
	// Perform initial scan of the check dir.
	if err := scanAndCopyFile(conf); err != nil && err != scanner.ErrNotFound {
		return err
	}

	// Get the playable file.
	f, err := playableFile(conf.CacheDir)
	if err == nil {
		p.Start(f)
		log.Printf("Player started with file: %s\n", f)
	}

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

	for {
		select {
		case <-n.Events:
			err := scanAndCopyFile(conf)
			if err != nil {
				break
			}
			f, err := playableFile(conf.CacheDir)
			if err != nil {
				break
			}
			p.Start(f)
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
		fmt.Println(err.Error())
		log.Fatalf("initConfig failed:%s", err.Error())
		os.Exit(1)
	}

	// Start the processes.
	if err := run(conf); err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Error starting the run process: %s", err.Error())
		os.Exit(2)
	}
}
