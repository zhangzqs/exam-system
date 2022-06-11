package controller

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	r.GET("/ping", Ping)
	r.POST("/login", Login)
	r.POST("/register", Register)

	r.POST("/questions", AddQuestion)
	r.DELETE("/questions/:id", DeleteQuestion)
	r.PUT("/questions/:id", UpdateQuestion)
	r.GET("/questions/:id", GetQuestion)
	r.GET("/users/:uid/questions", GetUserQuestions)

}
