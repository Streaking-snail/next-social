package api

import (
	"fmt"
	"net/http"
	"next-social/server/common/nt"
	"next-social/server/dto"
	"next-social/server/global/cache"
	"next-social/server/model"

	"github.com/gin-gonic/gin"
)

func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
	})
	// return c.JSON(200, maps.Map{
	// 	"code":    code,
	// 	"message": message,
	// })
}

func FailWithData(c *gin.Context, code int, message string, data interface{}) {
	// return c.JSON(200, maps.Map{
	// 	"code":    code,
	// 	"message": message,
	// 	"data":    data,
	// })
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func ShowError(c *gin.Context, msg interface{}) {
	msg = fmt.Sprintf("%v", msg)
	c.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    data,
	})
}

func GetToken(c *gin.Context) string {
	//token := c.Request().Header.Get(nt.Token)
	token := c.GetHeader(nt.Token)
	if len(token) > 0 {
		return token
	}
	return c.Param(nt.Token)
}

func GetCurrentAccount(c *gin.Context) (*model.User, bool) {
	token := GetToken(c)
	get, b := cache.TokenManager.Get(token)
	if b {
		return get.(dto.Authorization).User, true
	}
	return nil, false
}
