package Router

import (
	"Mmx/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(){
	G:=gin.Default()

	go Middlewares.Sec.InitIpLogger() //启动频次记录协程
	G.Use(Middlewares.Sec.Main) //安全中间件
}
