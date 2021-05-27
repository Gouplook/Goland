/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/8 13:31
@Description:

*********************************************/
package redisoperation

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 通过go 向redis 写入数据和读取数据
func GoRedisReadWrite() {
	// 1: 链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err) //
	}
	defer conn.Close()

	// 2：通过go向redis写入数据，key-value
	conn.Do("set", "name", "kuncun")
	conn.Do("set", "age", "19")
	conn.Do("set", "address", "上海市宝山区")

	conn.Do("HMSet", "user02", "name", "john", "age1", "22")

	// 3： 通过go从redis中读数据
	str, _ := redis.String(conn.Do("get", "address"))
	age, _ := redis.Int(conn.Do("get", "age"))
	age1, _ := redis.Int(conn.Do("get", "age1"))

	r, _ := redis.Strings(conn.Do("HMGet", "user02", "name", "age"))

	fmt.Println(str)
	fmt.Println(age)
	fmt.Println(age1)
	fmt.Println(r)

}

// 链接池
var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲链接数
		MaxActive:   0,   // 表示和数据库的最大链接数， 0 表示没有限制
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","127.0.0.1:6379")
		},
	}
}
func GoRedisPool(){
	// 1: 先从pool中取出一个链接
	conn := pool.Get()
	defer conn.Close()

	// 对数据进行操作
	_,err := conn.Do("set","Name","Mac")
	if err != nil {
		panic(err)
	}
	r, _ := redis.String(conn.Do("Get", "Name"))
	fmt.Println("1:==",r)

	//

	// 获得另一个链接
	conn = pool.Get()
	_,err = conn.Do("set","Name2","Mac2")
	if err != nil {
		panic(err)
	}
	r2, _ := redis.String(conn.Do("Get", "Name2"))
	fmt.Println("2:==",r2)

	// 获得另一个链接
	conn = pool.Get()
	_,err = conn.Do("set","Name3","Mac3")
	if err != nil {
		panic(err)
	}
	r3, _ := redis.String(conn.Do("Get", "Name3"))
	fmt.Println("2:==",r3)

 // go get github.com/gomodule/redigo/redis@v2.0.0+incompatible
}
