package main

import (
	"fmt"
	"net/http"

	"github.com/ckpns/media-sharing-platform/backend/src/common"
	"github.com/ckpns/media-sharing-platform/backend/src/controllers"
	"github.com/ckpns/media-sharing-platform/backend/src/data"
	"github.com/gin-gonic/gin"
)

func main() {

	common.InitConfig()
	data.InitDB()

	router := gin.Default()

	// Ping endpoint
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Static("/file", "./public")

	controllers.UserRoutesRegister(router)
	controllers.PostRoutesRegister(router)

	config := common.GetConfig()

	router.MaxMultipartMemory = config.MaxMultipartMemory

	router.Run(fmt.Sprintf("%s:%s", config.Host, config.Port))
}
