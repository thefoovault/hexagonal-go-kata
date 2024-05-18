package courses

import (
	"github.com/gin-gonic/gin"
	mooc "hexagonal-go-kata/internal"
	"net/http"
)

// createRequest DTO for the "create course" use case
type createRequest struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation
func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createRequest
		if err := context.BindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := mooc.NewCourse(req.Id, req.Name, req.Duration)
		if err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := courseRepository.Save(context, course); err != nil {
			context.JSON(http.StatusInternalServerError, err.Error())
		}

		context.Status(http.StatusCreated)
	}
}
