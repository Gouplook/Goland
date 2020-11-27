/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/27 上午10:11

*******************************************/
package base

import (
	"strings"
)

type Model struct {
	table   string                 // 表名
	data    map[string]interface{} // 存储新增、更新数据 入参或出参数据
	where   []WhereItem            // 条件查询入参或出参数据
	field   string
	orderBy []string
	limit   []interface{}
	sql     string
}

// 条件查询结构体
type WhereItem struct {
	Field string
	Value interface{}
}

//存储新增、更新数据
//使用示例
//maps := make(map[string]interface{})
//maps["name"] = "lidazhao"
//maps["age"]  = 21
//map["level"] = []interface{}{"inc", 1} //自增1
//map["level"] = []interface{}{"dec", 1} //自减1
//map["level"] = []interface{}{"concat", "asdf"} //字符串连接
func (m *Model) Data(param map[string]interface{}) *Model {
	m.data = make(map[string]interface{})
	m.data = param
	return m
}

//条件查询  刷选选出一条或多条
// map[string]interface{"id":1,"name":[]interface{}{"in", []int{1,2}}} )
// []base.WhereItem{ {userinfo.Field.F_sex, 1}, {userinfo.Field.F_reg_channel, []interface{}{"in", []int{1,2}}} }
func (m *Model) Where(param interface{}) *Model {
	// 刷选选出一条,先进行类型断言一下，转换成map类型。
	if where, ok := param.(map[string]interface{}); ok {
		if len(where) == 0 {
			return m
		}
		if m.where == nil {
			m.where = make([]WhereItem, 0)
		}
		for k, v := range where {
			m.where = append(m.where, WhereItem{k, v})
		}
	}
	// 刷选选出多条
	if where, ok := param.([]WhereItem); ok {
		if len(where) == 0 {
			return m
		}
		if m.where == nil {
			m.where = make([]WhereItem, 0)
		}
		m.where = append(m.where, where...)
	}

	return m
}

//查询的字段
//使用示例 Field([]string{"name","age"})
func (m *Model) Field(param ...[]string) *Model {
	if len(param) == 0 {
		m.field = "*"
	} else {
		s := strings.Join(param[0], ",")
		//m.field = strings.TrimRight(strings.Join(param[0], ","),",")
		m.field = strings.TrimRight(s, ",")
	}
	return m
}

//设置排序
//使用示例 OrderBy("id asc","age desc")
func (m *Model) OrderBy(params ...string) *Model {
	if len(params) == 0 {
		return m
	}
	if m.orderBy == nil {
		m.orderBy = make([]string, len(params))
	}
	for k, v := range params {
		v = strings.ToLower(v)
		m.orderBy[k] = v
	}

	return m
}

//设置查询范围
//使用示例 Limit(10) limit(0,10)
func (m *Model) Limit(start interface{}, limit ...interface{}) *Model {
	if m.limit == nil {
		m.limit = make([]interface{},0)
	}
	if len(limit) == 0{
		m.limit = append(m.limit, start)
	}else {
		m.limit = append(m.limit,start)
		m.limit = append(m.limit,limit[0])
	}
	return  m
}
