/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/23 下午5:14
@Des  : mysql CRID实现

*******************************************/
package mysqldb

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//定义变量
var (
	T_PREFIX = beego.AppConfig.String("db.prefix")
)

//定义Model结构体
type Model struct {
	table   string
	o       orm.Ormer
	limit   []interface{}
	orderBy []string
	where   []WhereItem
	data    map[string]interface{}
	field   string
	sql     string
}

// where条件结构体
type WhereItem struct {
	Field string
	Value interface{}
}

//操作另一张表,表名不需要扩展
func (m *Model) Table(table string) *Model {
	m.table = T_PREFIX + table
	return m
}



