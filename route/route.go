package route

import (
	"github.com/gin-gonic/gin"
	"../webcontroller"
	"net/http"
)

func Routerlist(router *gin.Engine) {
	router.GET("/routergroup", webcontroller.Routerhandle)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"pong"})
	})
	//get字段
	router.GET("/login/:name", func(context *gin.Context) {
		name :=context.Param("name")
		context.String(http.StatusOK,"hello %s",name)
	})

	router.GET("/cttest",webcontroller.Posthandle)



}
