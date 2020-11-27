/******************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/11/24 下午1:19

*******************************************/
package toml

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type BaseData struct {
	Db DataBase `toml:"dataBase"`
	Se Servers  `toml:"servers"`
	Us User     `toml:"user"`
}

type User struct {
	Name string `toml:"name"`
	Age  int    `toml:"age""`
	Sex  int    `toml:"sex"`
}

type DataBase struct {
	Servers       []string `toml:"servers"`
	ConnectionMax int      `toml:"connection_max"`
	Enabled       bool     `toml:"enabled"`
}

type Servers struct {
	A ServerEn `toml:"a"`
	B ServerEn `toml:"b"`
}

type ServerEn struct {
	IP   string `toml:"ip"`
	Port int    `toml:"port"`
}
type ServerER struct {
}

func ReadConf(fname string) (p *BaseData, err error) {
	var (
		fp       *os.File
		fcontent []byte
	)
	p = new(BaseData)
	if fp, err = os.Open(fname); err != nil {
		fmt.Println("open error ", err)
		return
	}

	if fcontent, err = ioutil.ReadAll(fp); err != nil {
		fmt.Println("ReadAll error ", err)
		return
	}

	if err = toml.Unmarshal(fcontent, p); err != nil {
		fmt.Println("toml.Unmarshal error ", err)
		return
	}


	return
}
