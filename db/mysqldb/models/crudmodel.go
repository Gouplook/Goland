/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/18 14:32
@Description: 数据库 增删改查的封装

*********************************************/
package models

import (
	"GoInduction/base"
	"github.com/astaxie/beego/orm"
)

// 表结构体
type NeedToModel struct {
	Model *base.Model      // 数据赠删改查的基本结构体
	Field NeedToModelField // 表字段
}

// 表字段
type NeedToModelField struct {
	// T_table      string `default:"table"` // 表名待处理
	// F_id         int    `default:"id"`
	// F_email      string `default:"email"`
	// F_password   string `default:"password"`
	// F_created_id string `default:"created_id"`

	Table     string `default:"table"` // 表名待处理
	Id        string `default:"id"`
	Email     string `default:"email"`
	Password  string `default:"password"`
	CreatedAt string `default:"created_at"`
	Uid       string `default:"uid"`
	Name      string `default:"name"`
}

// 初始化
// func (m *NeedToModel) Init(ormer ...orm.Ormer) *NeedToModel {
//	tool.ReflectModel(&m.Field)
//	m.Model = base.NewMode(m.Field.T_table, ormer...)
//	return m
// }
func (m *NeedToModel) Init(ormer ...orm.Ormer) *NeedToModel {
	// tool.ReflectModel(&m.Field)
	m.Model = base.NewMode(m.Field.Table, ormer...)
	m.Model = base.NewMode(m.Field.Table, ormer...)
	return m
}

// 新增数据
func (m *NeedToModel) Insert(data map[string]interface{}) int {
	result, _ := m.Model.Data(data).Insert()
	return result
}

// 批量添加
func (m *NeedToModel) InsertAll(data []map[string]interface{}) int {
	if len(data) == 0 {
		return 0
	}
	result, _ := m.Model.InsertAll(data)
	return result
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

// 查找(带范围）
func (m *NeedToModel) FindBetween(uid, start, end int) (data map[string]interface{}) {
	wh := []base.WhereItem{
		{m.Field.Uid, uid},
		{m.Field.CreatedAt, []interface{}{"between", []int{start, end}}},
	}

	return m.Model.Where(wh).Find()

}

// 查找 以name开头匹配的 在id中筛选
func (m *NeedToModel) FindLike(name string, ids []int)(data []map[string]interface{}){

	if len(name) == 0 {
		return make([]map[string]interface{},0)
	}

	wh := []base.WhereItem{
		{m.Field.Name,[]interface{}{"LIKE",name+"%"}},
		{m.Field.Id,[]interface{}{"IN",ids}},
	}

	return m.Model.Where(wh).Select()

}

// 根据条件获取单条数据(这种条件查询
func (m *NeedToModel) GetById(id int, fileld ...[]string) map[string]interface{} {
	if id <= 0 {
		return map[string]interface{}{}
	}
	if len(fileld) > 0 {
		m.Model.Field(fileld[0])
	}
	rs := m.Model.Where([]base.WhereItem{
		{m.Field.Id, id},
	}).Find()

	return rs
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

	// 需要修改,CreatedAt是字段
	return m.Model.Where(where).Limit(start, limit).OrderBy(m.Field.CreatedAt + " DESC ").Select()
}

// 根据[]int查数据
func (m *NeedToModel) GetByIds(ids []int) []map[string]interface{} {
	if len(ids) == 0 {
		return []map[string]interface{}{}
	}
	rs := m.Model.Where(map[string]interface{}{
			m.Field.Id: []interface{}{"in",ids},
	}).Select()
	return rs
}
