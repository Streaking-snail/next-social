package middleware

import (
	"next-social/server/api"
	"next-social/server/common/nt"
	"next-social/server/dto"
	"next-social/server/global/cache"
	"strings"

	"github.com/gin-gonic/gin"
)

var anonymousUrls = []string{"/login"}

func Auth(c *gin.Context) {
	uri := c.Request.RequestURI
	if uri == "/" || strings.HasPrefix(uri, "/#") {
		c.Next()
		return
	}
	// 路由拦截 - 登录身份、资源权限判断等
	for i := range anonymousUrls {
		if strings.HasPrefix(uri, anonymousUrls[i]) {
			c.Next()
			return
		}
	}

	token := api.GetToken(c)
	if token == "" {
		c.Abort()
		api.Fail(c, 401, "您的登录信息已失效，请重新登录后再试。")
		return
	}
	v, found := cache.TokenManager.Get(token)
	if !found {
		c.Abort()
		api.Fail(c, 401, "您的登录信息已失效，请重新登录后再试。")
		return
	}

	authorization := v.(dto.Authorization)

	if strings.EqualFold(nt.LoginToken, authorization.Type) {
		if authorization.Remember {
			// 记住登录有效期两周
			cache.TokenManager.Set(token, authorization, cache.RememberMeExpiration)
		} else {
			cache.TokenManager.Set(token, authorization, cache.NotRememberExpiration)
		}
	}
}

func Admin(c *gin.Context) {

	account, found := api.GetCurrentAccount(c)
	if !found {
		c.Abort()
		api.Fail(c, 401, "您的登录信息已失效，请重新登录后再试。")
		return
	}

	if account.Type != nt.TypeAdmin {
		c.Abort()
		api.Fail(c, 403, "permission denied.")
		return
	}

	c.Next()
	return
}
