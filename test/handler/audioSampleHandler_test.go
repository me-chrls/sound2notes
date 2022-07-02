package handler

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sound2notes/entity"
	"sound2notes/handler"
	"testing"
)

func TestAudioSampleHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/audioSamples/:audioSampleId", handler.GetAudioSampleHandler)
	req, _ := http.NewRequest("GET", "/audioSamples/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAudioSampleHandlerNotFound(t *testing.T) {
	mockResponse := `{"error":"Object with the given id does not exist"}`

	r := SetUpRouter()
	r.GET("/audioSamples/:audioSampleId", handler.GetAudioSampleHandler)
	req, _ := http.NewRequest("GET", "/audioSamples/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}

func TestGetAudioSampleHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.POST("/audioSamples", handler.GetAudioSamplesHandler)
	req, _ := http.NewRequest("POST", "/audioSamples", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var musicSheets []entity.MusicSheet
	json.Unmarshal(w.Body.Bytes(), &musicSheets)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, musicSheets)
}

func TestPostAudioSampleHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.POST("/audioSamples", handler.PostAudioSampleHandler)
	musicSheet := entity.AudioSample{
		MusicPiece: entity.MusicPiece{
			ID:              "4",
			Name:            "name",
			Composer:        "composer",
			Description:     "description",
			StorageLocation: "storageLocation",
		},
		StorageFormat: "mp3",
	}

	jsonValue, _ := json.Marshal(musicSheet)
	req, _ := http.NewRequest("POST", "/audioSamples", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostAudioSampleHandlerNil(t *testing.T) {
	mockResponse := `{"error":"invalid request"}`

	r := SetUpRouter()
	r.POST("/audioSamples", handler.PostAudioSampleHandler)

	req, _ := http.NewRequest("POST", "/audioSamples", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}

func TestDeleteAudioSampleHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/audioSamples/:audioSampleId", handler.DeleteAudioSampleHandler)
	req, _ := http.NewRequest("DELETE", "/audioSamples/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteAudioSampleHandlerNotFound(t *testing.T) {
	mockResponse := `{"error":"Object with the given id does not exist"}`

	r := SetUpRouter()
	r.GET("/audioSamples/:audioSampleId", handler.DeleteAudioSampleHandler)
	req, _ := http.NewRequest("GET", "/audioSamples/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}
