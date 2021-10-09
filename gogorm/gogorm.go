/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/25 11:11
@Description: gorm 数据库操作框架

*********************************************/
package gogorm

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

var db *gorm.DB
var err error

// 用户信息
type User struct {
	gorm.Model
	Name string
	Age  int
}

// 继承gorm.Model 不能有两个主键
type Card struct {
	ID        int
	Uid       int `gorm:"primaryKey"` // 自定义主键
	CardName  string
	CardPrcie float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

//Db = mysql
//DbUser = root
//DbPassword = 123456
//DbName = ginblog
//DbHost = 0.0.0.0
//DbPort = 3306

// 初始化数据和参数配置
func initDb() {
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	sqlDns := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True&loc=Local", "root", "123456", "ginblog")

	db, err = gorm.Open(mysql.Open(sqlDns),&gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	_ = db.AutoMigrate(&Card{},&User{})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)


}

// 新增用户
func CreateUser() {
	// 更具指定字段插入数据
	//user := User{Name: "goland22",Age: 32}
	//rs := db.Select("Name").Create(&user)
	//fmt.Println(rs.Error)
	//// 返回记录
	//fmt.Println(rs.RowsAffected)

	// 根据Map插入数据
	//rs := db.Model(&User{}).Create(map[string]interface{}{
	//	"Name":"admin1",
	//	"Age":19,
	//})
	//fmt.Println(rs.Error)
	//// 返回记录
	//fmt.Println(rs.RowsAffected)

	// 根据[]Map插入数据
	rs := db.Model(&User{}).Create([]map[string]interface{}{
		{"Name": "admin2", "Age": 29},
		{"Name": "admin3", "Age": 39},
	})
	fmt.Println(rs.Error)
	// 返回记录
	fmt.Println(rs.RowsAffected)

}

//新增card
func CreateCards() {
	rs := db.Model(&Card{}).Create([]map[string]interface{}{
		{"CardName": "single", "CardPrcie": 23.90},
		{"CardName": "rcrad", "CardPrcie": 123.70},
		{"CardName": "hncard", "CardPrcie": 203.50},
	})
	fmt.Println(rs.RowsAffected)
}



// 查找数据
func FindUser(){
	var user []User
	db.First(&user, 5)
}
// 查询多条
func FindUserList(){
	var user []User
	db.Find(&user,[]int{5,6})
	fmt.Println(user)
}

// 获取全部记录
func FindAllUser(){
	var users []User
	//result := db.Find(&users)
	result := db.Where("name <> ?", "jinzhu").Find(&users)

	fmt.Println(result.RowsAffected)
}


type Result struct {
	ID   int
	Name string
	Age  int
}
// 原生sql
func AotiSql() {
	var result Result
	db.Raw("select * from users where id = ?",3).Scan(&result)
	fmt.Println(result)
}

// 上海敦和汽车装备有限公司
