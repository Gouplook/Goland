/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/17 17:01
@Description: 配置基础数据库引擎/配置基础数据库参数

*********************************************/
package driver

import (
	"GoInduction/mysqldb/model"
	"GoInduction/utils"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"fmt"
)

// 如何将这个文件设置为最先启动
func Init() {
//func init() {
	logs.Info("Init driver.go mysql start")

	//"root:123456@tcp(127.0.0.1:3306)/macmysql?charset=utf8"

	// 读取配置文件
	// 读配置文件还有另外一种方法，利用key/value
	confModel := new(utils.YamlConfig)
	cfig := confModel.GetConfig()
	// 设置驱动数据库连接参数
	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s", cfig.Name, cfig.Pwrd, cfig.Host, cfig.Port, cfig.Dbname, cfig.Charset)
	logs.Info("DatabaseDriverConnect String:", dataSource)
	//maxIdle, _ := strconv.Atoi(cfig.Maxidle)
	//maxConn, _ := strconv.Atoi(cfig.Maxconn)
	// 设置注册数据库
	orm.RegisterDataBase("default", cfig.Type, dataSource)
	// 注册mode
	//orm.RegisterModel(new(model.User))
	orm.RegisterModel(new(model.NeedToModel))  // 存在包重复问题
	// 生成表
	// 第一个参数是数据库别名，第二个参数是是否强制更新
	orm.RunSyncdb("default", false, true)

}
