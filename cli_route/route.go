package cli_route

import "strings"

type Handler func(ctx IContext) error

type IContext interface {
	Params() []string
}

type RouteContext struct {
	params []string
}

func (ctx *RouteContext) Params() []string {
	return ctx.params
}

type CliRoute struct {
	path       []string
	controller string
	endpoint   string

	params []string

	withEndpoint bool

	handler Handler
}

func New(handler Handler, path ...string) CliRoute {
	// if path is empty, it is middleware
	if len(path) == 0 {
		return CliRoute{
			handler: handler,
		}
	}

	// endpoint
	var endpoint string
	var withEndpoint bool

	if len(path) > 1 {
		endpoint = path[1]
		withEndpoint = true
	}

	return CliRoute{
		path: path,

		handler: handler,

		controller: strings.TrimSpace(path[0]),
		endpoint:   strings.TrimSpace(endpoint),

		withEndpoint: withEndpoint,
	}
}
