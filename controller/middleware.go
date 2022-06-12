package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/global"
	"log"
	"strings"
)

func GetUid(c *gin.Context) int {
	uidAny, _ := c.Get("uid")
	return uidAny.(int)
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
