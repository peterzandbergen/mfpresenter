package scanner

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var ErrNotFound = errors.New("No matching file found.")

func buildExtRexexp(exts []string) (*regexp.Regexp, error) {
	var e []string

	for _, s := range exts {
		e = append(e, `\.`+s)
	}
	res := "(" + strings.Join(e, "|") + ")$"
	return regexp.Compile(res)
}

// find newest returns the newest file that ends in one of the extensions.
func FindNewest(path string, exts []string) (string, error) {
	var newestPath string
	var newestFi os.FileInfo
	var re *regexp.Regexp
	var err error
	// Build the reg expression to test if we have exts.
	if exts != nil {
		re, err = buildExtRexexp(exts)
		if err != nil {
			return "", err
		}
	}
	// Make path absolute.
	ap, err := filepath.Abs(path)
	if err == nil {
		path = ap
	}
	filepath.Walk(path, func(p string, fi os.FileInfo, err error) error {
		// Test for an error.
		if err != nil {
			return nil
		}
		// Only regular files.
		if !fi.Mode().IsRegular() {
			return nil
		}
		// Test if the file name matches the expression.
		// Only check if re != nil.
		if re == nil || re.Match([]byte(p)) {
			// Test if the file is newer than the current one.
			if newestFi == nil || fi.ModTime().After(newestFi.ModTime()) {
				newestFi = fi
				newestPath = p
			}
		}
		return nil
	})
	if newestFi == nil {
		return "", ErrNotFound
	}
	return newestPath, nil
}
