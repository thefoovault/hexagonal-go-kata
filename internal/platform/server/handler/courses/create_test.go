package courses

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"hexagonal-go-kata/internal/platform/storage/storagemocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	tests := []struct {
		name            string
		expectedStatus  int
		createCourseReq createRequest
		wantErr         bool
	}{
		{
			name:           `given a payload with missing fields request it returns 400`,
			expectedStatus: http.StatusBadRequest,
			createCourseReq: createRequest{
				Id:   "8a1c5cdc-ba57-445a-994d-aa412d23723f",
				Name: "Demo Course",
			},
			wantErr: true,
		},
		{
			name:           `given another payload with invalid fields it returns 400`,
			expectedStatus: http.StatusBadRequest,
			createCourseReq: createRequest{
				Id:       "invalid uuid",
				Name:     "Demo Course",
				Duration: "10 minutes",
			},
			wantErr: true,
		},
		{
			name:           `given a valid request it returns 201`,
			expectedStatus: http.StatusCreated,
			createCourseReq: createRequest{
				Id:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
				Name:     "Demo Course",
				Duration: "10 months",
			},
			wantErr: true,
		},
	}
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.createCourseReq)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)
		})
	}
}
