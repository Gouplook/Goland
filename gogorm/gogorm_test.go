/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/25 11:31
@Description:

*********************************************/
package gogorm

import "testing"

func TestCreateUser(t *testing.T) {
	initDb()
	// 新增用户
	//CreateUser()

	// 新增卡项
	CreateCards()
}

func TestFindUser(t *testing.T) {
	initDb()
	//FindUser()
	FindAllUser()
}

func TestAotiSql(t *testing.T) {
	initDb()
	AotiSql()
}

