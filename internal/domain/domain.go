package domain

import "go.uber.org/fx"

type Domain interface{
	SayHello() string
}

type Object struct{
	greeting string
}

func (obj *Object) SayHello() string {
	return obj.greeting + "hi"
}

var Module = fx.Module(
	"domain",
	fx.Provide(func() Domain {
		return &Object{greeting: "hello"}
	}),
)