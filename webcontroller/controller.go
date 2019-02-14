package webcontroller

import "github.com/gin-gonic/gin"

func Posthandle(ctx *gin.Context) {
	ctx.String(200, "hello")
}

//使用路由文件加载的路由
func Routerhandle(ctx *gin.Context) {
	ctx.String(200, "router test")
}
func RealPostHandle(ctx *gin.Context) {
	ctx.post
}
