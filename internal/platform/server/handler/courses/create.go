package courses

import (
	"errors"
	"github.com/gin-gonic/gin"
	mooc "hexagonal-go-kata/internal"
	"hexagonal-go-kata/internal/creating"
	"hexagonal-go-kata/kit/command"
	"net/http"
)

// createRequest DTO for the "create course" use case
type createRequest struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation
func CreateHandler(CommandBus command.Bus) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createRequest
		if err := context.BindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := CommandBus.Dispatch(context, creating.NewCourseCommand(
			req.Id,
			req.Name,
			req.Duration,
		))

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseId),
				errors.Is(err, mooc.ErrEmptyCourseName),
				errors.Is(err, mooc.ErrEmptyCourseDuration):
				context.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				context.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		context.Status(http.StatusCreated)
	}
}
