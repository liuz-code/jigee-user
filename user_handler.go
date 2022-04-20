package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandleInfo struct {
	RelativePath string
	HttpMethod   string
	Handlers     []gin.HandlerFunc
}

func getHandlers() []*HandleInfo {
	var handlers []*HandleInfo = make([]*HandleInfo, 0)
	postLogin := HandleInfo{
		RelativePath: "/user/login",
		HttpMethod:   http.MethodPost,
		Handlers:     []gin.HandlerFunc{userLogin},
	}
	getLogout := HandleInfo{
		RelativePath: "/user/logout",
		HttpMethod:   http.MethodGet,
		Handlers:     []gin.HandlerFunc{userLogout},
	}
	getTokenCheck := HandleInfo{
		RelativePath: "/token/check",
		HttpMethod:   http.MethodGet,
		Handlers:     []gin.HandlerFunc{tokenCheck},
	}
	return append(handlers, &postLogin, &getLogout, &getTokenCheck)
}

func userLogin(c *gin.Context) {
}

func userLogout(c *gin.Context) {
}

func tokenCheck(c *gin.Context) {
}
