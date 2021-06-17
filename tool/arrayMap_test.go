/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午5:00

*******************************************/
package tool

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestArrayKeys(t *testing.T) {
	ma := map[int]interface{}{
		1: "name",
		2: "age",
	}
	slile := ArrayKeys(ma)
	fmt.Println(slile)
}
func TestArrayString(t *testing.T) {
	tagIds := []string{"1#", "2#", "#4"}
	fmt.Println("tagIds: ", tagIds)
	str := ArrayString(",", tagIds)
	fmt.Println(str)
}

func SortMaps(field string, maps []map[string]interface{}) []map[string]interface{} {

	var tempMap = make(map[string]interface{})
	var keys = make([]string, 0)

	for _, v := range maps {
		vs := v[field]
		if vp, ok := vs.(float64); ok {
			vs = strconv.FormatFloat(vp, 'f', -1, 64)
		}
		if vp, ok := vs.(int); ok {
			vs = strconv.FormatInt(int64(vp), 10)
		}
		if vp, ok := vs.(string); ok {
			vs = vp
		}
		keys = append(keys, vs.(string))
		tempMap[vs.(string)] = v
	}
	sort.Strings(keys)
	remapData := make([]map[string]interface{}, 0)

	// 遍历keys，keys中现在是排好序的, 根据keys中的值进行访问查询
	for _, v := range keys {
		remapData = append(remapData, tempMap[v].(map[string]interface{}))
	}
	return remapData
}
func TestStr(t *testing.T){
	startTimePointArr := strings.Split("22:00", ":")
	startTimePointStr := startTimePointArr[0] + startTimePointArr[1]
	startTimePointInt, _ := strconv.Atoi(startTimePointStr)
	fmt.Println(startTimePointInt)
}


// 根据字段提取，[]map[string]interface{} 中的string信息
func TestArrayValue2Array(t *testing.T) {
	maps := make([]map[string]interface{},4)
	maps[0] = make(map[string]interface{})
	maps[0]["name"] = "综合0"
	maps[0]["card_id"] = 11
	maps[0]["price"] = 100.0
	maps[0]["salesnum"] = 1003

	maps[1] = make(map[string]interface{})
	maps[1]["name"] = "综合1"
	maps[1]["card_id"] = 12
	maps[1]["price"] = 100.1
	maps[1]["salesnum"] = 1004

	maps[2] = make(map[string]interface{})
	maps[2]["name"] = "综合2"
	maps[2]["card_id"] = 13
	maps[2]["price"] = 200.2
	maps[2]["salesnum"] = 1005

	maps[3] = make(map[string]interface{})
	maps[3]["name"] = "综合3"
	maps[3]["card_id"] = 14
	maps[3]["price"] = 300.5
	maps[3]["salesnum"] = 1006

	cardIds := ArrayValue2Array("card_id",maps)
	fmt.Println(cardIds)
}