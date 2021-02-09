package main
/*
safe，防止xss攻击
*/
import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed1,err:%v\n", err)
		return
	}

	name := "小泽玛利亚"

	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("parse template failed2,err:%v\n", err)
		return
	}
}

// xss攻击，修改请求返回的结果，执行js弹出窗口
func xss(w http.ResponseWriter, r *http.Request){
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string)template.HTML{
			return template.HTML(str)

		},
	}).ParseFiles("./xss.tmpl")

	if err != nil {
		fmt.Printf("parse template failed2,err:%v\n", err)
		return
	}

	// 模板渲染
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='http://liwenzhou.com'> 哪吒</a>"

	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})

}

func main() {
	// 1.配置路由
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)

	// 2.监听服务
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err:%v", err)
		return
	}
}
