package router

import (
	"github.com/gin-gonic/gin"
	"sound2notes/handler"
)

func AddAudioRoutes(group *gin.RouterGroup) {
	entry := group.Group("/audio")

	entry.POST("", handler.PostAudioHandler)
}
