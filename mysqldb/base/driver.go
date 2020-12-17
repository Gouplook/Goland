/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/17 17:01
@Description: 配置基础数据库引擎/配置基础数据库参数

*********************************************/
package base

import (
	"GoInduction/utils"
	"fmt"
	"github.com/astaxie/beego/logs"
)

func init(){
	logs.Info("Init driver.go mysql start")

	//"root:123456@tcp(127.0.0.1:3306)/macmysql?charset=utf8"
	// 读取配置文件
	y := new(utils.YamlConfig)
	//设置驱动数据库连接参数
	dataSource := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s",y.Name,y.Pwrd,y.Host,y.Port,y.Dbname,y.Charset)
	logs.Info("DatabaseDriverConnect String:",dataSource)

}
