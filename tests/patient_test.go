package tests

import (
	"gin/controllers"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/staff/create", controllers.StaffPost)
	return router
}

func TestPostStaff(t *testing.T) {
	router := SetupRouter()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/staff/create", strings.NewReader(`{
		"username": "mary",
		"password": "demo@1234",
		"hospital_id":1
	}`))
	request.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(recorder, request)
	t.Run("Returns 200 status code", func(t *testing.T) {
		if recorder.Code != 200 {
			t.Error("Expected 200, got ", recorder.Code)
		}
	})

}
