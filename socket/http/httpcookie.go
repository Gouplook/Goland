/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  httpcookie
 * @Version: 1.0.0
 * @Date: 2020/12/20 21:35
 */
package http

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	//创建Cookie
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: true,
	}


	//将Cookie发送给浏览器 (方法一）
	w.Header().Set("---Set-Cookie", cookie.String())
	//添加第二个Cookie
	//w.Header().Add("Set-Cookie", cookie2.String())

	//直接调用http的SetCookie函数设置Cookie (方法二 常用）
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

//获取Cookie
func getCookie(w http.ResponseWriter, r *http.Request) {
	//获取请求头中所有的Cookie
	// cookies := r.Header["Cookie"]
	//如果想得到某一个Cookie，可以直接调用Cookie方法
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的Cookie有：", cookie)
}
