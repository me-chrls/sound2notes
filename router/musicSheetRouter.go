package router

import (
	"github.com/gin-gonic/gin"
	"sound2notes/handler"
)

func AddMusicSheetRoutes(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/musicSheets")

	users.POST("", handler.PostMusicSheetHandler)
	users.GET("/:musicSheetId", handler.GetMusicSheetHandler)
	users.GET("", handler.GetMusicSheetsHandler)
	users.DELETE("/:musicSheetId", handler.DeleteMusicSheetHandler)
}
