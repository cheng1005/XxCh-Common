package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 最简响应结构
type Response struct {
	Code int         `json:"code"` // 状态码，0表示成功
	Msg  string      `json:"msg"`  // 消息
	Data interface{} `json:"data"` // 数据
}

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: 0, // 0 表示成功
		Msg:  "success",
		Data: data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: code, // 非0表示错误
		Msg:  msg,
		Data: nil,
	})
}

// ErrorWithData 带数据的错误响应
func ErrorWithData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
