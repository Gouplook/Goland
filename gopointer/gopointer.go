/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  gopointer
 * @Version: 1.0.0
 * @Date: 2021/8/15 19:57
 */
package gopointer

import "fmt"

// 数组指针
func ArrayPointer() {

	nums := []int{1, 2, 3, 4, 5, 6, 6, 7, 8}
	var p *[]int
	p = &nums
	fmt.Println(*p)

}
//指针数组
func PointerArray(){
   var p[]*int = make([]*int,2)

   var i int = 10
   var j int = 20
   p[0] = &i
   p[1] = &j
   fmt.Println(p[0])
   fmt.Println(*p[0])
}