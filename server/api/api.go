package api

import (
	"next-social/server/common/maps"
	"next-social/server/common/nt"
	"next-social/server/dto"
	"next-social/server/global/cache"
	"next-social/server/model"
	"github.com/gin-gonic/gin"
)


func Fail(c *gin.Context, code int, message string) error {
	return c.JSON(200, maps.Map{
		"code":    code,
		"message": message,
	})
}

func FailWithData(c *gin.Context, code int, message string, data interface{}) error {
	return c.JSON(200, maps.Map{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func Success(c *gin.Context, data interface{}) error {
	return c.JSON(200, maps.Map{
		"code":    1,
		"message": "success",
		"data":    data,
	})
}

func GetToken(c *gin.Context) string {
	token := c.Request().Header.Get(nt.Token)
	if len(token) > 0 {
		return token
	}
	return c.QueryParam(nt.Token)
}

func GetCurrentAccount(c *gin.Context) (*model.User, bool) {
	token := GetToken(c)
	get, b := cache.TokenManager.Get(token)
	if b {
		return get.(dto.Authorization).User, true
	}
	return nil, false
}
