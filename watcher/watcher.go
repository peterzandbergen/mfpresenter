/*
Package watcher monitors a set of directories for changes.
Changes can be Create, Write, Remove, Rename, Chmod.

We are mostly interested in Create and Remove of directories.
*/
package watcher

import "github.com/fsnotify/fsnotify"


// Dir watches a directory for changes.
type Dir struct {
    w *fsnotify.Watcher
    // Events     
}

func New() (* Dir, error) {
    return nil, nil 
}

