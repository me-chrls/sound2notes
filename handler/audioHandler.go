package handler

import (
	"github.com/gin-gonic/gin"
	"sound2notes/utils/upload"
)

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
