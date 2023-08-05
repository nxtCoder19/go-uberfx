package main

import (
	"flag"

	"abc.io/internal/framework"
	"go.uber.org/fx"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "localhost:3000", "--addr <host:port>")
	flag.Parse()

	app := fx.New(
		fx.Provide(func() framework.HttpAddr {
			return framework.HttpAddr(addr)
		}),
		framework.Module,
		// loggerfx.Module,
		// serverfx.Module,
		// fx.Provide(
		// 	handlers.NewHandler,
		// ),
	)
	app.Run()
}