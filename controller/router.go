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
	questionsGroup := r.Group("/questions", JwtAuthMiddleware)
	{
		questionsGroup.POST("/", AddQuestion)
		questionsGroup.GET("/", GetUserQuestions)
		questionsGroup.DELETE("/:id", DeleteQuestion)
		questionsGroup.PUT("/:id", UpdateQuestion)
		questionsGroup.GET("/:id", GetQuestion)
	}
	papersGroup := r.Group("/papers", JwtAuthMiddleware)
	{
		papersGroup.POST("/")
		papersGroup.GET("/")
		papersGroup.DELETE("/:id")
		papersGroup.PUT("/:id")
		papersGroup.GET("/:id")
	}
}
