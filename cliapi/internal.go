package cliapi

import "github.com/lowl11/lazy-cli/cli_route"

func (cli *Cli) matchRoute() *cli_route.CliRoute {
	return cli.routes.Single(func(iterator cli_route.CliRoute) bool {
		return iterator.Match(cli.arguments)
	})
}

func (cli *Cli) params() []string {
	return cli.arguments
}
