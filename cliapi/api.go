package cliapi

import (
	"errors"
	"github.com/lowl11/lazy-cli/cli_route"
)

func (cli *Cli) Middleware(middlewareHandler cli_route.Handler) *Cli {
	route := cli_route.New(middlewareHandler)
	route.Params(cli.params()...)

	cli.middlewares.Push(route)
	return cli
}

func (cli *Cli) Route(handler cli_route.Handler, path ...string) *Cli {
	if len(path) == 0 {
		return cli
	}

	route := cli_route.New(handler, path...)
	route.Params(cli.params()...)

	cli.routes.Push(route)
	return cli
}

func (cli *Cli) Start() error {
	// sort routes by path length
	cli.routes.Sort(func(i, j int) bool {
		left := cli.routes.Get(i)
		right := cli.routes.Get(j)
		return left.PathLen() > right.PathLen()
	})

	// find match route
	route := cli.matchRoute()
	if route == nil {
		return errors.New("route not found")
	}

	// run middlewares before running handler
	for _, middlewareHandler := range cli.middlewares.Slice() {
		if err := middlewareHandler.Run(); err != nil {
			return err
		}
	}

	return route.Run()
}
