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
	dec := decimal.NewFromFloat(123.66).Sub(decimal.NewFromFloat(22))

	fmt.Println(dec)
}