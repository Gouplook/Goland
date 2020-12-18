/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/26 上午9:04

*******************************************/
package base

import (
	"testing"
)

// 新增数据
func TestOrmInsert(t *testing.T) {
	OrmInit()
	OrmInsert()
}


func TestOrmUpdate(t *testing.T) {
	OrmInit()
	OrmUpdate()
}

func TestOrmUpdate2(t *testing.T) {
	OrmInit()
	OrmUpdate2()
}

func TestOrmRead(t *testing.T) {
	OrmInit()
	OrmRead()
}
func TestOrmDelet(t *testing.T) {
	OrmInit()
	OrmDelet()
}
