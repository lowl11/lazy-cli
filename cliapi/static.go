package cliapi

import (
	"flag"
	"os"
)

func Arguments() []string {
	flag.Parse()
	return os.Args[1:]
}
