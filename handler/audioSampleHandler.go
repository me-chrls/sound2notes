package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sound2notes/entity"
)

var audioSamples = []entity.AudioSample{
	{MusicPiece: entity.MusicPiece{ID: "1", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, StorageFormat: "mp3"},
	{MusicPiece: entity.MusicPiece{ID: "2", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, StorageFormat: "mp3"},
	{MusicPiece: entity.MusicPiece{ID: "3", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, StorageFormat: "mp3"},
}

func GetAudioSamplesHandler(context *gin.Context) {
	context.JSON(http.StatusOK, audioSamples)
}

func GetAudioSampleHandler(context *gin.Context) {
	for i := range audioSamples {
		if audioSamples[i].ID == context.Param("audioSampleId") {
			context.JSON(http.StatusOK, audioSamples[i])
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"error": "Object with the given id does not exist",
	})
}

func PostAudioSampleHandler(context *gin.Context) {

	var newAudioSample entity.AudioSample
	if err := context.ShouldBindJSON(&newAudioSample); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	audioSamples = append(audioSamples, newAudioSample)
	context.JSON(http.StatusCreated, newAudioSample)
}

func DeleteAudioSampleHandler(context *gin.Context) {
	for i := range musicSheets {
		if audioSamples[i].ID == context.Param("audioSampleId") {
			audioSamples[i] = audioSamples[len(audioSamples)-1]
			audioSamples = audioSamples[:(len(audioSamples) - 1)] // destroys the correct order, but is only for testing purposes
			context.Status(http.StatusNoContent)
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"error": "Object with the given id does not exist",
	})
}
