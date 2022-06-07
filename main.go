package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/controller"
)

func main() {
	r := gin.Default()
	controller.InitRouter(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
