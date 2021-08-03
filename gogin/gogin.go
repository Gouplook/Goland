/**
 * @Author: yinjinlin
 * @File:  gogin.go
 * @Description:
 * @Date: 2021/8/3 下午4:08
 */

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 上传文件，postman需要选择文件类型

func main() {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

		file,_ := c.FormFile("file")
		log.Println(file.Filename)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})


	router.Run(":8082")
}
