package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sound2notes/router"
	"testing"
)

func SetUpMusicSheetRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	router.AddMusicSheetRoutes(v1)
	return r
}

func TestAddMusicSheetRoutesPostRequest(t *testing.T) {
	mockResponse := "post"

	r := SetUpMusicSheetRouter()
	req, _ := http.NewRequest("POST", "/v1/musicSheets", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, response)
}

func TestAddMusicSheetRoutesGetRequest(t *testing.T) {
	mockResponse := "get"

	r := SetUpMusicSheetRouter()
	req, _ := http.NewRequest("GET", "/v1/musicSheets/musicSheetId", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, response)
}

func TestAddMusicSheetRoutesDeleteRequest(t *testing.T) {
	mockResponse := "delete"

	r := SetUpMusicSheetRouter()
	req, _ := http.NewRequest("DELETE", "/v1/musicSheets/musicSheetId", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, response)
}
