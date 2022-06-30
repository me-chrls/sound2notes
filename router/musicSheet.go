package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addMusicSheetRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/musicsheets")

	users.GET("/musicSheetId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})
	users.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	users.DELETE("/musicSheetid", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete")
	})
}
