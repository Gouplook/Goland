/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  gomain.go
 * @Version: 1.0.0
 * @Date: 2021/7/22 22:36
 */
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)
func read(path string) ([]byte,error)  {
	return ioutil.ReadFile(path)
}
func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type","application/pdf")
		content,err := read("./goweb/0218.pdf")
		if err != nil {
			log.Println(err.Error())
		}
		writer.Write(content)
	})
	log.Fatal(http.ListenAndServe(":8089",nil))
}