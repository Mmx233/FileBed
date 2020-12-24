package Secure

import (
	"Mmx/Public"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	IpLogger map[string]int
	IpChan  = make(chan string)
)

func InitIpLogger(){
	for{
		ip:=<-IpChan
		if IpLogger[ip]>=0 {//配合封禁
			IpLogger[ip]++
			go func(){//仅记录五秒内的访问
				time.Sleep(time.Second*5)
				IpLogger[ip]--
				if IpLogger[ip]==0 {
					delete(IpLogger, ip)
				}
			}()
		}
	}
}

func Sec(c *gin.Context){
	//防扫描
	IpChan<-c.ClientIP()
	if IpLogger[c.ClientIP()]>5 || IpLogger[c.ClientIP()]<0 {//五秒内最多五次访问
		c.AsciiJSON(403,Public.CallErrorWithCode(1))
		c.Abort()
		return
	}else if IpLogger[c.ClientIP()]>=50{//每秒超十次封禁Ip
		IpLogger[c.ClientIP()]=-1//使被拦截
		go func(ip string){//截除拦截
			time.Sleep(time.Hour)//一小时后截除
			delete(IpLogger,ip)
		}(c.ClientIP())
		c.AsciiJSON(403,Public.CallErrorWithCode(1))
		c.Abort()
		return
	}
	c.Next()
}
