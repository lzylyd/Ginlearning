package main

import (
	"./route"
	"github.com/gin-gonic/gin"
)

func main(){
	//默认配置路由 没有使用中间件
	r:=gin.Default()

	//使用route里的routerlist方法加载路由
	route.Routerlist(r)
	r.Run()
	//http.ListenAndServe(":8005",r)
}