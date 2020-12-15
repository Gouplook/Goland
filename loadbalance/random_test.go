/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/15 15:56
@Description:

*********************************************/
package loadbalance

import (
	"GoInduction/registry"
	"testing"
	"fmt"
	"context"
)

func TestSelect(t *testing.T) {
	balance := &RandomBalance{}
	var weights = [3]int{50, 100, 150}
	var nodes []*registry.Node
	for i := 0; i < 4; i++ {
		node := &registry.Node{
			Ip:     fmt.Sprintf("127.0.0.%d", i),
			Port:   8080,
			Weigh: weights[i%3],
		}
		fmt.Printf("node:%#v\n", node)
		nodes = append(nodes, node)
	}
	countStat := make(map[string]int)
	for i := 0; i < 1000; i++ {
		node, err := balance.Select(context.TODO(), nodes)
		if err != nil {
			//t.Fatalf("select failed, err:%v", err)
			continue
		}
		countStat[node.Ip]++
	}

	for key, val := range countStat {
		fmt.Printf("ip:%s count:%v\n", key, val)
	}
	return
}
