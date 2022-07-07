package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sound2notes/entity"
	"sound2notes/utils/upload"
)

var audioSamples1 = []entity.AudioSample{
	{MusicPiece: entity.MusicPiece{ID: "1", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, StorageFormat: "mp3"},
	{MusicPiece: entity.MusicPiece{ID: "2", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, StorageFormat: "mp3"},
	{MusicPiece: entity.MusicPiece{ID: "3", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, StorageFormat: "mp3"},
}

func GetAudiosHandler(context *gin.Context) {
	context.JSON(http.StatusOK, audioSamples)
}

func GetAudioHandler(context *gin.Context) {
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

// file: File
func PostAudioHandler(context *gin.Context) {
	/*
		var newAudioSample entity.AudioSample
		if err := context.ShouldBindJSON(&newAudioSample); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		audioSamples1 = append(audioSamples1, newAudioSample)
		context.JSON(http.StatusCreated, newAudioSample)*/
	upload.Upload(context, "file")
}

func DeleteAudioHandler(context *gin.Context) {
	for i := range musicSheets {
		if audioSamples[i].ID == context.Param("audioSampleId") {
			audioSamples1[i] = audioSamples1[len(audioSamples1)-1]
			audioSamples1 = audioSamples1[:(len(audioSamples1) - 1)] // destroys the correct order, but is only for testing purposes
			context.Status(http.StatusNoContent)
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"error": "Object with the given id does not exist",
	})
}
