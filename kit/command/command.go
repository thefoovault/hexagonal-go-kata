package command

import "context"

type Bus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(ctx context.Context, command Command) error
}
