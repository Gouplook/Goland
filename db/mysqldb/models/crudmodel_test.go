/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/18 14:52
@Description:

*********************************************/
package models

import (
	"GoInduction/dbtool"
	"fmt"
	"testing"
	"time"
)

// dataSource : root:123456@tcp(127.0.0.1:3306)/macmysql?charset=utf8
// 新增数据测试
func TestNeedToModel_Insert(t *testing.T) {
	//启动
	dbtool.Init()

	//初始化
	insetModel := new(NeedToModel).Init()

	// 数据导入
	data := map[string]interface{}{
		insetModel.Field.F_email:      "wangyi@163.com",
		insetModel.Field.F_password:   "qwer123",
		insetModel.Field.F_created_id: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
	id := insetModel.Insert(data)
	if id <= 0 {
		fmt.Println("Insert failed .....")
	}
	fmt.Println("Insert success !")
}
