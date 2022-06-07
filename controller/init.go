package controller

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.POST("/login", Login)
	r.POST("/register", Register)
}
