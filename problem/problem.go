/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/7 下午3:03

*******************************************/
package problem

import "fmt"

func Problem1(){
	x := 1
	fmt.Println(x)     //prints 1
	{
		fmt.Println(x) //prints 1
		x := 2
		fmt.Println(x) //prints 2
	}
	fmt.Println(x)     //prints 1 (bad if you need 2)
}

func Problem2(){
	data := "A\xfe\x02\xff\x04"
	for _,v := range data {
		fmt.Printf("%#x ",v)
	}
	//prints: 0x41 0xfffd 0x2 0xfffd 0x4 (not ok)

	fmt.Println()
	for _,v := range []byte(data) {
		fmt.Printf("%#x ",v)
	}
	//prints: 0x41 0xfe 0x2 0xff 0x4 (good)
}

func Problem3(){
	m := map[string]int{"one":1,"two":2,"three":3,"four":4}
	for k,v := range m {
		fmt.Println(k,v)
	}
}

