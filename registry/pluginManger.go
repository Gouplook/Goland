/**
 * @Author: Yinjinlin
 * @Description: 管理注册类插件
 * @File:  pluginManger
 * @Version: 1.0.0
 * @Date: 2020/12/5 20:50
 */
package registry

import (
	"context"
	"fmt"
	"sync"
)
var (
	// 初始化一个全局对象
	pluginMgr = &PluginManger{
		plugins: make(map[string]Registry),
	}
)
type PluginManger struct {
	// Registr --- interface
	plugins map[string]Registry
	lock sync.Mutex  //exclusion 互斥锁

}
// 注册插件
func RegisterPlugin(registry Registry)(err error){
	return pluginMgr.registerPlugin(registry)
}
// 初始化注册中心
func InitRegistry(ctx context.Context,name string,opts ...Option)(registry Registry, err error){
	return pluginMgr.initPlugin(ctx, name, opts...)
}
func (p *PluginManger)registerPlugin(plugin Registry)(err error){
	p.lock.Lock()
	defer p.lock.Unlock()
	// 调用registry
	_,ok := p.plugins[plugin.Name()]
	if ok{
		err = fmt.Errorf("duplicate registry plugin")
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

func (p *PluginManger)initPlugin(ctx context.Context,name string,opts ...Option)(registry Registry,err error){
	//查找对应的插件是否存在
	p.lock.Lock()
	defer p.lock.Unlock()
	//判断插件是否存在
	plugin, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exists",name)
		return
	}
	registry = plugin
	//调用registry 初始化
	err = plugin.Init(ctx, opts...)
	return
}
