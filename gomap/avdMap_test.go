/**
 * @Author: yinjinlin
 * @File:  avdMap_test
 * @Description:
 * @Date: 2021/6/17 上午10:16
 */

package gomap

import (
	"fmt"
	"testing"
)

// map和map切片测试
func TestAdvanceMap(t *testing.T) {
	AdvanceMap()
}

func TestSileIn(t *testing.T) {
	SileIn()
}

// 双map测试
func TestAdvMapMap(t *testing.T) {
	AdvMapMap()
}

func TestSlieIn2(t *testing.T) {
	SlieIn2()
}

func TestMapMap(t *testing.T) {
	MapMap2()
}

func TestMapSplitToStruct2(t *testing.T) {
	// MapSplitToStruct2(CardIcad{})
	fmt.Println(15 * 75)
	fmt.Println(100.3*550 + 1125)                              // 56290
	fmt.Println(125.7 * 750)                                   // 95625
	fmt.Println(125.7*750 - 4*550 - 100.3*550 - (95625 * 0.1)) // 28697.5
	fmt.Println(125.7*750 - 4*550 - 100.3*550)                 // 36910

	// expenses
	expenses := 2000/5 + 500/5 + 1200 // 固定成本
	L := 20 * 13 * 20                 // 个 量
	LS := 20 * 6 * 20                 // 个人收入
	L3D := 20*1000 - 20*400
	L3D1 := 20*1000 - 20*350

	// 11


	fmt.Println("费用=", expenses)    // 1700
	fmt.Println("L", L)              // 5200
	fmt.Println("Ls", LS)            //
	fmt.Println("G", L-LS-expenses)  // 2D利润
	fmt.Println("3D输入", 20*400)     // 8000
	fmt.Println("L3D", L3D-expenses) // 3D利润
	fmt.Println("L3D-1",L3D1-expenses) //
	fmt.Println("L", 20*13*20) // 5200
	fmt.Println("3D", 20*1000) // 20000

	fmt.Println()
}
