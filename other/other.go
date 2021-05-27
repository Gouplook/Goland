/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/27 上午10:46

*******************************************/
package other

import "fmt"

// --------- Go ok 	用法总结------------

// 1： ok判断key是否在map中
func IsOk(){
	var nameList = map[string]string{"姓名": "李四", "性别": "男"}
	name, ok := nameList["姓名"] // 假如 key 存在，则 ok = true，否则，ok = false
	if ok {
		fmt.Println(name)
	}
}
// 2： 类型判断
func TypeJudgment(){
	var a interface{}
	a = "yin"
	// 如果 ok 是 true，则说明 变量 a 是字符串类型，而 newA 就是 string 类型的变量，a 的实际值
	newA, ok := a.(string)

	fmt.Println(newA,ok)
}


//3: 判断 gochannel 是否 已关闭 且 有没有数据
func ChannelIsData(){
	ch := make(chan string,1)
	for {
		x, ok := <- ch
		if !ok {
			break // 通道 已关闭 且 没有数据，则跳出循环
		}

		fmt.Println(x)
	}
}