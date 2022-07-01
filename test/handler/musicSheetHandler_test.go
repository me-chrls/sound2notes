package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sound2notes/entity"
	"sound2notes/handler"
	"testing"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetMusicSheetHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.GET("/musicSheets/:musicSheetId", handler.GetMusicSheetHandler)
	req, _ := http.NewRequest("GET", "/musicSheets/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetMusicSheetHandlerNotFound(t *testing.T) {
	mockResponse := `{"error":"Object with the given id does not exist"}`

	r := SetUpRouter()
	r.GET("/musicSheets/:musicSheetId", handler.GetMusicSheetHandler)
	req, _ := http.NewRequest("GET", "/musicSheets/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}

func TestGetMusicSheetsHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.POST("/musicSheets", handler.GetMusicSheetsHandler)
	req, _ := http.NewRequest("POST", "/musicSheets", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var musicSheets []entity.MusicSheet
	json.Unmarshal(w.Body.Bytes(), &musicSheets)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, musicSheets)
}

func TestPostMusicSheetHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.POST("/musicSheets", handler.PostMusicSheetHandler)
	musicSheet := entity.MusicSheet{
		MusicPiece: entity.MusicPiece{
			ID:              "4",
			Name:            "name",
			Composer:        "composer",
			Description:     "description",
			StorageLocation: "storageLocation",
		},
		Difficulty: 5,
	}

	jsonValue, _ := json.Marshal(musicSheet)
	req, _ := http.NewRequest("POST", "/musicSheets", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestPostMusicSheetHandlerNil(t *testing.T) {
	mockResponse := `{"error":"invalid request"}`

	r := SetUpRouter()
	r.POST("/musicSheets", handler.PostMusicSheetHandler)

	req, _ := http.NewRequest("POST", "/musicSheets", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}

func TestDeleteMusicSheetHandlerSuccess(t *testing.T) {
	r := SetUpRouter()
	r.DELETE("/musicSheets/:musicSheetId", handler.DeleteMusicSheetHandler)
	req, _ := http.NewRequest("DELETE", "/musicSheets/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteMusicSheetHandlerNotFound(t *testing.T) {
	mockResponse := `{"error":"Object with the given id does not exist"}`

	r := SetUpRouter()
	r.GET("/musicSheets/:musicSheetId", handler.GetMusicSheetHandler)
	req, _ := http.NewRequest("GET", "/musicSheets/5", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, mockResponse, string(responseData))
}
