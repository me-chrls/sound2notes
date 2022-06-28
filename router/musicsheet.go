package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func addMusicSheetRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/musicsheets/id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})
	users.POST("/musicsheets", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	users.DELETE("/musicsheets/id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete")
	})
}
