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

	// 耗时时间
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

// 7.（1）定义一个map存下面数据
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
	return (cycleNum_010(N-1,k)+k-1)%N+1

}

// 11：

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
	x := cycleNum_010(5,3)
	fmt.Println(x)
}
