package Controllers

import (
	"Mmx/Modles"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
)

func countSize(size int64)string{
	FSize:=float32(size)
	var format string
	var S float32
	switch{
	case size<1024:
		format="B"
	case size<1024*1024:
		S=FSize/(1024)
		format="KB"
	case size<1024*1024*1024:
		S=FSize/(1024*1024)
		format="MB"
	case size<1024*1024*1024*1024:
		S=FSize/(1024*1024*1024)
		format="GB"
	}
	return fmt.Sprintf("%.1f"+format,S)
}

func Index(c *gin.Context){
	dir:="FILE/"+c.Param("token") //获取目录

	if !Modles.File.Exists(dir){//判断是否存在
		c.Status(403)
		c.Abort()
		return
	}

	//读取目录
	var content string
	if rd, err := ioutil.ReadDir(dir + "/");err!=nil{
		c.Status(500)
		c.Abort()
		return
	}else{
		for _,file :=range rd{
			content=content+"<tr><td>"+file.Name()+"</td><td>"+countSize(file.Size())+"</td><td onclick=\"download(this,'"+c.Param("token")+"')\">下载</td></tr>"
		}
	}
	c.HTML(200,"index.html",gin.H{
		"content":template.HTML(content),
	})
}
