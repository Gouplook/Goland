package goInterface

import "fmt"

// 继承和接口区别
// 1： 继承主要解决代码的复用性
// 2： 接口价值设计好各种规范 让其他自定义类型去实现方法。
// 3： 继承是is-a 关系 接口是like - a 关系

type Ainterface interface {
	Flying()
}
type Binterface interface {
	Swimming()
}

type Mokeny struct {
	Name string
}

type LitterMokeny struct {
	Mokeny
}

func (m *Mokeny) climbing() {
	fmt.Println("climbing.....")
}

func (l *LitterMokeny) Swimming() {
	fmt.Println("swimming...")
}
func (l *LitterMokeny) Flying() {
	fmt.Println("flying...")
	fmt.Println(l.Name)
}

func AAinterface() {
	mokeny := LitterMokeny{
		Mokeny{
			Name: "swk",
		},
	}

	fmt.Println(mokeny.Name)
	mokeny.climbing()
	mokeny.Flying()
	mokeny.Swimming()
}
