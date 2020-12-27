package Controllers

import (
	"Mmx/Modles"
	"github.com/gin-gonic/gin"
)

type downloadFORM struct {
	Dir string `form:"signkey"`
	Name string `form:"name"`
}

func Download(c *gin.Context){
	var form downloadFORM
	if c.ShouldBind(&form)!=nil{
		c.Status(501)
		c.Abort()
		return
	}
	Path:="FILE/"+form.Dir+"/"+form.Name
	if !Modles.File.Exists("FILE/"+form.Dir+"/"+form.Name){//不存在
		c.Status(404)
		c.Abort()
		return
	}
	if Modles.File.IsDir(Path){//不支持文件夹
		c.Status(404)
		c.Abort()
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+form.Name)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	c.File(Path)
}
