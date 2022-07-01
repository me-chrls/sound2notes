package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sound2notes/entity"
)

var musicSheets = []entity.MusicSheet{
	{MusicPiece: entity.MusicPiece{ID: "1", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, Difficulty: 1},
	{MusicPiece: entity.MusicPiece{ID: "2", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, Difficulty: 5},
	{MusicPiece: entity.MusicPiece{ID: "3", Name: "name1", Composer: "composer1", Description: "abcd", StorageLocation: "/test/location"}, Difficulty: 10},
}

func GetMusicSheetsHandler(context *gin.Context) {
	context.JSON(http.StatusOK, musicSheets)
}

func GetMusicSheetHandler(context *gin.Context) {
	for i := range musicSheets {
		if musicSheets[i].ID == context.Param("musicSheetId") {
			context.JSON(http.StatusOK, musicSheets[i])
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"error": "Object with the given id does not exist",
	})
}

func PostMusicSheetHandler(context *gin.Context) {
	var newMusicSheet entity.MusicSheet
	if err := context.ShouldBindJSON(&newMusicSheet); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	musicSheets = append(musicSheets, newMusicSheet)
	context.JSON(http.StatusCreated, newMusicSheet)
}

func DeleteMusicSheetHandler(context *gin.Context) {
	for i := range musicSheets {
		if musicSheets[i].ID == context.Param("musicSheetId") {
			musicSheets[i] = musicSheets[len(musicSheets)-1]
			musicSheets = musicSheets[:(len(musicSheets) - 1)] // destroys the correct order, but is only for testing purposes
			context.Status(http.StatusNoContent)
			return
		}
	}
	context.JSON(http.StatusBadRequest, gin.H{
		"error": "Object with the given id does not exist",
	})
}
