/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  gomain.go
 * @Version: 1.0.0
 * @Date: 2021/7/22 22:36
 */
package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)
func read(path string) ([]byte,error)  {
	return ioutil.ReadFile(path)
}
// func main()  {
// 	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
// 		writer.Header().Set("Content-Type","application/pdf")
// 		content,err := read("./goweb/0218.pdf")
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		writer.Write(content)
// 	})
// 	log.Fatal(http.ListenAndServe(":8089",nil))
// }

var (
	url = "https://file.900sui.cn/e668a530c9de452c9b10cf7adfd69eb1.pdf"
)
func main(){


	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f,err := os.Create("e668a530c9de452c9b10cf7adfd69eb1.pdf")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)



}