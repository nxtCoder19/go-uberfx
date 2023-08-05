package serverfx

import (
	"net/http"

	"context"

	"go.uber.org/fx"
	// "go.uber.org/fx/internal/lifecycle"
)

func RegisterHandlers(handler http.Handler){
	http.Handle("/", handler)
}

func InitServer(lifecycle fx.Lifecycle) {
	server := &http.Server{
		Addr: ":8080",
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error{
			go server.ListenAndServe()
			return nil
		},
		OnStop: func (ctx context.Context) error {
			return server.Close()
		},
	})
}

var Module = fx.Options(
	fx.Invoke(
		RegisterHandlers,
		InitServer,
	),
)
