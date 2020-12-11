/*************************************
 * @Author: Yinjinlin
 * @Description:
 * @File:  options
 * @Version: 1.0.0
 * @Date: 2020/12/5 19:49
 ***************************************/
package registry

import "time"

type Options struct {
	Addrs        []string
	Timeout      time.Duration // 超时时间
	RegistryPath string
	HeartBeat    int64 // 心跳
}

// 选项模式
type Option func(opts *Options)

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}

}
func WithRegistryPath(registryPath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registryPath
	}
}
func WithHeartbeat(heartBeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartBeat
	}
}
