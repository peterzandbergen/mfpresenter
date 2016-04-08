package mfpresenter

import (
	"log"
	"os"
)

var std *log.Logger

func SetDefault(l *log.Logger) {
	std = l
}

func Logger() *log.Logger {
	return std
}

func init() {
	std = log.New(os.Stdout, "", log.LstdFlags)
}
