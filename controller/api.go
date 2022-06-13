package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func errorApiResponse(c *gin.Context, errorCode int, errorMsg ...any) {

	c.JSON(200, ApiResponse{
		Code: errorCode,
		Msg:  fmt.Sprintln(errorMsg...),
		Data: nil,
	})
}

const (
	SuccessfulCode              = 0
	RequestFormatErrorCode      = 1
	RequestContentErrorCode     = 2
	LoginErrorCode              = 3
	RegisterUserExistsErrorCode = 4
	OperationNeedLoginErrorCode = 5
	TokenInvalidErrorCode       = 6
	DatabaseErrorCode           = 7
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessfulApiResponse(c *gin.Context, data ...any) {
	if len(data) == 0 {
		c.JSON(200, ApiResponse{
			Code: SuccessfulCode,
			Msg:  "Successful",
			Data: nil,
		})
		return
	}
	if len(data) == 1 {
		c.JSON(200, ApiResponse{
			Code: SuccessfulCode,
			Msg:  "Successful",
			Data: data[0],
		})
		return
	}
	c.JSON(200, ApiResponse{
		Code: SuccessfulCode,
		Msg:  "Successful",
		Data: data,
	})
}

func RequestFormatError(c *gin.Context, errorMsg ...any) {
	errorApiResponse(c, RequestFormatErrorCode, "请求格式有误", errorMsg)
}

func RequestContentError(c *gin.Context, msg string, errorMsg ...any) {
	errorApiResponse(c, RequestContentErrorCode, "请求内容有误：", msg, errorMsg)

}

func LoginError(c *gin.Context, errorMsg ...any) {
	errorApiResponse(c, LoginErrorCode, "用户名或密码有误", errorMsg)
}

func RegisterUserExistsError(c *gin.Context, errorMsg ...any) {
	errorApiResponse(c, RegisterUserExistsErrorCode, "待注册用户已存在", errorMsg)
}

func OperationNeedLoginError(c *gin.Context, errorMsg ...any) {
	errorApiResponse(c, OperationNeedLoginErrorCode, "该操作需要登录", errorMsg)
}

func TokenInvalidError(c *gin.Context, errorMsg ...any) {
	errorApiResponse(c, TokenInvalidErrorCode, "Token已失效，请重新登录", errorMsg)
}

func DatabaseError(c *gin.Context, errorMsg ...any) {
	errorApiResponse(c, DatabaseErrorCode, "数据库异常", errorMsg)
}
