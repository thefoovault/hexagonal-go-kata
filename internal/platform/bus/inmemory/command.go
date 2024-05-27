package inmemory

import (
	"context"
	"hexagonal-go-kata/kit/command"
	"log"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(context context.Context, command command.Command) error {
	handler, ok := b.handlers[command.Type()]
	if !ok {
		return nil
	}

	go func() {
		err := handler.Handle(context, command)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", command.Type(), err)
		}

	}()

	return nil
}

func (b *CommandBus) Register(commandType command.Type, handler command.Handler) {
	b.handlers[commandType] = handler
}
