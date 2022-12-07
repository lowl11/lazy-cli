package cliapi

import "os"

func Arguments() []string {
	return os.Args[1:]
}
