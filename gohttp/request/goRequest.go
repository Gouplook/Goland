/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  goRequest
 * @Version: 1.0.0
 * @Date: 2021/7/26 8:25
 */
package main

import (
 "fmt"
 "io/ioutil"
 "net/http"
 "net/url"
)

func printBody(r *http.Response) {
 defer func() {_ = r.Body.Close()}()

 content, err := ioutil.ReadAll(r.Body)
 if err != nil {
  panic(err)
 }

 fmt.Printf("%s", content)
}

func requestByParams() {
 request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
 if err != nil {
  panic(err)
 }

 params := make(url.Values)
 params.Add("name", "poloxue")
 params.Add("age", "18")

 request.URL.RawQuery = params.Encode()

 r , err:= http.DefaultClient.Do(request)
 if err != nil {
  panic(err)
 }
 printBody(r)
}

func requestByHead() {
 request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
 if err != nil {
  panic(err)
 }

 request.Header.Add("user-agent", "chrome")
 r , err:= http.DefaultClient.Do(request)
 if err != nil {
  panic(err)
 }
 printBody(r)
}