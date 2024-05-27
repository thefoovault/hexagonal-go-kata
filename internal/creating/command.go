package creating

import (
	"context"
	"errors"
	"hexagonal-go-kata/kit/command"
)

const CourseCommandType = "command.creating.course"

type CourseCommand struct {
	id       string
	name     string
	duration string
}

func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c CourseCommand) Type() command.Type {
	return CourseCommandType
}

type CourseCommandHandler struct {
	service CourseService
}

func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

func (c CourseCommandHandler) Handle(context context.Context, cmd command.Command) error {
	createCourseCmd, ok := cmd.(CourseCommand)
	if ok != true {
		return errors.New("unexpected command")
	}

	return c.service.CreateCourse(
		context,
		createCourseCmd.id,
		createCourseCmd.name,
		createCourseCmd.duration,
	)
}
