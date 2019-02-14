package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"./webcontroller"
)

func main(){
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{"message":"pong"})
	})
	r.GET("/wang", func(context *gin.Context) {
		context.JSON(200,gin.H{"name":"xin"})
	})
	r.GET("/login/:name", func(context *gin.Context) {
		name :=context.Param("name")
		context.String(http.StatusOK,"hello %s",name)
	})
	r.GET("/cttest",webcontroller.Posthandle)
	r.Run()
	//http.ListenAndServe(":8005",r)
}