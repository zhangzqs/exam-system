package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/service"
	"strconv"
)

func GeneratePaper(c *gin.Context) {
	uid := GetUid(c)

	var paper service.PaperRequestBody
	err := c.BindJSON(&paper)
	if err != nil {
		RequestFormatError(c, err)
		return
	}

	pid, err := service.AddPaperManual(uid, &paper)
	if err != nil {
		DatabaseError(c, err)
		return
	}

	SuccessfulApiResponse(c, gin.H{
		"id": pid,
	})
}

func GetPaper(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		RequestFormatError(c, err)
	}
	paper, err := service.GetPaper(GetUid(c), pid, true)
	if err != nil {
		DatabaseError(c, err)
	}
	SuccessfulApiResponse(c, paper)
}

func GetUserPapers(c *gin.Context) {
	ps, err := service.GetUserPapers(GetUid(c))
	if err != nil {
		DatabaseError(c, err)
	}
	SuccessfulApiResponse(c, gin.H{
		"papers": ps,
	})
}
