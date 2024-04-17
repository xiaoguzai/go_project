package router

import (
	"gimchat1/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetIndex)
	return r
}
