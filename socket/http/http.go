/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  http.go
 * @Version: 1.0.0
 * @Date: 2020/11/28 12:34
 */
package http

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

// ---------------------  链接 http -----------------------
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func WebHttp() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// 添加应答头
		w.Header().Set("Content-type", "application/json")
		w.Header().Set("author", "kangcun.com")

		//获取头部信息 get忽略大小写
		//w.Header().Get("Content-type")

		// 写入的几种方式
		//fmt.Fprintf(w, "hello world HTML!")
		//w.Write([]byte("hello world HTML!"))
		//io.WriteString(w, "hello world HTML!")
		p := User{Id: 123, Name: "Jim"}
		json.NewEncoder(w).Encode(p)

	})

	http.ListenAndServe(":10086", nil)
}

// 利用结构体存放内容
type HandHello struct {
	context string
	name    string
}

func (h *HandHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 将数据写到浏览器上
	w.Write([]byte(h.context))

}

// 第二种写法
func WebHttp2() {
	http.Handle("/hello", &HandHello{context: "webhttp2", name: "Jack"})
	http.ListenAndServe(":10086", nil)
}


func WebHttp3() {
	s := &http.Server{
		Addr: ":10086",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world HTML333!"))
	})
	// 利用server结构体
	s.ListenAndServe()
}

// 利用postman 传入json格式的，写到网页中
func WebHttp4() {
	http.HandleFunc("/user/add", func(writer http.ResponseWriter, request *http.Request) {
		var params map[string]string
		decoder := json.NewDecoder(request.Body)
		decoder.Decode(&params)
		io.WriteString(writer, "post json: "+params["name"])
	})
	http.ListenAndServe(":10086", nil)
}

// 利用postman 传入x-www-form-urlencoded 格式的，写到网页中
func WebHttp5() {
	http.HandleFunc("/user/del", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()
		io.WriteString(writer, "from:"+request.Form.Get("name"))
	})
	http.ListenAndServe(":10086", nil)
}

// --------------------- 请求 ------------------------
func Get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	context, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(context))
}

func Post() {
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()
	context, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(context))
}

func Put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)

	if err != nil {
		panic(err)
	}
	defer func() {
		_ = request.Body.Close()
	}()
	r, err := http.DefaultClient.Do(request) //enter
	if err != nil {
		panic(err)
	}
	context, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", context)
}

// --------------------- 编码问题 ------------------------

func Encoding() {
	r, err := http.Get("https://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = r.Body.Close()
	}()

	//可以通过网页的头部测试编码信息
	bufReader := bufio.NewReader(r.Body)
	bytes, _ := bufReader.Peek(1024) //peek 不会移动只读取位置
	e, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))
	fmt.Println(e)
	//
	bodyReader := transform.NewReader(bufReader, e.NewDecoder())
	content, _ := ioutil.ReadAll(bodyReader)
	fmt.Printf("%s", content)
}

// --------------http 提交POSt form和json数据-------------
