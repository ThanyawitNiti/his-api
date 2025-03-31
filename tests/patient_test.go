package tests

import (
	"gin/controllers"
	"net/http/httptest"
	"strings"
	"testing"

	// ตำแหน่งของ mock

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/staff/create", controllers.StaffPost)
	// router.POST("/", addCommonHeaders, post)
	// router.PUT("/:id", addCommonHeaders, put)
	// router.DELETE("/:id", addCommonHeaders, remove)
	return router
}

func TestPostStaff(t *testing.T) {
	router := SetupRouter()
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/", strings.NewReader(`{
		"username": "Bloody Mary",
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
	// t.Run("Returns app name", func(t *testing.T) {
	// 	if recorder.Body.String() != "\"Cocktail service\"" {
	// 		t.Error("Expected '\"Cocktail service\"', got ", recorder.Body.String())
	// 	}
	// })
	// t.Run("Returns Server header", func(t *testing.T) {
	// 	if recorder.Header().Get("Server") != "gin-gonic/1.33.7" {
	// 		t.Error("Expected 'gin-gonic/1.33.7', got ", recorder.Header().Get("Server"))
	// 	}
	// })
}
