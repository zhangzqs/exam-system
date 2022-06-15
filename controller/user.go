package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzqs/exam-system/global"
	"github.com/zhangzqs/exam-system/service"
)

type loginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponseBody struct {
	Uid      int    `json:"uid"`
	JwtToken string `json:"jwtToken"`
}

func Login(c *gin.Context) {
	var rb loginRequestBody
	err := c.BindJSON(&rb)
	if err != nil {
		RequestFormatError(c)
		return
	}
	uid, err := service.Login(rb.Username, rb.Password)
	if err != nil {
		LoginError(c)
		return
	}

	SuccessfulApiResponse(c, loginResponseBody{
		Uid:      uid,
		JwtToken: global.GetJwt().GenerateToken(uid),
	})
}

func Register(c *gin.Context) {
	var rb loginRequestBody
	err := c.BindJSON(&rb)
	if err != nil {
		RequestFormatError(c)
		return
	}
	uid, err := service.Register(rb.Username, rb.Password)
	if err != nil {
		RegisterUserExistsError(c)
		return
	}
	SuccessfulApiResponse(c, loginResponseBody{
		Uid:      uid,
		JwtToken: global.GetJwt().GenerateToken(uid),
	})
}
