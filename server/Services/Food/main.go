package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"Food/helpers/gredis"
	"Food/helpers/logging"
	"Food/helpers/setting"
	"Food/helpers/util"
	"Food/repository"
	"Food/routers"
)

func init() {
	setting.Setup()
	repository.Setup()
	logging.Setup()
	gredis.Setup()
	util.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @license.name MIT
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%s", setting.ServerSetting.HTTPPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// if setting.ServerSetting.SSL {

	// 	SSLKeys := &struct {
	// 		CERT string
	// 		KEY  string
	// 	}{}

	// 	//Generated using sh generate-certificate.sh
	// 	SSLKeys.CERT = "./cert/myCA.cer"
	// 	SSLKeys.KEY = "./cert/myCA.key"

	// 	routersInit.RunTLS(":"+setting.ServerSetting.HTTPPort, SSLKeys.CERT, SSLKeys.KEY)
	// } else {
	// 	routersInit.Run(":" + setting.ServerSetting.HTTPPort)
	// }

	repository.CloseDB()
}
