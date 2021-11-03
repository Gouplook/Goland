/**
 * @Author: yinjinlin
 * @File:  goSensitive
 * @Description:
 * @Date: 2021/11/2 上午11:00
 */

package golibrary

import (
	"fmt"
	"github.com/importcjj/sensitive"
)


// 广告过滤敏感词汇

func Filter(){
	filter := sensitive.New()
	err := filter.LoadWordDict("./dict.txt")
	if err != nil {
		fmt.Println("没有获得文件")
	}
	// filter.UpdateNoisePattern(`x`)

	// str := filter.FindAll("最低价")

	// 验证内容是否ok，如果含有敏感词，则返回
	b, vStr:= filter.Validate("史上最低价")
	fmt.Println(b,vStr)

	fmt.Println()
	fmt.Println("")

}






