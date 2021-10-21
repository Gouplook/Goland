/**
 * @Author: yinjinlin
 * @File:  gotoexcel.go
 * @Description: 查找到数据，生产excel表，导出
 * @Date: 2021/10/9 上午10:56
 */

package goexcel

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

// 思路：
// 调用接口，传过来数据生产excel表格，根据选择的页面生产excel数据

var (
	OrderTypeRelation    = map[int]string{}
	PayChannelRelation   = map[int]string{}
	PayTypeRelation      = map[int]string{}
	OrderChannelRelation = map[int]string{}
)

// 后台财务订单列表-出参
type GetAdminFinanceOrderListBase struct {
	OrderSn              string // 订单编号
	ShopId               int
	ShopName             string  // 门店名称
	CityId               int     // 城市ID
	CityName             string  // 城市名称
	OrderChannel         int     // 订单来源渠道：1-普通订单 2=店内码牌购卡 3=店内码牌支付
	OrderType            int     // 订单类型 1=单项目订单 2=卡项订单  3=商品订单
	TotalAmount          float64 // 订单总金额
	PayChannel           int     // 支付渠道 1=原生支付 2=杉德支付 3=建行直连 4=平安银行
	PayType              int     // 支付方式 1=支付宝直连 2=微信直连 3=现金 4=渠道支付
	PayTimeStr           string  // 支付成功时间，格式：“2021/09/01 14:05:01”
	PayFee               float64 // 支付手续费
	DepositRatio         string  // 存管比例
	DepositAmount        float64 // 存管金额
	InsureAmount         float64 // 保费
	RenewInsureAmount    float64 // 续保费用
	ComServiceChargeRate string  // 综合服务费费率
	PlatformAmount       float64 // 平台手续费
	BusAmount            float64 // 商家应得金额

}

func CreatePayOrderExcel(data []GetAdminFinanceOrderListBase) (fileName string) {
	file := xlsx.NewFile()
	// 设计表格字体
	xlsx.SetDefaultFont(12, "FangSong")
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	titles := []string{
		"订单编号", "门店名称", "订单类型", "城市", "订单来源", "订单总金额", "支付渠道", "支付方式", "付款时间",
		"代收商户银行手续费", "存管比例", "存管金额", "保费", "续保保费", "综合服务费费率", "综合服务费金额", "商家应得金额",
	}
	// 添加一行表头&赋值
	titleRow := sheet.AddRow()
	for _, v := range titles {
		cell := titleRow.AddCell()
		cell.SetString(v)
	}
	for _, v := range data {
		// 添加每一行的数值
		rowCellValues := sheet.AddRow()

		// 订单编号
		cell := rowCellValues.AddCell()
		cell.SetString(v.OrderSn)

		// 门店名称
		cell = rowCellValues.AddCell()
		if v.ShopName == "" {
			v.ShopName = "--"
		}
		cell.SetString(v.ShopName)

		// 订单类型 1=单项目订单 2=卡项订单  3=商品订单
		cell = rowCellValues.AddCell()
		value, ok := OrderTypeRelation[v.OrderType]
		if !ok {
			value = "--"
		}
		cell.SetString(value)

		// 城市名称
		cell = rowCellValues.AddCell()
		if v.CityName == "" {
			v.CityName = "--"
		}
		cell.SetString(v.CityName)

		// 订单来源 1-普通订单 2=店内码牌购卡 3=店内码牌支付
		cell = rowCellValues.AddCell()
		value, ok = OrderChannelRelation[v.OrderChannel]
		if !ok {
			value = "--"
		}
		cell.SetString(value)

		// 订单总金额
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.TotalAmount)

		// 支付渠道 1=原生支付 2=杉德支付 3=建行直连 4=平安银行 5=工商银行 6=宁波银行 7=pingpp 8=新大陆国通 9=银盛支付
		cell = rowCellValues.AddCell()
		value, ok = PayChannelRelation[v.PayChannel]
		if !ok {
			value = "--"
		}
		cell.SetString(value)

		// 支付方式 1=支付宝直连 2=微信直连 3=现金 4=渠道支付
		cell = rowCellValues.AddCell()
		value, ok = PayTypeRelation[v.PayType]
		if !ok {
			value = "--"
		}
		cell.SetString(value)

		// 支付成功时间，格式：“2021/09/01 14:05:01”
		cell = rowCellValues.AddCell()
		cell.SetString(v.PayTimeStr)

		// 支付手续费
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.PayFee)
		// 存管比例
		cell = rowCellValues.AddCell()
		// cell.SetFloat(v.DepositRatio)
		cell.SetString(v.DepositRatio)
		// 存管金额
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.DepositAmount)
		// 保费
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.InsureAmount)
		// 续保费用
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.RenewInsureAmount)

		cell = rowCellValues.AddCell()
		cell.SetString(v.ComServiceChargeRate)

		// 综合服务费金额
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.PlatformAmount)
		// 商家应得金额
		cell = rowCellValues.AddCell()
		cell.SetFloat(v.BusAmount)
	}
	// page 表示页数
	page := 1
	// 保存文件
	fileName = fmt.Sprintf("订单信息_%d.xls", page)
	err = file.Save(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	return

}
