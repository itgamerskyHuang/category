package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

// 设置配置中心
func GetConsulConfig(host string, prefix string) (config.Config, error) {
	consulSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(host),
		// 设置前缀， 不设置前缀的话为默认前缀 /micro/config
		consul.WithPrefix(prefix),
		// 是否移除前缀，这里设置为true,表示可以不带前缀，直接获取consul配置
		consul.StripPrefix(true),
	)
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}
	// 加载配置
	err = config.Load(consulSource)
	return config, nil
}
