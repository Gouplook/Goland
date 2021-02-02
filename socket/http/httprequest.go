/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  httprequest
 * @Version: 1.0.0
 * @Date: 2020/12/19 16:54
 */
package http

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息有：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))

	//获取请求体重内容的长度
	len := r.ContentLength
	body := make([]byte,len)
	//将请求体中的内容读到body中
	r.Body.Read(body)
	//在浏览器中显示请求体中的内容
	fmt.Fprintln(w, "请求体中的内容有：", string(body))

	//解析表单，在调用r.Form之前必须执行该操作
	r.ParseForm()
	//获取请求参数
	//如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	//那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面
	fmt.Fprintln(w, "请求参数有：", r.Form)
	fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)
	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("username"))
}

// 请求
func HttpRequest(){
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":10086",nil)
}


