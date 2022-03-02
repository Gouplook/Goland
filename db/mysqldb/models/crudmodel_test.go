/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/18 14:52
@Description:

*********************************************/
package models

import (
	"GoInduction/dbtool"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

// dataSource : root:123456@tcp(127.0.0.1:3306)/macmysql?charset=utf8
// 新增数据测试
func TestNeedToModel_Insert(t *testing.T) {
	//启动
	dbtool.Init()

	orm.RegisterModel(new(NeedToModel))

	// 第一个参数是数据库别名，第二个参数是是否强制更新
	orm.RunSyncdb("default", false, true)
	//初始化
	insetModel := new(NeedToModel).Init()

	// 数据导入
	//  data := map[string]interface{}{
	//	insetModel.Field.F_email:      "wangyi@163.com",
	//	insetModel.Field.F_password:   "qwer123",
	//	insetModel.Field.F_created_id: time.Now().Local().Format("2006-01-02 15:04:05"),
	//}
	data := map[string]interface{}{
		insetModel.Field.Email:      "wangyi@163.com",
		insetModel.Field.Password:   "qwer123",
		insetModel.Field.CreatedAt: time.Now().Local().Format("2006-01-02 15:04:05"),
	}
	id := insetModel.Insert(data)
	if id <= 0 {
		fmt.Println("Insert failed .....")
	}
	fmt.Println("Insert success !")


	// fmt.Println(42635+21500)
}
