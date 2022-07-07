package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"sound2notes/utils"
	"sound2notes/utils/upload"
)

var (
	router = gin.Default()
	port   = ":8080"
)

func Run() {
	utils.SetupIp()
	utils.SetPort(port)

	err := router.SetTrustedProxies(nil) // https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	if err != nil {
		log.Fatal("SetTrustedProxies:\n", err)
		return
	}

	getRoutes()

	upload.ConfigureUpload(router)

	router.Run(utils.Port)
}

func getRoutes() {
	v1 := router.Group("/v1")
	AddMusicSheetRoutes(v1)
	AddAudioSampleRoutes(v1)
	AddAudioRoutes(v1)
}
