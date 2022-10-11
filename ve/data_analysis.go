package ve

import (
	"bufio"
	"log"
	"os"
)

// AnalyzeDataSrc 解析数据源
func AnalyzeDataSrc(dataSrc string) {
	file, _ := os.Open(dataSrc)
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		log.Println(fileScanner.Text())
		lineCount++
	}
	defer file.Close()
}
