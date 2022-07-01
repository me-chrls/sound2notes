package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddAudioSampleRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/audioSamples")

	users.GET("/audioSampleId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})
	users.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})
	users.DELETE("/audioSampleId", func(c *gin.Context) {
		c.JSON(http.StatusOK, "delete")
	})
}
