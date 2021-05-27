/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/30 15:13
@Description:

*********************************************/
package utils

import (
	"fmt"
	"github.com/shopspring/decimal" // 大数据处理框架
)

func BigDataProcess(){
	add := decimal.NewFromFloat(123.66).Add(decimal.NewFromFloat(22))
	dec := decimal.NewFromFloat(123.66).Sub(decimal.NewFromFloat(22))
	mul := decimal.NewFromFloat(10).Mul (decimal.NewFromFloat(22))
	div := decimal.NewFromFloat(float64(12)).Div(decimal.NewFromFloat(float64(7)))
	div2,b := decimal.NewFromFloat(float64(12)).Div(decimal.NewFromFloat(float64(7))).Truncate(4).Float64()

	fmt.Println(add)
	fmt.Println(dec)
	fmt.Println(mul)
	fmt.Println(div)
	fmt.Println(div2,b)
	fmt.Println(add)

	var RealPrice float64 = 10
	var Price float64 = 50

	// 保留2位有效数字
	// 充值卡： 充10 送50元，  10/（10+50） 面值 60，实际支付 10 元
	discount, _ := decimal.NewFromFloat(RealPrice).Div(decimal.NewFromFloat(Price + RealPrice)).Truncate(1).Float64()
	fmt.Println("dicount= ", discount)
}