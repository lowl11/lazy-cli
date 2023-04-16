package cli_tools

import "log"

func Print(args ...any) {
	log.SetPrefix("")
	log.Println(args...)
}
