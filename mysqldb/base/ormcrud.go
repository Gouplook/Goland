/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/26 上午8:53
@Des  : mysql 数据库操作
		Insert / Update / Delete /Read

*******************************************/
package base

import (
	"GoInduction/mysqldb/model"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// 先进行初始化
func OrmInit(){
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/macmysql?charset=utf8")
	// 注册mode
	orm.RegisterModel(new(model.User))
	// 生成表
	// 第一个参数是数据库别名，第二个参数是是否强制更新
	orm.RunSyncdb("default", false, true)
}

func OrmInsert() {
	o := orm.NewOrm()
	o.Using("default")
	user := new(model.User)
	user.Password = "qwert123"
	user.Email = "tengxun@163.com"
	user.CreatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	o.Insert(user)
}

// 默认情况下，orm更新所有数据，没有的写的字段，将清空
func OrmUpdate(){
	o := orm.NewOrm()
	user := model.User{Id: 1104}
	if o.Read(&user) == nil {
		user.Password = "xiugiamima123"
	}
	o.Update(&user)
}
//Update 默认更新所有的字段，可以更新指定的字段
func OrmUpdate2(){
	o := orm.NewOrm()
	//user := new(model.User)
	user := model.User{Id: 1106}
	user.CreatedAt = time.Now().Local().Format("2006-01-02 15:04:05")
	o.Update(&user,"CreatedAt")
}
// 查询
// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
func OrmRead(){
	o := orm.NewOrm()
	user := model.User{Id :1107}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Password)
	}
}

// 删除
func OrmDelet(){
	o := orm.NewOrm()
	user := model.User{Id: 1106}
	o.Delete(&user)
}




