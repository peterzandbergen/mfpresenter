package config

import (
	"os"
	"testing"
)

// Test if the values of a new configuration are correct.
func TestNewConfig(t *testing.T) {
	c := NewConfig()

	if c.CheckDir != DefaultCheckDir {
		t.Errorf("CheckDir value is not correct: %s", c.CheckDir)
	}
	if c.CacheDir != DefaultCacheDir {
		t.Errorf("CheckDir value is not correct: %s", c.CacheDir)
	}
	if c.PlayerExec != DefaultPlayerExec {
		t.Errorf("PlayerProgram value is not correct: %s", c.PlayerExec)
	}
	if len(c.MediaFileExtensions) != 1 {
		t.Errorf("MediaExtensions len not equal to 1: %d", len(c.MediaFileExtensions))
	}

	if c.MediaFileExtensions[0] != "mp4" {
		t.Errorf("MediaExtension[0] not equal to \"mp4\": %s", c.MediaFileExtensions[0])
	}
}

const testRootDir = "/home/peza/Documents/workspace-fileplayer/test/mfpresenter/"
const testCacheDir = testRootDir + "cachedir"
const testCheckDir = testRootDir + "checkdir"
const testPlayerExec = testRootDir + "test.exec.sh"

func TestCheckConfig(t *testing.T) {
	c := &Config{
		CheckDir:   testCheckDir,
		CacheDir:   testCacheDir,
		PlayerExec: testPlayerExec,
	}

	if err := c.EnvValid(); err != nil {
		t.Errorf("envValid failed: %s", err.Error())
	}
}

func TestShowPermissions(t *testing.T) {
	var fi os.FileInfo
	var err error
	var path string

	// 	t.Logf("0x400 = %d", 0400)

	path = testCacheDir
	fi, err = os.Stat(path)
	if err != nil {
		t.Errorf("Cannot stat %s: %s", path, err.Error())
	}
	// t.Logf("fi of %s: %s", path, fi.Mode().Perm().String())

	path = testCheckDir
	fi, err = os.Stat(path)
	if err != nil {
		t.Errorf("Cannot stat %s: %s", path, err.Error())
	}
	_ = fi
	// t.Logf("fi of %s: %s", path, fi.Mode().Perm().String())
}
