/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2020/12/5 19:44
 */
package registry

type Service struct {
	Name  string  `json:"name"`
	Nodes []*Node `json:"nodes"`
}

// 节点
type Node struct {
	Id   string `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Weigh int `json:"weigh"` // ---###
}
