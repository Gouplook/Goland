/**
 * @Author: yinjinlin
 * @File:  gotoexcel_test
 * @Description:
 * @Date: 2021/10/9 上午11:16
 */

package goexcel

import "testing"

func TestCreatExcel(t *testing.T) {
	CreatExcel()
}


func TestCreteExcel2(t *testing.T) {
	CreteExcel2()
}

func TestCreateexcel3(t *testing.T) {
	Createexcel3()
}

func TestCreateExecel4(t *testing.T) {
	CreateExecel4()
}

func TestCreateexcel5(t *testing.T) {
	Createexcel5()
}


func TestCreateExcel(t *testing.T) {

	page := 5
	lists := []ExecelData{
		{
			OrderSn:   "S1245",
			ShopId:    1,
			ShopName:  "店面1",
			CityId:    1,
			CityName:  "北京",
			PayChance: 1,
		},
		{
			OrderSn:   "S1243",
			ShopId:    1,
			ShopName:  "店面2",
			CityId:    1,
			CityName:  "上海",
			PayChance: 2,
		},
		{
			OrderSn:   "S1244",
			ShopId:    1,
			ShopName:  "店面3",
			CityId:    1,
			CityName:  "广州",
			PayChance: 1,
		},
	}

	CreateExcel(page,lists)
}
