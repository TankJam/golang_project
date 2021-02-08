package main

/*
安装gin框架包:
	- 更换国内源
		- go env -w GO111MODULE=on
		- go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
	- 安装gin包
		- go get -u github.com/gin-gonic/gin

	- 编译go文件:
		- go build:
			- 报错 --> go: cannot find main module; see 'go help modules'
			- 因为缺少依赖包，无法加载main包，所以需要mod文件

		- go mod init

*/
import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context){
	// c.JSON：返回JSON格式的数据
	c.JSON(200, gin.H{
		"message": "hello tank!",
	})
}

func main() {
	// 1.创建一个默认的路由引擎
	r := gin.Default()
	// 2.GET: 请求方式
	// 当客户端以get方法请求/hello路径时，会执行后面sayHello函数
	r.GET("/hello", sayHello)
	// 3.启动HTTP服务, 默认在 0.0.0.0:8080 启动服务
	r.Run(":9090")
	// 4.访问路径: http://127.0.0.1:9090/hello
}


