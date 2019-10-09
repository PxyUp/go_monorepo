package router

import (
	"github.com/gin-gonic/gin"
	"go_monogrepo/packages/shared/handlers/health"
)

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health("Second app"))
	return r
}