package router

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func Run() {
	getRoutes()
	router.Run(":8080")
}

func getRoutes() {
	v1 := router.Group("/v1")
	AddMusicSheetRoutes(v1)
	AddAudioSampleRoutes(v1)
}
