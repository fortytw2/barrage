package home

import (
	"log"
	"os"
)

var l *log.Logger

func init() {
	l = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
}
