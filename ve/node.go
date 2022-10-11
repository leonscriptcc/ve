package ve

import "github.com/prometheus/client_golang/prometheus"

type Node struct {
	descs   []*prometheus.Desc
	metrics map[string][]prometheus.Metric
}

// NewNode 节点工厂
func NewNode(vms ...VirtualMetric) *Node {
	// 参数声明
	descs := make([]*prometheus.Desc, 0, len(vms))
	metrics := make([]prometheus.Metric, len(vms))
	metricsMap := make(map[string][]prometheus.Metric, len(vms))

	// 添加数据描述
	for i, vm := range vms {
		descs = append(descs, prometheus.NewDesc(
			vm.FqName,
			vm.Help,
			vm.VariableLabels,
			vm.ConstLabels,
		))

		// 添加采集数据
		for _, d := range vm.Data {
			metrics = append(metrics, prometheus.MustNewConstMetric(descs[i], prometheus.GaugeValue, d.Value, d.VlParams...))
		}

		// 数据汇总
		metricsMap[descs[i].String()] = metrics
	}

	return &Node{descs: descs, metrics: metricsMap}
}

// Describe 描述
func (n *Node) Describe(ch chan<- *prometheus.Desc) {
	for _, desc := range n.descs {
		ch <- desc
	}
}

// Collect 数据获取
func (n *Node) Collect(ch chan<- prometheus.Metric) {
	for _, desc := range n.descs {
		metric := n.metrics[desc.String()]
		for _, m := range metric {
			ch <- m
		}
	}
}

type VirtualMetric struct {
	FqName         string
	Type           prometheus.ValueType
	Help           string
	VariableLabels []string
	ConstLabels    prometheus.Labels

	Data []Data // 采集需要的数据
}

type Data struct {
	VlParams []string
	Value    float64
}
