/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/20 14:46
@Description: 公共方法

*********************************************/
package functions

import "fmt"

func GetFileNameByHash(hash string)string{
	return  fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s",hash[0:3],hash[3:6],hash[6:9],hash[9:12],hash[12:15],hash[15:18],hash[18:])
}

func GetFileName (name string)string{
	if len(name) > 64 {
		return name[0:64]
	}
	return name
}
