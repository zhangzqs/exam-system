package controller

import (
	"github.com/gin-gonic/gin"
)

type requestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type responseBody struct {
	JwtToken string `json:"jwtToken"`
}

func Login(c *gin.Context) {
	var rb requestBody
	err := c.BindJSON(&rb)
	if err != nil {
		UnknownErrorApiResponse(c, err.Error())
		return
	}
	SuccessfulApiResponse(c, rb)
}

func Register(c *gin.Context) {
	var rb requestBody
	err := c.BindJSON(&rb)
	if err != nil {
		UnknownErrorApiResponse(c, err.Error())
		return
	}
	SuccessfulApiResponse(c, rb)
}
