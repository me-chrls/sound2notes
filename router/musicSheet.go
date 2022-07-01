package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMusicSheetRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/musicSheets")

	users.GET("/musicSheetId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})
	users.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	users.DELETE("/musicSheetId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete")
	})
}
