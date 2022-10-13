package ve

import (
	"fmt"
	"github.com/leonscriptcc/ve/gconfig"
	"github.com/leonscriptcc/ve/gtool"
	"github.com/spf13/viper"
	"log"
)

type PrometheusConfig struct {
	ScrapeConfigs []scrapes `yaml:"scrape_configs"`
}

type scrapes struct {
	JobName       string    `yaml:"job_name"`
	StaticConfigs []statics `yaml:"static_configs"`
}

type statics struct {
	Targets []string
}

// CreatePrometheusConfig 工厂
func CreatePrometheusConfig() {
	// 生成地址目标
	ip, _ := gtool.GetLocalIP()
	targets := make([]string, 0, gconfig.Parameters.Exporters.Num)
	for i := 0; i < gconfig.Parameters.Exporters.Num; i++ {
		targets = append(targets, fmt.Sprintf("%s:%d", ip, i+gconfig.Parameters.Exporters.StartPort))
	}

	pfg := PrometheusConfig{ScrapeConfigs: []scrapes{
		{
			JobName:       "virtual job",
			StaticConfigs: []statics{{Targets: targets}},
		},
	}}

	// 从yaml文件获取nacos配置
	vconfig := viper.New()
	// 添加读取的配置文件路径
	vconfig.AddConfigPath("./")
	// 读取文件类型
	vconfig.SetConfigType("yaml")
	// 设置文件内容
	vconfig.Set("scrape_configs", pfg.ScrapeConfigs)
	err := vconfig.WriteConfigAs("prometheus.yaml")
	if err != nil {
		log.Println("write config failed: ", err)
	}
}
