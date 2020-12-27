package Router

import (
	"Mmx/Controllers"
	"Mmx/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(){
	G:=gin.Default()

	go Middlewares.Sec.InitIpLogger() //启动频次记录协程
	G.Use(Middlewares.Sec.Main) //安全中间件

	G.GET("/download",Controllers.Download)//单文件分发接口

	//HTML页渲染
	G.LoadHTMLGlob("templates/*.html")
	G.Static("/files","templates/files")

	G.GET("/hi/:token",Controllers.Index)//主页-文件列出

	G.Run(":10080")
}
