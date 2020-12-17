/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/15 10:04
@Description:数据库框架封装

*********************************************/
package mysqldb

import (
	"GoInduction/mysqldb/base"
	"GoInduction/tool"
	"github.com/astaxie/beego/orm"
)

// 表结构体
type NeedToModel struct {
	Model *base.Model      // 数据赠送改查的基本结构体
	Field NeedToModelField // 表字段
}

// 表字段
type NeedToModelField struct {
	T_table  string `default:"table"`  // 表名待处理
	F_bus_id string `default:"bus_id"`
}

// 初始化
func (m *NeedToModel) Init(ormer ...orm.Ormer) *NeedToModel {
	tool.ReflectModel(&m.Field)
	m.Model = base.NewMode(m.Field.T_table, ormer...)
	return m
}

// 新增数据
func (m *NeedToModel) Insert(data map[string]interface{}) int {
	result, _ := m.Model.Data(data).Insert()
	return result
}

//批量添加
func (m *NeedToModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := m.Model.InsertAll(data)
	return  result
}

// 更新数据
func (m *NeedToModel) Update(where, data map[string]interface{}) bool {
	if len(where) == 0 {
		return false
	}
	_, err := m.Model.Where(where).Data(data).Update()
	if err != nil {
		return false
	}
	return true
}

// 单条数据查询
func (m *NeedToModel) Find(where map[string]interface{}) map[string]interface{} {
	if len(where) == 0 {
		return make(map[string]interface{})
	}
	return m.Model.Where(where).Find()
}

// 基础查询（多条）
func (m *NeedToModel) Select(where map[string]interface{}) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}
	return m.Model.Where(where).Select()
}

// 带分页查询 （多条）
func (m *NeedToModel) SelectByPage(where map[string]interface{}, start, limit int) []map[string]interface{} {
	if len(where) == 0 {
		return make([]map[string]interface{}, 0)
	}

	return m.Model.Where(where).Limit(start, limit).OrderBy(m.Field.F_bus_id + " DESC ").Select()
}
