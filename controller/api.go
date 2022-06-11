package controller

import "github.com/gin-gonic/gin"

func errorApiResponse(c *gin.Context, errorCode int, errorMsg string) {
	r := ApiResponse{
		Code: errorCode,
		Msg:  errorMsg,
		Data: nil,
	}
	c.JSON(200, r)
}

const (
	SuccessfulCode = iota
	RequestFormatErrorCode
	LoginErrorCode
	RegisterUserExistsErrorCode
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessfulApiResponse(c *gin.Context, data any) {
	r := ApiResponse{
		Code: SuccessfulCode,
		Msg:  "Successful",
		Data: data,
	}
	c.JSON(200, r)
}

func RequestFormatError(c *gin.Context) {
	errorApiResponse(c, RequestFormatErrorCode, "请求格式有误")
}

func LoginError(c *gin.Context) {
	errorApiResponse(c, LoginErrorCode, "用户名或密码有误")
}

func RegisterUserExistsError(c *gin.Context) {
	errorApiResponse(c, RegisterUserExistsErrorCode, "待注册用户已存在")
}
