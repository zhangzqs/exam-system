package controller

import "github.com/gin-gonic/gin"

var (
	SuccessfulCode   = 0
	UnknownErrorCode = 1
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessfulApiResponse(c *gin.Context, data interface{}) {
	r := ApiResponse{
		Code: SuccessfulCode,
		Msg:  "Successful",
		Data: data,
	}
	c.JSON(200, r)
}

func ErrorApiResponse(c *gin.Context, errorCode int, errorMsg string) {
	r := ApiResponse{
		Code: errorCode,
		Msg:  errorMsg,
		Data: nil,
	}
	c.JSON(200, r)
}

func UnknownErrorApiResponse(c *gin.Context, errorMsg string) {
	ErrorApiResponse(c, UnknownErrorCode, errorMsg)
}
