package cli_route

import (
	"errors"
)

func (route *CliRoute) Run() error {
	if route.handler == nil {
		return errors.New("handler is null")
	}

	return route.handler(&RouteContext{
		params: route.params,
	})
}

func (route *CliRoute) Params(params ...string) {
	if len(params) == 0 {
		return
	}

	realParams := make([]string, 0, len(params))
	for _, param := range params {
		if route.controller == param || route.endpoint == param {
			continue
		}

		realParams = append(realParams, param)
	}

	route.params = realParams
}

func (route *CliRoute) PathLen() int {
	return len(route.path)
}

func (route *CliRoute) Match(arguments []string) bool {
	if len(arguments) == 0 {
		return false
	}

	if route.withEndpoint {
		if len(arguments) < 2 {
			return false
		}

		return route.controller == arguments[0] && route.endpoint == arguments[1]
	}

	return route.controller == arguments[0]
}
