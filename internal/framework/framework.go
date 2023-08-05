package framework

import (
	"context"
	"log"

	"abc.io/internal/app"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type HttpAddr string

var Module = fx.Module(
	"framework",
	fx.Provide(func() *fiber.App {
		return fiber.New()
	}),

	fx.Invoke(func (app *fiber.App, lf fx.Lifecycle, addr HttpAddr) {
		 lf.Append(fx.Hook{
			OnStart: func( context.Context) error {
				go func() {
					log.Fatalln(app.Listen(string(addr)))
				}()
				return nil
			},
		})
	}),
	app.Module,
)

