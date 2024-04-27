package courses

import (
	"github.com/gin-gonic/gin"
	mooc "go_test/internal/platform"
	"net/http"
)

// createRequest DTO for the "create course" use case
type createRequest struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation
func CreateHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		var req createRequest
		if err := context.BindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.Id, req.Name, req.Duration)
		context.String(http.StatusCreated, "Creating the course with Id: %s, name: %s and %s description", course.Id(), course.Name(), course.Duration())
	}
}
