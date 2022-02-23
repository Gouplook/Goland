/**
 * @Author: yinjinlin
 * @File:  enumeration_test.go
 * @Description:
 * @Date: 2021/11/9 上午10:32
 */

package goenumeration

import (
	"fmt"
	"testing"
)

func TestFishType_String(t *testing.T) {
	var f FishType = A

	fmt.Println(f)
	fmt.Println(f)

	fmt.Println(A,B)

	switch f {
	case A:
		fmt.Println("=====A")
	case B:
		fmt.Println("=====B")
	case C:
		fmt.Println("=====C")

	}

	fmt.Println("======D")
	fmt.Println(f.String())

	// 450 + 1000
	//
	// Enmeration()
}


func TestEnmeration(t *testing.T) {

	Enmeration()
	//  10万
	//  3D 7万（100单元） + 3万（3000）
	// 成本 100*160 = 16000+ 15000 = 31000
	// 16000+15000 = 31000   //

	fmt.Println()
}