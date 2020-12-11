/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/11 17:34
@Description: 负载均衡--随机算法

*********************************************/
package loadbalance

import (
	"GoInduction/errno"
	"GoInduction/registry"
	"context"
)

// 随机算法
type RandomBalance struct {
	name string
}


//func NewRandomBalance() LoadBalance {
//	return  &RandomBalance{
//		name: "random",
//	}
//}

func NewRandomBalance() LoadBalance{
	return  &RandomBalance{
		name: "random",
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
			//setSelected(ctx, node)
		}
	}()
	return
}