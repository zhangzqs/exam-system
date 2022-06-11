package controller

import "github.com/gin-gonic/gin"

func errorApiResponse(c *gin.Context, errorCode int, errorMsg string) {
	c.JSON(200, ApiResponse{
		Code: errorCode,
		Msg:  errorMsg,
		Data: nil,
	})
}

const (
	SuccessfulCode              = 0
	RequestFormatErrorCode      = 1
	RequestContentErrorCode     = 2
	LoginErrorCode              = 3
	RegisterUserExistsErrorCode = 4
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessfulApiResponse(c *gin.Context, data any) {
	c.JSON(200, ApiResponse{
		Code: SuccessfulCode,
		Msg:  "Successful",
		Data: data,
	})
}

func RequestFormatError(c *gin.Context) {
	errorApiResponse(c, RequestFormatErrorCode, "请求格式有误")
}

func RequestContentError(c *gin.Context, msg string) {
	errorApiResponse(c, RequestContentErrorCode, "请求内容有误："+msg)

}

func LoginError(c *gin.Context) {
	errorApiResponse(c, LoginErrorCode, "用户名或密码有误")
}

func RegisterUserExistsError(c *gin.Context) {
	errorApiResponse(c, RegisterUserExistsErrorCode, "待注册用户已存在")
}
