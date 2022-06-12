package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", Ping)
	sessionGroup := r.Group("/session")
	{
		sessionGroup.POST("/login", Login)
		sessionGroup.POST("/register", Register)
	}
	questionsGroup := r.Group("/questions")
	questionsGroup.Use(JwtAuthMiddleware)
	{
		questionsGroup.POST("/", AddQuestion)
		questionsGroup.DELETE("/:id", DeleteQuestion)
		questionsGroup.PUT("/:id", UpdateQuestion)
		questionsGroup.GET("/:id", GetQuestion)
	}
	r.GET("/users/:uid/questions", GetUserQuestions)

}
