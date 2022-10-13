package main

import (
	"fmt"
	"github.com/leonscriptcc/ve/gconfig"
	"github.com/leonscriptcc/ve/ve"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func init() {
	// 读取配置文件
	err := gconfig.Load()
	if err != nil {
		log.Panic("config.yml is invalid:", err.Error())
	}
}

func main() {
	// 读取文件中的测试数据
	vm, err := ve.AnalyzeDataSrc(gconfig.Parameters.Exporters.DataSrcs)
	if err != nil {
		log.Panic("analyse data source fail:", err.Error())
	}

	// 初始化exporter
	node := ve.NewNode(vm)

	// 创建prometheus和http绑定的方法
	reg := prometheus.NewRegistry()
	reg.MustRegister(node)
	AHandler := promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	)

	// 端口初始化
	port := gconfig.Parameters.Exporters.StartPort

	// 启动http服务
	for i := 0; i < gconfig.Parameters.Exporters.Num; i++ {
		aMux := http.NewServeMux()
		aMux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
			AHandler.ServeHTTP(w, r)
		})
		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", port+i),
			Handler: aMux,
		}
		go func() {
			log.Println(server.ListenAndServe())
		}()
	}

	// 生成prometheus配置文件
	ve.CreatePrometheusConfig()

	select {}

}
