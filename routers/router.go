package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/gofast-bbs/api/v1"
	"github.com/gofast-bbs/config"
)

func InitRouter() {
	gin.SetMode(config.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("test", v1.Test)
	}
	r.Run(":" + config.HttpPort)
}
