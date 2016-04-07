package scanner

import (
	"regexp"
	"testing"
)

const extMp4RegExp = `(\.mp4|\.mp3|\.divx)$`

func TestBuildExtRexexp(t *testing.T) {
	var exts = []string{
		"mp4",
		"mp5",
	}
	re, err := buildExtRexexp(exts)
	if err != nil {
		t.Errorf("buildExtRexexp failed: %s", err.Error())
	}
	_ = re
	if !re.Match([]byte("playfile.mp4")) {
		t.Error("regexp.Match failed")
	}
	if !re.Match([]byte("playfile.mp5")) {
		t.Error("regexp.Match failed on mp5")
	}
	if re.Match([]byte("playfile.mp4.divx")) {
		t.Error("regexp.Match passed on divx")
	}

}

func TestRegExp(t *testing.T) {
	re, err := regexp.Compile(extMp4RegExp)
	if err != nil {
		t.Errorf("regexp compile failed: %s", err.Error())
	}
	_ = re
	if !re.Match([]byte("playfile.mp4")) {
		t.Error("regexp.Match failed")
	}
	if !re.Match([]byte("playfile.mp4")) {
		t.Error("regexp.Match failed")
	}
	if !re.Match([]byte("playfile.mp4.divx")) {
		t.Error("regexp.Match failed")
	}
}

//func TestFindNewestPass(t *testing.T) {
//	var exts = []string{
//		"mp4",
//		"mp5",
//		"mp6",
//	}
//	s, err := findNewest(testCheckDir, exts)
//	// Should pass.
//	if err != nil {
//		t.Errorf("error: %s", err.Error())
//	}
//	_ = s
//	// t.Logf("found path: %s", s)
//}
