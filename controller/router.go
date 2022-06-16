package controller

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.GET("/ping", Ping)
	sessionGroup := r.Group("/session") // 会话管理
	{
		sessionGroup.POST("/login", Login)       // 登录
		sessionGroup.POST("/register", Register) // 注册
	}
	questionsGroup := r.Group("/questions", JwtAuthMiddleware) // 题库管理
	{
		questionsGroup.POST("/", AddQuestion)     // 增加题目
		questionsGroup.GET("/", GetUserQuestions) // 获取用户创建的题库
		questionsGroup.DELETE("/:qid")            // 删除用户添加的题目
		questionsGroup.PUT("/:qid")               // 修改用户添加的题目
		questionsGroup.GET("/:qid", GetQuestion)  // 获取某道题
	}
	papersGroup := r.Group("/papers", JwtAuthMiddleware) // 试卷管理
	{
		papersGroup.POST("/", GeneratePaper) // 生成试卷
		papersGroup.GET("/", GetUserPapers)  // 获取用户创建的试卷
		papersGroup.DELETE("/:pid")          // 删除用户创建的试卷
		papersGroup.PUT("/:pid")             // 修改用户创建的试卷
		papersGroup.GET("/:pid", GetPaper)   // 获取某张试卷的所有信息
	}
	roomsGroup := r.Group("/rooms", JwtAuthMiddleware)
	{
		roomsGroup.POST("/", CreateRoom)  // 创建考场
		roomsGroup.GET("/", GetUserRooms) // 获取创建的所有考场列表
		roomsGroup.DELETE("/:rid")        // 删除考场信息
		roomsGroup.PUT("/:rid")           // 修改考场信息
		roomsGroup.GET("/:rid", GetRoom)  // 获取考场信息

		roomGroup := roomsGroup.Group("/:rid")
		{
			studentsGroup := roomGroup.Group("/users")
			{
				studentsGroup.POST("/", AddStudents)
				studentGroup := studentsGroup.Group("/:uid") // 考场中的考生
				{
					studentGroup.GET("/") // 获取考生成绩与评语信息
					studentGroup.PUT("/") // 评语修改及是否选择发放成绩信息
				}
			}
		}

		testGroup := roomGroup.Group("/test")
		{
			// 考生调用
			testGroup.GET("/enter", EnterRoom)            //进入考场获得试卷开始考试
			testGroup.POST("/submitAnswer", SubmitAnswer) // 提交答案
			testGroup.POST("/submitPaper", SubmitPaper)   // 提交答案
			testGroup.GET("/detail", GetTestResult)       // 获取考生答卷详情
		}
	}
}
