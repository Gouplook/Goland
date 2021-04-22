/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/20 10:19
@Description:

*********************************************/
package tool

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type CardBase struct {
	Name   string  // 名称
	Price  float64 // 价格
	CardId int
	Clicks int     // 点击量
}

type Card struct {
	CardBase
	Sales int  // 销量
}

// 定义一个map
func AdvanceMap() {
	maps := make([]map[string]interface{}, 3)
	maps[0] = make(map[string]interface{})
	maps[0]["name"] = "综合0"
	maps[0]["card_id"] = 11
	maps[0]["price"] = 100.0
	maps[1] = make(map[string]interface{})
	maps[1]["name"] = "综合1"
	maps[1]["card_id"] = 12
	maps[1]["price"] = 100.1
	maps[2] = make(map[string]interface{})
	maps[2]["name"] = "综合2"
	maps[2]["card_id"] = 13
	maps[2]["price"] = 200.2

	fmt.Println("maps=====", maps)
	var outStruct []CardBase

	_ = mapstructure.WeakDecode(maps, &outStruct)
	//var cardMap = map[string]Card{}





}
