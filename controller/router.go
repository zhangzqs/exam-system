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
		questionsGroup.DELETE("/:qid", DeleteQuestion)
		questionsGroup.PUT("/:qid", UpdateQuestion)
		questionsGroup.GET("/:qid", GetQuestion)
	}
	papersGroup := r.Group("/papers", JwtAuthMiddleware)
	{
		papersGroup.POST("/", GeneratePaper)
		papersGroup.GET("/", GetUserPapers)
		papersGroup.DELETE("/:pid")
		papersGroup.PUT("/:pid")
		papersGroup.GET("/:pid", GetPaper)
	}
}
