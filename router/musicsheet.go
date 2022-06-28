package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addMusicSheetRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/musicsheets")

	users.GET("/musicsheetId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})
	users.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	users.DELETE("/musicsheetid", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete")
	})
}
