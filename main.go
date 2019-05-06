package main

import (
	"fmt"
	"github.com/rayyu03/gin-blog/models"
	"github.com/rayyu03/gin-blog/pkg/logging"
	"github.com/rayyu03/gin-blog/pkg/setting"
	"github.com/rayyu03/gin-blog/routers"
	"net/http"
)

func main()  {
	setting.Setup()
	models.Setup()
	logging.Setup()

	router := routers.InitRouter()
	server := &http.Server{
		Addr:			fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:		router,
		ReadTimeout:	setting.ServerSetting.ReadTimeout,
		WriteTimeout:	setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}



