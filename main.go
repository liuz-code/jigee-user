package main

import (
	"github.com/gin-gonic/gin"
)

const (
	API_VERSION = "/api/v1"
)

func main() {
	r := gin.Default()
	g := r.Group(API_VERSION)
	httpHandlers := getHandlers()
	for _, h := range httpHandlers {
		g.Handle(h.HttpMethod, h.RelativePath, h.Handlers...)
	}
	r.Run(":8000")
}
