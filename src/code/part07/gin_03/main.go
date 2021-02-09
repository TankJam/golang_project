package main
/*
模板继承

父页面:
	base.html 中的 {{block "content" .}}{{end}}
	可以声明子html 通过content参数来渲染，实现继承base页面。

子页面:
	- 继承根模板
		{{template "base.tmpl" .}}
	- 重新定义块模板
		{{define "content"}}
		<h1>这是index2页面</h1>
		<p>Hello {{ . }}</p>
		{{end}}
*/

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request){
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil{
		fmt.Println("parse template failed,err:%v\n", err)
		return
	}
	msg := "暴走大事件之张全蛋参演新喜剧之王"

	// 渲染模板
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request){
	// 定义模板
	// 解析模板
	t,err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, error:%v\n")
	}
	msg := "赵铁柱"
	// 渲染模板
	t.Execute(w, msg)
}

func index2(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:%v\n", err)
		return
	}
	name := "张全蛋9527"
	t.ExecuteTemplate(w, "index2.tmpl", name)
}

func home2(w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Println("parse template failed, err:%v\n", err)
	}
	name := "张全蛋9527"
	t.ExecuteTemplate(w, "home2.tmpl", name)
}

func main() {

	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)

	err:= http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http server start failed, error: %v\n", err)
		return
	}


}
