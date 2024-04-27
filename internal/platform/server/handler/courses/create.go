package courses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusCreated, "Creating the course")
	}
}
