/**
 * @Author: yinjinlin
 * @File:  ClassicalAlogorithm
 * @Description:
 * @Date: 2021/8/25 下午2:44
 */

package main

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// 1： 9X9算法口诀表及耗时
func formulaList9X9_001() {
	starTime := time.Now() // 开始时间
	for i := 0; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d =%2d ", i, j, i*j)
		}
		fmt.Println()
	}
	tc := time.Since(starTime) // 表示经过时间

	fmt.Println("Time consuming: ", tc)
}

// 2: 求两个数的求最大公约数和最小公倍数
//    最小公约数 = a*b / 最大公约数
func getMaximumCommonDivisor_002(a, b int) int {

	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}

	return a

}

// 3：回文数的判断
// 回文数的概念：即是给定一个数，这个数顺读和逆读都是一样的。例如：121，1221是回文数，123，1231不是回文数。
func palindrome_003(s string) bool {

	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		} else {
			j--
			continue
		}
	}

	return true
}

// 4: 求水仙花数
// 水仙花数是指一个 3 位数，它的每个位上的数字的 3次幂之和等于它本身（例如：1^3 + 5^3+ 3^3 = 153）
func isDaffodilNumber(num int64) bool {
	numStru := strconv.FormatInt(num, 10)
	digit := len(numStru) // 位数不确定性
	fmt.Println(digit)
	//
	a := num / 100
	b := (num / 10) % 10
	c := num % 10

	if num == a*a*a+b*b*b+c*c*c {
		fmt.Println("Num = 是水仙花数", num)
		return true
	} else {
		return false
	}
}

// 5：求1-10000之间的同构数
func lsomorphicNumber_005() {
	var k, j int
	k = 10
	for i := 1; i <= 1000; i++ {
		if i == k {
			k *= 10
		}
		j = i * i
		if j%k == i {
			fmt.Printf("%d是同构数，%d的平方是%d\n", i, i, j)
		}
		//
	}

}

// 6：(1)根据工龄(整数)给员工涨工资(整数),工龄和基本工资通过键盘录入
// (2)涨工资的条件如下：
// [10-15) +5000
// [5-10) +2500
// [3~5) +1000
// [1~3) +500
// [0~1) +200
// (3)如果用户输入的工龄为10，基本工资为3000，程序运行后打印格式"您目前工作了10年，基本工资为 3000元,
// 应涨工资 5000元,涨后工资 8000元"
func salary_006(salaryNum float64) {
	baseSalary := 3000.0
	var totalSalary float64

	switch {
	case salaryNum >= 0.0 && salaryNum < 1.0:
		totalSalary = baseSalary + 200
		fmt.Printf("您目前工作了 %.1f 年,基本工资为 3000元,应涨工资 %d元,涨后工资 %.1f元", salaryNum, 200, totalSalary)
	case salaryNum >= 1.0 && salaryNum < 3.0:
		totalSalary = baseSalary + 500
		fmt.Printf("您目前工作了 %.1f 年,基本工资为 3000元,应涨工资 %d元,涨后工资 %.1f元", salaryNum, 500, totalSalary)
	case salaryNum >= 3.0 && salaryNum < 5.0:
		totalSalary = baseSalary + 200
		fmt.Printf("您目前工作了 %.1f 年,基本工资为 3000元,应涨工资 %d元,涨后工资 %.1f元", salaryNum, 1000, totalSalary)
	case salaryNum >= 5.0 && salaryNum < 10.0:
		totalSalary = baseSalary + 200
		fmt.Printf("您目前工作了 %.1f 年,基本工资为 3000元,应涨工资 %d元,涨后工资 %.1f元", salaryNum, 2500, totalSalary)
	case salaryNum >= 10.0 && salaryNum < 15.0:
		totalSalary = baseSalary + 5000
		fmt.Printf("您目前工作了 %.1f 年,基本工资为 3000元,应涨工资 %d元,涨后工资 %.1f元", salaryNum, 5000, totalSalary)
	default:
		fmt.Println("输入工龄有误....")
	}

}

// 7.（1）定义一个map存放下面数据
//       France 首都是 巴黎
//       Italy 首都是 罗马
// 		 Japan 首都是 东京
// 		 India 首都是 新德里
//   （2）检测American 的首都是否存在

func map_007() {

	cityMaps := make(map[string]string)
	cityMaps["France"] = "巴黎"
	cityMaps["Italy"] = "罗马"
	cityMaps["Japan"] = "东京"
	cityMaps["India"] = "新德里"

	// 判断map中的key是否存在
	if _, ok := cityMaps["American"]; ok {
		fmt.Println("American capital is", cityMaps["American"])
	} else {
		fmt.Println("Americal capital is not !")
	}
}

// 8:判断两个map是否拥有相同的键和值
func isMapValueEquality_008() {
	days := make(map[string]string)
	mons := make(map[string]string)

	days["Monday"] = "星期一"
	days["Tuesday"] = "星期二"
	days["Wednesday"] = "星期三"
	days["Thursday"] = "星期四"
	days["Friday"] = "星期五"
	days["Saturday"] = "星期六"
	days["Sunday"] = "星期日"
	days["22"] = "星期日22"

	mons["22"] = "星期日22"
	mons["January"] = "1月"
	mons["February"] = "2月"
	mons["March"] = "3月"
	mons["April"] = "4月"
	mons["May"] = "5月"
	mons["June"] = "6月"
	mons["July"] = "7月"
	mons["August"] = "8月"
	mons["September"] = "9月"
	mons["October"] = "10月"
	mons["November"] = "11月"
	mons["December"] = "12月"

	b := false
	for monkey, mon := range mons {
		for dayKey, day := range days {
			if monkey == dayKey && mon == day {
				fmt.Println(days[dayKey], mons[monkey])
				break
			} else {
				b = true
			}
		}
	}
	if b {
		fmt.Println("两个Map不存在")
	}

}

// 9： 定义一个map，存1到20的阶乘并顺序输出
func factorial_009() {
	m := make(map[int]int)
	for i := 0; i <= 20; i++ {
		if i == 0 {
			m[i] = 1
		} else {
			m[i] = m[i-1] * i
		}
		// fmt.Println(i,"的阶乘是",m[i])
	}
	//
	s := make([]int, 0)
	for k, _ := range m {
		s = append(s, k)
	}
	sort.Ints(s)
	for i := 0; i <= len(s)-1; i++ {
		fmt.Println(i, "的阶乘是", m[i])
	}

}

// 10: 编号为 1-N 的 N 个士兵围坐在一起形成一个圆圈，从编号为 1 的士兵开始依次报数（1，2，3…这样依次报），
//     数到 k 的 士兵会被杀掉出列，之后的士兵再从 1 开始报数。直到最后剩下一士兵，求这个士兵的编号。
func cycleNum_010(N int, k int) int {
	//
	if N == 1 {
		return k
	}
	return (cycleNum_010(N-1, k)+k-1)%N + 1
}

// 11：编写一个函数就地反转一个整型slice中的元素
func reversal_011(array []int) {

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	for j := 0; j < len(array); j++ {
		fmt.Printf("%d,", array[j])
		// fmt.Println(array)
	}

}

// 12 ： 斐波那契数列指的是这样一个数列 1, 1, 2, 3, 5, 8, 13, 21, 34, 55…
// 这个数列从第3项开始，每一项都等于前两项之和

func isFobonacc_012(n int) int {
	if n <= 1 {
		return 1
	}

	return isFobonacc_012(n-1) + isFobonacc_012(n-2)

}

// 13: 卖家将养的一缸金鱼分五次出售，
//     第一次出售上一次卖出全部的一半加二分之一条；
//     第二次卖出余下的三分之一加三分之一条；
//     第三次卖出余下的四分之一加四分之一条；
//     第四次卖出余下的五分之一加五分之一条；
//     最后卖出余下的11条。问原来的鱼缸中共有几条金鱼?
//     ((x/2 + 1) /3 + 1)/4 + 1)

func salefish_013() {
	res := 11
	for j := 4; j >= 1; j-- {
		res = (res*(j+1) + 1) / j
	}
	fmt.Println(res)
}

// 14：从数据库中筛选多条数据 存放[]map[string]interface{} 中，
//     从[]map[string]interface{}中 筛选出key 信息 存放到map[string]interface{}
//     在map[string]interface{} 找出需要的key信息，和在另一个数据查找出的数据进行比对。
//     筛选出新的

// 15： golang 实现单链表（增删改查）

type studentNode struct {
	no   int // 节点编号
	name string
	sex  string
	age  int
	next *studentNode

}

// list后面插入一个数据
func insertStudentNode(head *studentNode, newNode *studentNode) {
	temp := head

	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = newNode
}

// 删除链表
func delStudentNode(head *studentNode, id int) {

	temp := head
	flag := false

	for {
		if temp.next == nil {
			break // 说明遍历到最后一个了
		}
		if temp.next.no == id { // 找到需要删除的节点
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	}else {
		fmt.Println("删除的id不存在", id)
	}

}
// 显示链表
func listStudentNode (head *studentNode) {

	temp := head
	if temp.next == nil  {
		fmt.Println("This is empty list...")
		return
	}

	for {
		fmt.Printf("[%d,%s,%d,%s]",temp.next.no, temp.next.name,temp.next.age,temp.next.sex)

		temp  = temp.next
		if temp.next == nil {
			break // 表示list 遍历多结尾了。
		}
	}
	fmt.Println()

}
func main() {
	// 1 2 3 4 5 6
	// x := palindrome_003("126")
	// fmt.Println(x)
	// isDaffodilNumber(153)
	// lsomorphicNumber_005()
	// salary_006(40.5)
	// map_007()
	// isMapValueEquality_008()
	// factorial_009()
	// x := cycleNum_010(5,3)
	// fmt.Println(x)

	// slice := []int{1, 10, 48, 52}
	// reversal_011(slice)
	// for i := 0; i < 8; i++ {
	// 	fmt.Printf(" %d,", isFobonacc_012(i))
	// }

	// salefish_013()

	// head := &studentNode{}
	// subLisa := &studentNode{
	// 	no: 1,
	// 	name: "lisa",
	// 	sex: "W",
	// 	age: 32,
	// }
	// subJim := &studentNode{
	// 	no: 2,
	// 	name: "Jim",
	// 	sex: "M",
	// 	age: 29,
	// }
	// subTom := &studentNode{
	// 	no: 3,
	// 	name: "Tom",
	// 	sex: "W",
	// 	age:19,
	// }
	// subJack := &studentNode{
	// 	no:4,
	// 	name: "Jack",
	// 	sex: "M",
	// 	age:45,
	// }
	//
	// insertStudentNode(head,subLisa)
	// insertStudentNode(head,subJim)
	// insertStudentNode(head,subTom)
	// insertStudentNode(head,subJack)
	// fmt.Println("dele before :")
	// listStudentNode(head)
	// delStudentNode(head,2)
	// fmt.Println("dele after :")
	// listStudentNode(head)

}
