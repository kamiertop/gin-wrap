package router

import (
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	return engine
}
