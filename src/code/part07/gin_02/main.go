package main
/*
模板嵌套
*/
import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request){
	// 定义一个函数kua
	// 要么只有一个返回值，要么有两个返回值，第二个返回值必须是error类型
	k := func(name string) (string, error){
		return name + "年轻又帅气，18岁波波脆！！！", nil
	}

	// 定义模板
	// 创建一个名字是f的模板对象, 名字一定要与模板的名字能对应上
	t := template.New("f.tmpl")
	// FuncMap 模板对象自带的map类型
	t.Funcs(template.FuncMap{
		"kua99": k,
	})
	// 解析模板
	_, err := t.ParseFiles("./f.tmpl")
	if err!= nil{
		fmt.Printf("parse template failed, err:#{err}\n")
		return
	}

	// 渲染模板
	name := "张全蛋"
	t.Execute(w, name)  // name 传递给了 func匿名函数
}

func demo1(w http.ResponseWriter, r *http.Request){
	// 定义模板
	// 解析模板
	// 可一次性解析多个模板
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		fmt.Println("http server start failed, err: %v", err)
		return
	}
	name := "张全蛋"
	// 渲染模板
	t.Execute(w, name)
}

func main() {
	// 配置路由
	// 根目录
	http.HandleFunc("/", f1)

	// tmplDemo路径
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err: %v", err)
	}
}
