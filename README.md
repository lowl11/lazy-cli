# lazy-cli

> Library for routing in CLI application

Example
```go
func main() {
	cli := cliapi.New()
	cli.Route(controllerHandler, "controller")
	cli.Route(controllerNewHandler, "controller", "new")

	if err := cli.Start(); err != nil {
		log.Fatal("start cli error: ", err)
	}
}

func controllerHandler(ctx cli_route.IContext) error {
	fmt.Println("controllerHandler()")
	fmt.Println("params:", ctx.Params())
	return nil
}

func controllerNewHandler(ctx cli_route.IContext) error {
	fmt.Println("controllerNewHandler()")
	fmt.Println("params:", ctx.Params())
	return nil
}
```
