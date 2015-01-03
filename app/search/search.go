package search

import (
	"os"
	"log"
)

var l *log.Logger

func init()  {
	l = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
}
