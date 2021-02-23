package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {

	r := gin.Default()

	// 加载静态文件
	r.Static("/xxx", "./statics")

	// gin框架中给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})

	// r.LoadHTMLFiles("templates/index.tmpl", "templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		// 发送http协议的请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "liwenzhou.com",
		})
	})

	r.GET("/users/index", func(c *gin.Context){
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<a href='https://liwenzhou.com'>哪吒的博客</a>",
		})
	})

	// 返回从网上下载的模板
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	// 1.请求源码
	r.Run(":9090")
}

