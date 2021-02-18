/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/18 14:35
@Description: mysql数据的增删改查，

*********************************************/
package base

import (
	"errors"
	"fmt"
	"GoInduction/tool"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strconv"
	"strings"
)

//定义Model结构体
type Model struct {
	table   string                 // 表名
	o       orm.Ormer              // Ormer define the orm interface
	data    map[string]interface{} // 存储新增、更新数据 入参或出参数据
	where   []WhereItem            // 条件查询入参或出参数据
	field   string
	orderBy []string
	limit   []interface{}
	sql     string //
}

// 条件查询结构体
type WhereItem struct {
	Field string
	Value interface{}
}

//实例化Model引用
//@param string table 表名称
func NewMode(table string, ormer ...orm.Ormer) *Model {
	var ormers orm.Ormer
	if len(ormer) > 0 {
		ormers = ormer[0]
	} else {
		ormers = orm.NewOrm()
	}
	return &Model{
		o: ormers,
		//table: T_PREFIX + table, // 表名拼接
		table: table, // 需要conf定义，待conf处理
	}
}

func (m *Model) GetOrmer() orm.Ormer {
	return m.o
}

// ----------------------------------------基本字段组装 -----------------------------------------

//存储新增、更新数据
//使用示例
//maps := make(map[string]interface{})
//maps["name"] = "lidazhao"
//maps["age"]  = 21
//map["level"] = []interface{}{"inc", 1} //自增1
//map["level"] = []interface{}{"dec", 1} //自减1
//map["level"] = []interface{}{"concat", "asdf"} //字符串连接
//Data(maps)
func (m *Model) Data(param map[string]interface{}) *Model {
	m.data = make(map[string]interface{})
	m.data = param
	return m
}

//条件查询  刷选出一条或多条
// map[string]interface{"id":1,"name":[]interface{}{"in", []int{1,2}}} )
// []base.WhereItem{{userinfo.Field.F_sex, 1}, {userinfo.Field.F_reg_channel, []interface{}{"in", []int{1,2}}} }
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
		m.limit = make([]interface{}, 0)
	}
	if len(limit) == 0 {
		m.limit = append(m.limit, start)
	} else {
		m.limit = append(m.limit, start)
		m.limit = append(m.limit, limit[0])
	}
	return m
}

// ----------------------------------------增删改查封装 -----------------------------------------

// 插入数据
// SQL语法格式：INSERT INTO table_name (列1, 列2,...) VALUES (值1, 值2,....);
// OR：INSERT INTO 表名称 VALUES (值1, 值2,....);
func (m *Model) Insert() (int, error) {
	if len(m.data) == 0 {
		return 0, nil
	}
	//param := []interface{}{}
	param := make([]interface{}, 0)
	var colsName, colsValue = "", ""
	for i, v := range m.data {
		colsName += "`" + i + "`" + ","
		//如果为整型则转字符串类型
		colsValue += "?,"
		param = append(param, v)
	}
	colsName = strings.TrimRight(colsName, ",")
	colsValue = strings.TrimRight(colsValue, ",")
	// 组合数据写入SQL
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s);", m.table, colsName, colsValue)
	resData, err := m.o.Raw(sql, param...).Exec()
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", param))
	if err != nil {
		fmt.Println("Inser 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
		return 0, nil
	}
	lastId, err := resData.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastId), err
}

//批量添加
func (m *Model) InsertAll(data []map[string]interface{})(int ,error) {
	if len(data) == 0 {
		return 0, nil
	}
	var keys []string
	var colsName, colsValue = "", ""
	for i, _ := range data[0] {
		colsName += "`" + i + "`" + ","
		keys = append(keys, i)
	}
	colsName = strings.TrimRight(colsName, ",")
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES ", m.table, colsName)
	//values := []interface{}{}
	values := make([]interface{},0)
	for _, v := range data {
		colsValue += "("
		for _, k := range keys {
			colsValue += "?,"
			values = append(values, v[k])
		}
		colsValue = strings.TrimRight(colsValue, ",")
		colsValue += "),"
	}
	colsValue = strings.TrimRight(colsValue, ",")
	sql = fmt.Sprintf("%s %s;", sql, colsValue)
	retData, err := m.o.Raw(sql, values...).Exec()
	m.sql = fmt.Sprintf("%s - `%s`", sql, tool.ArrayString("`, `", values))
	if err != nil {
		fmt.Println("Inser 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
		return 0, nil
	}
	LastId, err := retData.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(LastId), err
}

// 更新数据
// SQL语法格式：UPDATE 表名称 SET 列名称 = 新值 WHERE 列名称 = 某值
func (m *Model) Update() (int, error) {
	//分析参数
	if len(m.data) == 0 {
		return 0, nil
	}
	var updateStr string
	field := []interface{}{}
	for i, v := range m.data {
		if val, ok := v.([]interface{}); ok {
			if len(val) == 2 {
				m.arrayData(i, val[0].(string), val[1], &updateStr, &field)
			}
		} else {
			updateStr += i + "=?,"
			field = append(field, v)
		}
	}
	updateStr = strings.TrimRight(updateStr, ",")
	where, param := m.whereString()

	if where == "" {
		return 0, errors.New("must have where")
	}

	sql := fmt.Sprintf("UPDATE %s SET %s%s", m.table, updateStr, where)
	param = append(field, param...)
	sqlSource, err := m.o.Raw(sql, param...).Exec()
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", param))
	if err != nil {
		fmt.Println("Update 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
		return 0, nil
	}
	rowAffectedId, _ := sqlSource.RowsAffected()
	return int(rowAffectedId), err
}

// 物理删除
// SQL语法格式：DELETE FROM 表名称 WHERE 列名称 = 值
func (m *Model) Delete() (int, error) {
	where, param := m.whereString()
	if where == "" {
		return 0, nil
	}
	sql := fmt.Sprintf("DELETE FROM %s%s", m.table, where)
	sqlSource, err := m.o.Raw(sql, param...).Exec()
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", param))
	if err != nil {
		fmt.Println("Delete 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
		return 0, nil
	}
	rowAffectId, _ := sqlSource.RowsAffected()
	return int(rowAffectId), err
}

//单条查询
func (m *Model) Find() map[string]interface{} {
	var field string
	if m.field == "" {
		m.field = "*"
	} else {
		field = m.field
	}

	where, param := m.whereString()
	// sql 语句
	sql := fmt.Sprintf("SELECT %s FROM %s%s LIMIT 1", field, m.table, where)
	// 执行查询
	var res []orm.Params
	_, err := m.o.Raw(sql, param...).Values(&res)
	// 将param（数组）转换为字符串
	str := tool.ArrayString("`, `", param)
	m.sql = fmt.Sprintf("%s-`%s`", sql, str)
	if err != nil {
		// 打印日志..... 后期统一处理
		fmt.Println("Find 打印日志..... 后期统一处理")
		//	if kcgin.KcConfig.RunMode != kcgin.PROD {
		//		logs.Error("Sql:", sql, " Error,", err.Error())
		//	}
	}
	if len(res) == 0 {
		return make(map[string]interface{})
	}
	return res[0]
}

// 查询多条数据
// SQL语法格式：SELECT column_name,column_name FROM table_name [WHERE Clause] [LIMIT N][ OFFSET M]
func (m *Model) Select() []map[string]interface{} {
	// 字段
	var field string
	if m.field == "" {
		field = "*"
	} else {
		field = m.field
	}
	// 排序
	where, param := m.whereString()
	var orderBy string
	if len(m.orderBy) > 0 {
		for _, v := range m.orderBy {
			orderBy += v + ","
		}
		orderBy = " ORDER BY " + strings.TrimRight(orderBy, ",")
	}
	// 分页
	var limit string
	if len(m.limit) > 0 {
		if len(m.limit) == 1 {
			limit = strconv.Itoa(m.limit[0].(int))
		} else {
			limit = strconv.Itoa(m.limit[0].(int)) + "," + strconv.Itoa(m.limit[1].(int))
		}
		limit = " LIMIT " + limit
	}

	sql := fmt.Sprintf("SELECT %s FROM %s%s%s%s", field, m.table, where, orderBy, limit)
	var res []orm.Params
	// 执行查询
	_, err := m.o.Raw(sql, param...).Values(&res)
	// 将param（数组）转换为字符串
	str := tool.ArrayString("`, `", param)

	m.sql = fmt.Sprintf("%s-`%s`", sql, str)
	if err != nil {
		// 打印日志..... 后期统一处理
		fmt.Println("Select 打印日志..... 后期统一处理")
		//	if kcgin.KcConfig.RunMode != kcgin.PROD {
		//		logs.Error("Sql:", sql, " Error,", err.Error())
		//	}
	}
	maps := make([]map[string]interface{}, len(res))
	if len(res) > 0 {
		for i, v := range res {
			maps[i] = v
		}
	}
	return maps
}

// ----------------------------------------聚合函数封装 ------------------------------------------

//统计数据
// SQL语法格式：SELECT COUNT(column_name) FROM table_name [WHERE Clause]
//Count() 或 Count("id") //id为字段名
func (m *Model) Count(param ...string) int {
	columnName := "*"
	if len(param) != 0 {
		columnName = param[0]
	}
	where, pargm := m.whereString()
	sql := fmt.Sprintf("SELECT COUNT(%s) FROM %s%s", columnName, m.table, where)
	var maps []orm.Params
	_, err := m.o.Raw(sql, pargm...).Values(&maps)
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", pargm))
	if err != nil {
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
		fmt.Println("Count 打印日志..... 后期统一处理")

	}
	num, _ := strconv.Atoi(maps[0]["COUNT("+columnName+")"].(string))
	return num
}

//聚合函数-sum
// SQL语法格式： SELECT SUM(column_name) FROM table_name [WHERE Clause]
func (m *Model) Sum(field string) float64 {
	where, pargm := m.whereString()
	sql := fmt.Sprintf("SELECT SUM(%s) FROM %s%s", field, m.table, where)
	var maps []orm.Params
	_, err := m.o.Raw(sql, pargm...).Values(&maps)
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", pargm))
	if err != nil {
		fmt.Println("Sum 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
	}
	var num float64
	if maps[0]["SUM("+field+")"] != nil {
		num, _ = strconv.ParseFloat(maps[0]["SUM("+field+")"].(string), 64)
	} else {
		num = 0
	}
	return num
}

//聚合函数-Avg
// SQL语法格式：SELECT AVG(column_name) FROM table_name [WHERE Clause]
func (m *Model) Avg(field string) float64 {
	where, pargm := m.whereString()
	sql := fmt.Sprintf("SELECT AVG(%s) FROM %s%s", field, m.table, where)
	var maps []orm.Params
	_, err := m.o.Raw(sql, pargm...).Values(&maps)
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", pargm))
	if err != nil {
		fmt.Println("Avg 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
	}
	var num float64
	if maps[0]["AVG("+field+")"] != nil {
		num, _ = strconv.ParseFloat(maps[0]["AVG("+field+")"].(string), 64)
	} else {
		num = 0
	}
	return num
}

//聚合函数-Min
// SELECT MIN(column_name) FROM table_name [WHERE Clause]
func (m *Model) Min(field string) float64 {
	where, pargm := m.whereString()
	sql := fmt.Sprintf("SELECT MIN(%s) FROM %s%s", field, m.table, where)
	var maps []orm.Params
	_, err := m.o.Raw(sql, pargm...).Values(&maps)
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", pargm))
	if err != nil {
		fmt.Println("Min 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
	}
	var num float64
	if maps[0]["MIN("+field+")"] != nil {
		num, _ = strconv.ParseFloat(maps[0]["MIN("+field+")"].(string), 64)
	} else {
		num = 0
	}
	return num
}

//聚合函数-max
// SELECT MAX(column_name) FROM table_name [WHERE Clause]
func (m *Model) Max(field string) float64 {
	where, pargm := m.whereString()
	sql := fmt.Sprintf("SELECT MAX(%s) FROM %s%s", field, m.table, where)
	var maps []orm.Params
	_, err := m.o.Raw(sql, pargm...).Values(&maps)
	m.sql = fmt.Sprintf("%s-`%s`", sql, tool.ArrayString("`, `", pargm))
	if err != nil {
		fmt.Println("Max 打印日志..... 后期统一处理")
		//if kcgin.KcConfig.RunMode != kcgin.PROD {
		//	logs.Error("Sql:", sql, " Error,", err.Error())
		//}
	}
	var num float64
	if maps[0]["MAX("+field+")"] != nil {
		num, _ = strconv.ParseFloat(maps[0]["MAX("+field+")"].(string), 64)
	} else {
		num = 0
	}
	return num
}

//--------------------------------------辅助查询组织的语句-----------------------------------------

// 组织where字符串
func (m *Model) whereString() (string, []interface{}) {
	var where string = ""
	param := []interface{}{}
	if len(m.where) != 0 {
		for _, v := range m.where {
			if k, ok := v.Value.([]interface{}); ok {
				if len(k) == 2 {
					m.arrayWhere(v.Field, k[0].(string), k[1], &where, &param)
				}
			} else if k, ok := v.Value.([]string); ok {
				if len(k) == 2 {
					m.arrayWhere(v.Field, k[0], k[1], &where, &param)
				}
			} else {
				where += v.Field + "=? AND "
				param = append(param, v.Value)
			}
		}
		where = " WHERE " + strings.TrimRight(where, " AND")
	}
	m.where = make([]WhereItem, 0)
	return where, param
}

//
func (m *Model) arrayWhere(name string, condition string, v interface{}, where *string, param *[]interface{}) {
	switch strings.ToLower(condition) {
	case "eq":
		condition = "="
		break
	case "neq":
		condition = "<>"
		break
	case "gt":
		condition = ">"
		break
	case "egt":
		condition = ">="
		break
	case "lt":
		condition = "<"
		break
	case "elt":
		condition = "<="
		break
	case "between":
		*where += name + " " + condition + " ? and ? AND "
		*param = append(*param, v)
		return
	}

	ars := "?"
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		for i := 1; i < reflect.ValueOf(v).Len(); i++ {
			ars += ",?"
		}
	}
	*where += name + " " + condition + "(" + ars + ") AND "
	*param = append(*param, v)
}

func (m *Model) arrayData(name string, condition string, v interface{}, field *string, param *[]interface{}) {
	switch strings.ToLower(condition) {
	case "inc":
		*field += name + "=" + name + "+?,"
		break
	case "dec":
		*field += name + "=" + name + "-?,"
		break
	case "concat":
		*field += name + "=concat(" + name + ",?),"
		break
	}
	*param = append(*param, v)
}

//------------------------------------------事务处理-----------------------------------------

//事务开始
func (m *Model) Begin() *Model {
	err := m.o.Begin()
	if err != nil {
		logs.Error("Begin Error", err.Error())
	}
	return m
}

//事务提交
func (m *Model) Commit() *Model {
	err := m.o.Commit()
	if err != nil {
		logs.Error("Commit Error", err.Error())
	}
	return m
}

//事务回滚
func (m *Model) RollBack() *Model {
	err := m.o.Rollback()
	if err != nil {
		logs.Error("RollBack Error", err.Error())
	}
	return m
}

