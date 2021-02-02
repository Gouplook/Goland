/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/15 13:48
@Description: 过滤

*********************************************/
package loadbalance

import (
	"GoInduction/registry"
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
)

type selectedNodes struct {
	selectedNodesMap map[string]bool
}

type loadbalanceFilterNodes struct {

}

// ------####
func WithBalanceContext(ctx context.Context) context.Context{
	selet := &selectedNodes{
		selectedNodesMap: make(map[string]bool),
	}
	return  context.WithValue(ctx,loadbalanceFilterNodes{},selet)
}

// ------####
func GetSelectedNodes(ctx context.Context) *selectedNodes {
	selet, ok := ctx.Value(loadbalanceFilterNodes{}).(*selectedNodes)
	if !ok {
		return nil
	}
	return selet
}

// 设置节点
func setSelected (ctx context.Context,node *registry.Node) {
	selet := GetSelectedNodes(ctx)
	if selet == nil {
		return
	}
	addr := fmt.Sprintf("%s:%d",node.Ip, node.Port)
	logs.Debug(ctx, "filter node %s",addr)
	selet.selectedNodesMap[addr] = true
}

// 过滤节点
func filterNodes(ctx context.Context, nodes []*registry.Node) []*registry.Node{
	var newNodes []*registry.Node
	selet := GetSelectedNodes(ctx)
	if selet == nil {
		return newNodes
	}
	for _,node := range nodes {
		addr := fmt.Sprintf("%s:%d",node.Ip, node.Port)
		_,ok := selet.selectedNodesMap[addr]
		if ok {
			logs.Debug(ctx, "addr: %s ok", addr)
			continue
		}
		newNodes = append(newNodes,node)
	}
	return newNodes
}