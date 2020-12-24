package Router

import (
	"Mmx/Middlewares/Secure"
	"github.com/gin-gonic/gin"
)

func InitRouter(){
	G:=gin.Default()

	go Secure.InitIpLogger() //启动频次记录函数
	G.Use(Secure.Sec) //安全中间件
}
