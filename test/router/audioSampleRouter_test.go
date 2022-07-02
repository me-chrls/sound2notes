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

func SetUpAudioSampleRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	router.AddAudioSampleRoutes(v1)
	return r
}

func TestAddAudioSampleRoutesPostRequest(t *testing.T) {
	mockResponse := "post"

	r := SetUpAudioSampleRouter()
	req, _ := http.NewRequest("POST", "/v1/audioSamples", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, response)
}

func TestAddAudioSampleRoutesGetRequest(t *testing.T) {
	mockResponse := "get"

	r := SetUpAudioSampleRouter()
	req, _ := http.NewRequest("GET", "/v1/audioSamples/audioSampleId", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, response)
}

func TestAddAudioSampleRoutesDeleteRequest(t *testing.T) {
	mockResponse := "delete"

	r := SetUpAudioSampleRouter()
	req, _ := http.NewRequest("DELETE", "/v1/audioSamples/audioSampleId", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var response string
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, mockResponse, response)
}
