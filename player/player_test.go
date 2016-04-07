package player

import (
	"testing"
	"time"
)

const exe = "sleep"
const parm = "100"
const exetest = "vlc --fullscreen --loop"
const file = "/media/sf_Peter/Desktop/raspitest.mp4"

func TestPlayShell(t *testing.T) {
	p, err := NewPlayer(exetest)
	if err != nil {
		t.Fatalf("Error creating player: %s", err.Error())
	}
	t.Log(p.exe)
	t.Log(p.args)
	err = p.Start(file)
	if err != nil {
		t.Fatalf("Error starting process: %s", err.Error())
	}

	// Sleep for 3 seconds.
	time.Sleep(3 * time.Second)
	//	// Stop the player.
	//	t.Log("Awake")
	//	err = p.Stop()
	//	if err != nil {
	//		t.Errorf("Error stopping the process: %s", err.Error())
	//	}

	//	time.Sleep(3 * time.Second)
	err = p.Start(file)
	if err != nil {
		t.Fatalf("Error starting process: %s", err.Error())
	}

	// Sleep for 30 seconds.
	time.Sleep(4 * time.Second)
	// Stop the player.
	t.Log("Awake")
	err = p.Stop()
	if err != nil {
		t.Errorf("Error stopping the process: %s", err.Error())
	}

}
