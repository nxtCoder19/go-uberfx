package loggerfx

import (
	"log"
	"os"

	"go.uber.org/fx"
)

type Logger interface{
	Println(v ...interface{})
}

func NewLogger() Logger {
	return log.New(os.Stdout, "[Logger]", 0)
}

var Module  = fx.Options(
	fx.Provide(
		NewLogger,
	),
)