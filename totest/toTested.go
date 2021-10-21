/**
 * @Author: yinjinlin
 * @File:  toTested
 * @Description:
 * @Date: 2021/10/21 上午11:20
 */

package totest

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

// 已经测试
func MapSplitToStruct(){
	// 获取数据
	// cprrlMap  map[string]interface
	// cprrlMap := cprrlM.GetById(rechargeId)

	// 在cpprlMap后面追加数据，在另外一张表中查找数据
	// cprrlMap["CardId"] = cprMap[cprM.Field.F_card_id]

	// 数据组装后，转化为struc   reply 是结构体
	// _ = mapstructure.WeakDecode(cpprlMap, reply)

}

type CardIcad struct {
	CardId int
	CardSn string
	Name string
}
//
func MapSplitToStruct2(reple CardIcad){
	cMak := map[string]interface{}{}
	cMak["CardId"] = 1002
	cMak["CardSn"] = "JS0003"

	fmt.Println("打印前：----",reple)
	_ = mapstructure.WeakDecode(cMak, &reple)
	fmt.Println("打印后：----",reple.CardSn)
	fmt.Println("打印后：----",reple.CardId)
}


type ReplyConsumeDataConf struct {
	SingleId   int
	SingleName string
	SspId      int
	SspName    string
	Type       int
	Num        int
	Price      float64
	StaffId    []int
}

// 数据中的string --> struct
func StringToStruct(reply ReplyConsumeDataConf){
	// 从数据库获取数据
	// singleLogConf := singleLog[logModel.Field.F_conf_data].(string)
	// 序列化成struct
	// _ = json.Unmarshal([]byte(singleLogConf),&reply)
}


type ArgsBatchHandlerBusFundDepos struct {

}
type ArgsBatchSendCommonConsumeMsg struct {

}
//
type ReplyBatchNewConsumeParams struct {
	BusFundDeposSlice []ArgsBatchHandlerBusFundDepos  //资金存管
	ConsumeMsgSlice   []ArgsBatchSendCommonConsumeMsg //消费确认短信
	RelationLogIds    []int                           //卡包关联消费ids
}

// 结构体中含有切片结构体
func StructSilce(reply ReplyBatchNewConsumeParams){
	// 对于切片需要make
	reply = ReplyBatchNewConsumeParams{
		BusFundDeposSlice: make([]ArgsBatchHandlerBusFundDepos,0),
		ConsumeMsgSlice: make([]ArgsBatchSendCommonConsumeMsg,0),
		RelationLogIds: make([]int, 0),
	}

}




