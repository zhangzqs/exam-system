package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/global"
	"log"
	"strconv"
	"strings"
)

func GetUid(c *gin.Context) int {
	uidAny, _ := c.Get("uid")
	return uidAny.(int)
}
func GetPageInfo(c *gin.Context) (pageId int, pageSize int) {
	var err error
	pid := c.Query("pageId")
	ps := c.Query("pageSize")
	pageId, err = strconv.Atoi(pid)
	if err != nil {
		pageId = 0
	}
	pageSize, err = strconv.Atoi(ps)
	if err != nil {
		pageSize = 10
	}
	return
}
func JwtAuthMiddleware(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		OperationNeedLoginError(c)
		c.Abort()
		return
	}
	// 去除Bearer
	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		authorizationHeader = authorizationHeader[7:]
	}
	log.Println(authorizationHeader)
	claims, err := global.GetJwt().ParseToken(authorizationHeader)
	if err != nil {
		TokenInvalidError(c)
		c.Abort()
		return
	}
	c.Set("uid", claims.Uid)
}
