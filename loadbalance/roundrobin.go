/********************************************
@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/11 17:56
@Description: 负载均衡--轮询算法、加权轮询
*********************************************/
package loadbalance

import (
	"GoInduction/errno"
	"GoInduction/registry"
	"context"
)

type RoundRobinBalance struct {
	name  string
	index int
}

func NewRoundRobinBalance() LoadBalance {
	return &RoundRobinBalance{
		name: "roundrobin", // 轮询算法
	}
}

func (r *RoundRobinBalance) Name() string {
	return r.name
}

func (r *RoundRobinBalance) Select(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error) {
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

	r.index = (r.index + 1) % len(nodes)
	node = nodes[r.index]
	return
}
