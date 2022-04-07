/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午4:58

*******************************************/
package tool

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//map转数组 (提取map中的key值）
func ArrayKeys(maps map[int]interface{}) []int {
	//分析参数
	if len(maps) == 0 {
		return make([]int, 0)
	}
	var arr = make([]int, 0)
	for i, _ := range maps {
		arr = append(arr, i)
	}
	return arr
}

//map数组转数组（根据字段提取信息）
func ArrayValue2Array(field string, maps []map[string]interface{}) []int {
	//分析参数
	if len(maps) == 0 {
		return make([]int, 0)
	}
	var arr = make([]int, 0)
	for _, m := range maps {
		v, ok := m[field]
		if ok {
			if vs, p := v.(string); p {
				n, _ := strconv.Atoi(vs)
				arr = append(arr, n)
			}
			if vs, p := v.(int); p {
				arr = append(arr, vs)
			}
		}
	}
	return arr
}

//map数组转map (根据字段，在map切片中提取map）
func ArrayRebuild(field string, maps []map[string]interface{}) map[string]interface{} {
	//分析参数
	if len(maps) == 0 {
		return make(map[string]interface{}, 0)
	}
	var reMap = make(map[string]interface{})
	for _, m := range maps {
		v, ok := m[field]
		if ok {
			if vs, p := v.(int); p {
				reMap[strconv.Itoa(vs)] = m
			}
			if vs, p := v.(string); p {
				reMap[vs] = m
			}
			if vs, p := v.(float64); p {
				reMap[strconv.FormatFloat(vs, 'f', -1, 64)] = m
			}
			if vs, p := v.(float32); p {
				reMap[strconv.FormatFloat(float64(vs), 'f', -1, 64)] = m
			}
		}
	}
	return reMap
}

// 数组map排序
// 思路：1： 先定义两个容器，mapData/keys 存放数据和存放key，
//      2： 对key进行排序 sort.string
//      3： 遍历key，将key中字段所对应的值，存放起来，return

func SortsMap(field string, maps []map[string]interface{}) []map[string]interface{} {
	var mapData = make(map[string]interface{}) // map make不需要指定大小
	var keys = make([]string, 0)               // 切片make时，需要指定大小
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
		mapData[vs.(string)] = v
		keys = append(keys, vs.(string))
	}
	sort.Strings(keys)
	remapData := make([]map[string]interface{}, 0)
	for _, v := range keys {
		remapData = append(remapData, mapData[v].(map[string]interface{}))
	}
	return remapData
}

//把数组转换为字符串
//@param  string separator       转换分隔符
//@param  interface{}  interface 待转换数据
//@return string
// []int{1,2,3} ---> str : 1,2,3  ,[]string也适用

func ArrayString(separator string, array interface{}) (str string) {
	str = strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", separator, -1)
	return str
}

// 数组去重 int
func ArrayUniqueInt(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	newArr := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if arr[i] == 0 {
			continue
		}
		if repeat == false {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 数组去重 string
func ArrayUniqueString(arr []string) []string {
	if len(arr) == 0 {
		return arr
	}
	newArr := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if arr[i] == "" {
			continue
		}
		if repeat == false {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}
