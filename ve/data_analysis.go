package ve

import (
	"bufio"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"strconv"
	"strings"
)

// AnalyzeDataSrc 解析数据源
func AnalyzeDataSrc(dataSrc string) (vm VirtualMetric, err error) {
	// 打开配置文件
	file, err := os.Open(dataSrc)
	if err != nil {
		return vm, err
	}
	defer file.Close()

	// 初始化map
	vm.ConstLabels = make(map[string]string, 1)

	// 读取文件第三行
	lineCount := 1
	var content string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		lineCount++
		if lineCount == 4 {
			content = fileScanner.Text()
			break
		}
	}
	// 根据第三行生成desc
	analyseData2Desc(content, &vm)

	// 继续读取文件内容
	for fileScanner.Scan() {
		content = fileScanner.Text()
		vm.Data = append(vm.Data, analyseData2Metric(content))
	}
	return vm, nil
}

// analyseData2Desc 分析数据产生desc
func analyseData2Desc(data string, vm *VirtualMetric) {
	pos1 := strings.Index(data, "{")
	pos2 := strings.Index(data, "}")

	// 获取动态标签
	var params []string
	labelContent := data[pos1+1 : pos2]
	labels := strings.Split(labelContent, ",")
	for _, label := range labels {
		params = strings.Split(label, "=")
		vm.VariableLabels = append(vm.VariableLabels, params[0])
	}

	// 生成desc
	vm.Help = "Virtual Metric"
	vm.Type = prometheus.GaugeValue
	vm.FqName = data[:pos1]
	vm.ConstLabels["mode"] = "test"
	vm.Data = append(vm.Data, analyseData2Metric(data))

	return
}

// analyseData2Metric 分析数据产生metric
func analyseData2Metric(data string) (d Data) {
	pos1 := strings.Index(data, "{")
	pos2 := strings.Index(data, "}")

	// 获取动态标签
	var params []string
	labelContent := data[pos1+1 : pos2]
	labels := strings.Split(labelContent, ",")
	for _, label := range labels {
		params = strings.Split(label, "=")
		d.VlParams = append(d.VlParams, params[0])

	}

	// 设置数值
	d.Value, _ = strconv.ParseFloat(data[len(data)-1:], 64)

	return d
}
