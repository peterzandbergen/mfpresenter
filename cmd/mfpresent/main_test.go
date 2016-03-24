package main

import (
	"testing"
)

// Test if the values of a new configuration are correct.
func TestNewConfig(t *testing.T) {
	c := newConfig()

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

func TestCheckConfig(t *testing.T) {
	c := Config{}
	if err := envValid(c); err != nil {
		t.Errorf("envValid failed: %s", err.Error())
	}
}
