/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/11 17:35
@Description:负载均衡解决方案

*********************************************/
package loadbalance

import (
	"GoInduction/registry"
	"context"
)

const (
	DefaultNodeWeight = 100
)


type LoadBalanceType int

const (
	LoadBalanceTypeRandom = iota // 随机算法
	LoadBalanceTypeRoundRobin   // 	轮询算法
)

type LoadBalance interface {
	Name() string
	Select(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error)
}


func GetLoadBalance(balanceType LoadBalanceType) (balancer LoadBalance) {

	switch (balanceType) {
	case LoadBalanceTypeRandom:
		balancer = NewRandomBalance()
	case LoadBalanceTypeRoundRobin:
		balancer = NewRoundRobinBalance()
	default:
		balancer = NewRandomBalance()
	}
	return
}
