package cliapi

import (
	"github.com/lowl11/lazy-cli/cli_route"
	"github.com/lowl11/lazy-collection/array"
)

type Cli struct {
	arguments []string
	routes    *array.Array[cli_route.CliRoute]
}

func New() *Cli {
	return &Cli{
		arguments: Arguments(),
		routes:    array.New[cli_route.CliRoute](),
	}
}
