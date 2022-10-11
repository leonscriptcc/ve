package gconfig

import (
	"errors"
	"github.com/leonscriptcc/ve/gtool"
	"github.com/spf13/viper"
)

var Parameters configParameters

// Load 获取配置参数
func Load() error {
	//表示 先预加载匹配的环境变量
	viper.AutomaticEnv()
	// 从yaml文件获取nacos配置
	vconfig := viper.New()
	// 添加读取的配置文件路径
	vconfig.AddConfigPath("./")
	// 设置读取的配置文件
	vconfig.SetConfigName("config")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	// 读取yaml
	err := vconfig.ReadInConfig()
	if err != nil {
		return err
	}
	// 转译yaml文件
	if err = vconfig.Unmarshal(&Parameters); err != nil {
		return err
	}

	// 校验文件是否存在
	for _, exporter := range Parameters.Exporters {
		if !gtool.IsExist(exporter.DataSrc) {
			return errors.New("dataSrc not exist:" + exporter.DataSrc)
		}

	}
	return nil
}

// configParameters 项目配置参数
type configParameters struct {
	Exporters []export `mapstructure:"exporters"`
}

type export struct {
	Port    string `mapstructure:"port"`
	DataSrc string `mapstructure:"dataSrc"`
}
