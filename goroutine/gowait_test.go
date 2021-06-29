/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/22 17:19
@Description:

*********************************************/
package goroutine

import (
	"fmt"
	"github.com/shopspring/decimal"
	"testing"
)

func TestGoWait(t *testing.T) {
	// GoWait()
	amount2020 := 218000 // 2020年高考总人数
	amount2021 := 191000 // 2021年高考总人数
	rank2020 := 65957
	rank2021 := decimal.NewFromInt(int64(amount2021)).Mul(decimal.NewFromInt(int64(rank2020))).Div(decimal.NewFromInt(int64(amount2020)))

	// 理论数据 57788 实际64838
	fmt.Println("=======数据分析结果=========")
	fmt.Println("2021年 整体数据技术理论数据排名", rank2021)

	gradeNumber2020 := 132098 //  2020年150分以上人数（理科）
	gradeNumber2021 := 102737 //  2020年150分以上人数（理科）
	ranking2020 := 65957      //  446在2020年平行排名

	ranking2021 := decimal.NewFromInt(int64(ranking2020)).Mul(decimal.NewFromInt(int64(gradeNumber2021))).Div(decimal.NewFromInt(int64(gradeNumber2020)))
	fmt.Println("2021年446分相对于2020年理论排名：", ranking2021)

	college2020 := 101844 // 2020年本科达线人数 本控线：359分
	college2021 := 93737  // 2021年本科达线人数 本控线：336分

	ranking2021 = decimal.NewFromInt(int64(college2021)).Mul(decimal.NewFromInt(int64(ranking2020))).Div(decimal.NewFromInt(int64(college2020)))
	fmt.Println("2021年本科控制线446分相对于2020年理论排名：", ranking2021)


}
