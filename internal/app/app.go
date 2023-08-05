package app

import (
	"abc.io/internal/domain"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"app",
	domain.Module,
	fx.Invoke(
		func (app *fiber.App, d domain.Domain) error {
			app.Get("/healthy", func(c *fiber.Ctx) error {
				s := d.SayHello()
				return c.JSON(map[string]string{"msg": s})
			})
			return nil
		},
	),
)