package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func InitDist() {
	app := gin.Default()
	// 加载静态资源
	app.Use(static.Serve("/", static.LocalFile("dist", true)))
	// 静态文件的文件夹相对目录
	app.StaticFS("/dist", http.Dir("./dist"))
	// 第一个参数是api 第二个参数是具体的文件名字
	app.StaticFile("/favicon.ico", "./favicon.ico")
	app.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})
	app.Run("0.0.0.0:8888")
}
