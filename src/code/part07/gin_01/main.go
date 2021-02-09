package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct{
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request){
	//	// 定义模板
	//	// 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil{
		fmt.Println("parse template failed, err:%v", err)
		return
	}

	// 模板渲染
	// 获取一个User对象
	u1 := User{
		Name: "张全蛋",
		Gender: "女",
		Age: 18,
	}
	// 创建一个map类型
	m1 := map[string]interface{}{
		"name": "张全蛋2",
		"gender": "man",
		"age": 18,
	}
	// 创建一个数组类型
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}

	// 将三种数据类型，装进map中，通过response对象返回给访问者
	t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
		"hobby": hobbyList,
	})

}

func main() {
	// 监听路由

	// 跟路径，request请求到sayHello方法中
	http.HandleFunc("/", sayHello)
	fmt.Println("server is run, click go to http://127.0.0.1:9000/")
	// 监听服务端口，127.0.0.1:9000
	err := http.ListenAndServe(":9000", nil)
	if err != nil{
		fmt.Println("HTTP SERVER START FAILED, ERROR:", err)
	}
}
