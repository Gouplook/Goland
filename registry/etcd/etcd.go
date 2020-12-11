/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/11 10:21
@Description:

*********************************************/
package etcd

import (
	"GoInduction/registry"
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"path"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaixServiceNum         = 8 //
	MaxSyncServiceInterval = time.Second * 10
)

//etcd 注册插件
type EtcdRegistry struct {
	options            *registry.Options
	client             *clientv3.Client
	serviceCh          chan *registry.Service
	value              atomic.Value // 原子操作
	lock               sync.Mutex
	registryServiceMap map[string]*RegisterService
}

type AllServiceInfo struct {
	serviceMap map[string]*registry.Service
}
type RegisterService struct {
	id          clientv3.LeaseID
	service     *registry.Service
	registered  bool // 判断是否注册
	keepAliceCh <-chan *clientv3.LeaseKeepAliveResponse
}

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		serviceCh:          make(chan *registry.Service, MaixServiceNum),
		registryServiceMap: make(map[string]*RegisterService, MaixServiceNum),
	}
)

func init() {
	allserviceInfo := &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaixServiceNum),
	}
	// -----#####
	etcdRegistry.value.Store(allserviceInfo)
	registry.RegisterPlugin(etcdRegistry)

	go etcdRegistry.run()

}

// 插件的名字
func (e *EtcdRegistry) Name() string {
	return "etcd"
}

// 初始化
func (e *EtcdRegistry) Init(ctx context.Context, opts ...registry.Option) (err error) {
	e.options = &registry.Options{}
	for _, opt := range opts {
		opt(e.options)
	}
	e.client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.options.Addrs,
		DialTimeout: e.options.Timeout,
	})
	if err != nil {
		err = fmt.Errorf("init etcd failed, err:%v", err)
		return
	}
	return
}

// 服务注册
func (e *EtcdRegistry) Register(ctx context.Context, service *registry.Service) (err error) {
	select {
	case e.serviceCh <- service:
	default:
		err = fmt.Errorf("regiser chan is full")
		return
	}
	return
}

// 服务反注册
func (e *EtcdRegistry) Unregister(ctx context.Context, service *registry.Service) (err error) {
	return
}

// 服务发现
func (e *EtcdRegistry) GetService(ctx context.Context, name string)(service *registry.Service, err error) {
	// 一般情况下，都会从缓存中读取
	service,ok := e.getServiceFormCache(ctx, name)
	if ok {
		return
	}
	//如果缓存中没有这个service，则从etcd中读取

	e.lock.Lock()
	defer e.lock.Unlock()
	// 先检测是否已经从etcd中加载成功
	service,ok = e.getServiceFormCache(ctx, name)
	if ok {
		return
	}
	//从etcd中读取指定服务名字的服务信息
	key := e.servicePath(name)
	resp, err := e.client.Get(ctx,key,clientv3.WithPrefix())
	if err != nil {
		return
	}

	service = &registry.Service{
		Name: name,
	}
	for _,kv := range resp.Kvs{
		value := kv.Value
		var tmpService registry.Service
		err = json.Unmarshal(value,&tmpService)
		if err != nil {
			return
		}
		for _, node := range tmpService.Nodes{
			service.Nodes = append(service.Nodes,node)
		}
	}
	allServiceInfoOld := e.value.Load().(*AllServiceInfo)
	var allServiceInfoNew = &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaixServiceNum),
	}

	for key, val := range allServiceInfoOld.serviceMap {
		allServiceInfoNew.serviceMap[key] = val
	}

	allServiceInfoNew.serviceMap[name] = service
	e.value.Store(allServiceInfoNew)
	return

}
// 启动一个定时器
func (e *EtcdRegistry) run() {
	ticker := time.NewTicker(MaxSyncServiceInterval)
	for {
		select {
		case service := <-e.serviceCh:
			registryService, ok := e.registryServiceMap[service.Name]
			if ok {
				for _, node := range service.Nodes {
					registryService.service.Nodes = append(registryService.service.Nodes, node)
				}
				registryService.registered = false
				break
			}
			registryService = &RegisterService{
				service: service,
			}
			e.registryServiceMap[service.Name] = registryService

		case <-ticker.C:
			// ---#####
			e.syncServiceFromEtcd()
		default:
			e.registerOrKeepAlice()
			time.Sleep(time.Millisecond *5)
		}
	}
}

func (e *EtcdRegistry) serviceNodePath(service *registry.Service) string{
	nodeIP := fmt.Sprintf("%s:%d",service.Nodes[0].Ip,service.Nodes[0].Port)
	return path.Join(e.options.RegistryPath, service.Name, nodeIP)
}

func (e *EtcdRegistry) registerService(registryService *RegisterService) (err error) {
	// 租约
	resp,err := e.client.Grant(context.TODO(),e.options.HeartBeat)
	if err != nil {
		return
	}
	registryService.id = resp.ID
	for _,node := range registryService.service.Nodes{
		tmp := &registry.Service{
			Name: registryService.service.Name,
			Nodes: []*registry.Node{
				node,
			},
		}

		data,err := json.Marshal(tmp)
		if err != nil {
			continue
		}
		//
		key := e.serviceNodePath(tmp)
		fmt.Printf("register key:%s\n", key)
		_,err = e.client.Put(context.TODO(), key, string(data),clientv3.WithLease(resp.ID))
		if err != nil {
			continue
		}
		ch, err := e.client.KeepAlive(context.TODO(), resp.ID)
		if err != nil {
			continue
		}

		registryService.keepAliceCh = ch
		registryService.registered = true
	}
	return
}

func (e *EtcdRegistry) registerOrKeepAlice() {
	for _, registryService := range e.registryServiceMap {
		if registryService.registered{
			e.keepAlive(registryService)
			continue
		}
		e.registerService(registryService)
	}
}

func (e *EtcdRegistry) keepAlive(registryService *RegisterService) {
	select {
	case resp := <-registryService.keepAliceCh:
		if resp == nil {
			registryService.registered = false
			return
		}
	}
	return
}

func (e *EtcdRegistry) syncServiceFromEtcd() {
	var allSerceInfoNew = &AllServiceInfo{
		serviceMap: make(map[string]*registry.Service, MaixServiceNum),
	}
	ctx := context.TODO()
	allServiceInfo := e.value.Load().(*AllServiceInfo)
	// 对于缓存每一服务，都需要从etcd中进行更新
	for _, service := range allServiceInfo.serviceMap {
		key := e.servicePath(service.Name)
		resp, err := e.client.Get(ctx, key, clientv3.WithPrefix())
		if err != nil {
			allSerceInfoNew.serviceMap[service.Name] = service
			continue
		}
		seriveNew := &registry.Service{
			Name: service.Name,
		}
		for _, kv := range resp.Kvs {
			value := kv.Value
			var tmpService registry.Service
			err = json.Unmarshal(value, &tmpService)
			if err != nil {
				fmt.Printf("unmarshal failed, err:%v value:%s", err, string(value))
				return
			}

			for _, node := range tmpService.Nodes {
				seriveNew.Nodes = append(seriveNew.Nodes, node)
			}
		}

		allSerceInfoNew.serviceMap[seriveNew.Name] = seriveNew

	}
	e.value.Store(allSerceInfoNew)

}

func (e *EtcdRegistry) servicePath(name string) string {
	return path.Join(e.options.RegistryPath, name)
}


// 服务发现-------
func (e *EtcdRegistry)getServiceFormCache(ctx context.Context,name string)(service *registry.Service, ok bool) {
	allServiceInfo := e.value.Load().(*AllServiceInfo)

	service,ok = allServiceInfo.serviceMap[name]
	return
}

