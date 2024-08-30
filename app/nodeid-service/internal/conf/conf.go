package conf

import (
	configutil "github.com/go-micro-saas/service-kit/config"
)

var (
	serviceConfig = &ServiceConfig{}
)

func GetServiceConfig() *ServiceConfig {
	return serviceConfig
}

// LoadServiceConfig 加载服务配置
// 由 setuputil.NewLauncherManager 进行加载赋值
func LoadServiceConfig() configutil.Option {
	return configutil.WithOtherConfig(serviceConfig)
}
