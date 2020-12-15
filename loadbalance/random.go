/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/11 17:34
@Description: 负载均衡--随机算法/加权

*********************************************/
package loadbalance

import (
	"GoInduction/errno"
	"GoInduction/registry"
	"context"
	"math/rand"
)

// 随机算法
type RandomBalance struct {
	name string
}



func NewRandomBalance() LoadBalance{
	return  &RandomBalance{
		name: "random", // 随机算法
	}
}

func (r *RandomBalance)Name() string{
	return r.name
}

func (r *RandomBalance)Select(ctx context.Context, nodes[]*registry.Node)(node *registry.Node, err error) {
	if len(nodes) == 0 {
		err = errno.NotHaveInstance
		return
	}

	defer func() {
		if node != nil {
			setSelected(ctx, node)
		}
	}()

	var newNodes = filterNodes(ctx, nodes)
	if len(newNodes) == 0 {
		err = errno.AllNodeFailed
		return
	}
	// 加权
	var totalWeight int
	for _,val := range newNodes {
		if val.Weigh == 0 {
			val.Weigh = DefaultNodeWeight
		}
		totalWeight += val.Weigh
	}
	curWeight := rand.Intn(totalWeight)
	curIndex := -1
	for Index, node := range nodes {
		curWeight -= node.Weigh
		if curWeight < 0 {
			curIndex = Index
			break
		}
	}
	if curIndex == -1 {
		err = errno.AllNodeFailed
		return
	}

	node = nodes[curIndex]

	return

}