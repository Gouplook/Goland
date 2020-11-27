package model

// 测试自动生产数据库表
type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`    // 邮箱
	Password  string `json:"password"` // 密码
	CreatedAt string `json:"created_at"`
}

// 建好数据表，连接数据库
type StudentUser struct {
	Id          int
	Email       string
	Password    string
	Createdtime string
}


