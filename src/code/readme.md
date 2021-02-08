# GO study

```go
// go环境配置
/// 用户环境
GOPATH ---> Go语言开发目录
PATH ---> Go语言开发查找编译后的文件目录

/// 系统环境
GOROOT ---> Go语言编译器的目录
PATH ---> GO编译器的bin目录

// 快速导入包工具
ctrl + shirt + x
Auto Import
// 自动格式化
Prettier - Code formatter




// 安装依赖工具，若无法安装
/// 1.开启代理设置，Go 1.13 以上默认启用，可跳过此步
go env -w GO111MODULE=on  // 该设置需要重新设置会off

/// 2.设置代理
go env -w GOPROXY=https://goproxy.io,direct

/// 3.重启VScode再安装工具，享受SUCCEEDED的提示吧
ctrl + shirt + p
install all

// go build  编译当前目录下的go文件
// go build -o hello.exe  编译当前目录下的go文件
// go install 安装，先编译当前目录下的文件，然后放到bin目录下

// 若想windows下编写的代码可以在linux平台中使用
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

// 配置自动补全代码
ctrl + shirt + p, 输入snippets，然后选择go
"print": {
		"prefix": "pln",  // 快捷键
		"body": "fmt.Println($0)",  // $0光标选中
		"description": "打印行..."
	}
```