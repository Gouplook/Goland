/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/4/19 13:05
@Description:

*********************************************/
package tool

import (
	"encoding/json"
	"fmt"
)

// 场景1：
// 传过来的是结构体，需要序列化一下，存到数据库中
// specInfo, _ := json.Marshal(specIds)
// mSS.Field.F_spec_info: string(specInfo),

// 场景2：
// rpc添加描述 struct ---> string 前端传过来的数据 存储到数据库中 marshl
// rpc获取描述 string ---> struct 从数据库中取出数据给前端 Unmarshl（以结构体形式返回）

// Product 商品信息
type Product struct {
	Name      string `json:"name"`
	ProductID int64  `json:"product_id,omitempty"` // omitempty，tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
	Number    int    `json:"-"`                    // 表示不进行序列化
	Price     float64
	IsOnSale  bool
}




// Marshal
func JsonMarshal() {
	p := &Product{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}

// type，有些时候，我们在序列化或者反序列化的时候，可能结构体类型和需要的类型不一致，
// 这个时候可以指定,支持string,number和boolean
type Good struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func JsonUnmarshal() {
	var data = `{"name":"Xiao mi 6","product_id":"10","number":"10000","price":"2499","is_on_sale":"true"}`
	// 一般从数据查到的数据或网落之前传送过来的。
	g := &Good{}
	err := json.Unmarshal([]byte(data), g)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")
	fmt.Println(*g)
	fmt.Println(g.Name)
	fmt.Println(g.ProductID)
	fmt.Println(g.Number)
	fmt.Println(g.Price)
}
