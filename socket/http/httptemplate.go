/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  httptemplate
 * @Version: 1.0.0
 * @Date: 2020/12/19 17:42
 */
package http

import (
	"GoInduction/socket/http/model"
	"html/template"
	"net/http"
)

//创建处理器函数
func templatehandle(w http.ResponseWriter,r *http.Request){
	//解析模板文件
	// t, _ := template.ParseFiles("./httptemplatehtml/index.html")
	//通过Must函数让Go帮我们自动处理异常  多个参数时，最后一个参数生效
	t := template.Must(template.ParseFiles("./httptemplatehtml/index.html", "./httptemplatehtml/index2.html"))

	//执行
	//t.Execute(w, "Hello Template")
	//将响应数据在index2.html文件中显示 对多个tmeplate使用下面这个方法
	// name : specified data object
	t.ExecuteTemplate(w, "index2.html", "跳转index2.html中")
}
func HttpTemplate(){
	http.HandleFunc("/testTemplate",templatehandle)
	http.ListenAndServe(":10086",nil)
}

//----------------
//测试if
func ifHandler(w http.ResponseWriter, r *http.Request){
	//解析模板文件
	t := template.Must(template.ParseFiles("./httptemplatehtml/if.html"))
	age := 17
	//执行
	t.Execute(w, age > 18)
}

//测试range
func rangeHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("./httptemplatehtml/range.html"))
	var emps []*model.Employee
	emp := &model.Employee{
		ID:       1,
		LastName: "jack",
		Email:    "lxl@jnl.com",
	}
	emps = append(emps, emp)
	emp2 := &model.Employee{
		ID:       2,
		LastName: "baidu",
		Email:    "bbh@cyf.com",
	}
	emps = append(emps, emp2)
	emp3 := &model.Employee{
		ID:       3,
		LastName: "wangyi",
		Email:    "mr@wbq.com",
	}
	emps = append(emps, emp3)

	//执行
	t.Execute(w, emps)
}

// 设置动作
//测试with
func withHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("./httptemplatehtml/with.html"))
	//执行
	t.Execute(w, "hello-->")
}

// 包含动作
//测试template
func templateHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("./httptemplatehtml/template1.html", "./httptemplatehtml/template2.html"))
	//执行
	t.Execute(w, "能在两个文件中显示吗？")
}

// 定义动作
//测试define
func defineHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("./httptemplatehtml/define.html"))
	//执行
	t.ExecuteTemplate(w, "model", "")
}
//测试testDefine2
func define2Handler(w http.ResponseWriter, r *http.Request) {
	age := 20
	var t *template.Template
	if age < 18 {
		//解析模板文件
		t = template.Must(template.ParseFiles("./httptemplatehtml/define2.html"))
	} else {
		//解析模板文件
		t = template.Must(template.ParseFiles("./httptemplatehtml/define2.html", "./httptemplatehtml/content1.html"))
	}
	//执行
	t.ExecuteTemplate(w, "model", "")
}

// template action
func TemplateAction(){
	//http.HandleFunc("/testIf", ifHandler)
	//http.HandleFunc("/testRange", rangeHandler)
	//http.HandleFunc("/testWith",withHandler)
	//http.HandleFunc("/template",templateHandler)
	//http.HandleFunc("/testDefine",defineHandler)
	http.HandleFunc("/testdefine2Handler",define2Handler)

	http.ListenAndServe(":10086", nil)

}


