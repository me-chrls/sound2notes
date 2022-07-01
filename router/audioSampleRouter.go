package router

import (
	"github.com/gin-gonic/gin"
	"sound2notes/handler"
)

func AddAudioSampleRoutes(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/audioSamples")

	users.POST("", handler.PostAudioSampleHandler)
	users.GET("/:audioSampleId", handler.GetAudioSampleHandler)
	users.GET("", handler.GetAudioSamplesHandler)
	users.DELETE("/:audioSampleId", handler.DeleteAudioSampleHandler)
}
