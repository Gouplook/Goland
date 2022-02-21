/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/20 15:05
@Description:

*********************************************/
package file

import (
	"fmt"
	"testing"
)

// 根据url地址获取远程图片
func TestSaveImgFromUrl(t *testing.T) {
	// u= 33d4385f-6021-40c7-a90f-04b21f6e6049
	// u := uuid.NewV4()
	// hash := strings.Replace(u.String(), "-","",-1)
	// u2 := functions.GetFileNameByHash(hash)
	// fmt.Println(u)
	// fmt.Println(u2)
	// filename := ".jpg"
	// path := filepath.Join("./upload/image",filename)
	// fmt.Println(path)

	//
	url := "https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fpic1.win4000.com%2Fwallpaper%2Fa%2F583654178e0cf.jpg%3Fdown&refer=http%3A%2F%2Fpic1.win4000.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=jpeg?sec=1616400817&t=9e12e9bd96e909505fc452402857254f"
	SaveImgFromUrl(url)

	fmt.Println()
}
