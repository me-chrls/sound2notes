package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"sound2notes/utils"
	"sound2notes/utils/upload"
)

var (
	router = gin.Default()
)

func Run() {
	err := router.SetTrustedProxies([]string{"nil"}) // https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies
	if err != nil {
		log.Fatal("SetTrustedProxies:\n", err)
		return
	}

	getRoutes()

	upload.ConfigureUpload(router)
	utils.SetupIp()

	router.Run(":8080")
}

func getRoutes() {
	v1 := router.Group("/v1")
	AddMusicSheetRoutes(v1)
	AddAudioSampleRoutes(v1)
	AddAudioRoutes(v1)
}
