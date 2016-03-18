package main

import (
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
)

// TestWatcher is a main loop for testing the fsnotify functionality.
func TestWatcher() {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Error creating Watcher: %s", err.Error())
	}
	defer w.Close()
	w.Add("/media/peza")

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Println("Starting watch...")
		for {
			select {
			case ev := <-w.Events:
				log.Printf("Event: %s, %d", ev.Name, ev.Op)
				// case err := <- w.Errors:
			}
		}

	}()
	wg.Wait()
}

func main() {
	log.Print("Starting the program.")
	TestWatcher()
}
