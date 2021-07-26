/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  gobase
 * @Version: 1.0.0
 * @Date: 2021/7/26 8:21
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func post() {
	// body 是请求参数
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) // enter 键
	if err != nil {
		panic(err)
	}

	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func del() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request) // enter 键
	if err != nil {
		panic(err)
	}

	defer func() { _ = r.Body.Close() }()

	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", content)
}

func main() {
	// client request 和 response

	//get()
	put()
	// head
	// options
	//del()
}
