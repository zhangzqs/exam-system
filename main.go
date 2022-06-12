package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/controller"
	"github.com/zhangzqs/exam-system/global"
	"log"
)

func main() {
	gin.SetMode(gin.DebugMode)
	conf := global.GetConfig()
	r := gin.Default()
	controller.InitRouter(r)
	err := r.Run(fmt.Sprintf(
		"%s:%d",
		conf.Server.ListenIp,
		conf.Server.ListenPort,
	))
	if err != nil {
		log.Fatalln(err)
	}
}
