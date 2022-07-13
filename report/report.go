package report

import (
	"fmt"
	"fofa/logger"
	"regexp"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var (
	fields = map[string]string{"A1": "Host", "B1": "IP", "C1": "Port", "D1": "Server", "E1": "Domain",
		"F1": "Title", "G1": "Country", "H1": "Province", "I1": "City", "J1": "ICP"}
	colWidth = map[string]float64{"A": 27, "B": 15, "C": 7, "D": 20,
		"E": 20, "F": 40, "G": 8, "H": 15, "I": 15, "J": 20}
)

func cleanSheetName(sheetName string) (result string) {
	regexp1 := regexp.MustCompile(`[\[\]:\*\?/\\]+`)
	result = regexp1.ReplaceAllString(sheetName, "-")
	regexp2 := regexp.MustCompile(`[']+`)
	result = regexp2.ReplaceAllString(result, "-")

	if len(result) > 30 {
		result = result[0:30]

		pos := len(result) - 1
		lenx := 0
		for result[pos] < 32 || result[pos] > 126 {
			lenx++
			if pos == 0 {
				break
			} else {
				pos--
			}
		}
		remainder := lenx % 3
		result = result[0 : len(result)-remainder]
		return
	}
	return
}

func formatHost(host string, protocol string) (nhost string) {
	nhost = host
	if !strings.Contains(host, "http") {
		if protocol == "http" {
			nhost = "http://" + host
		} else if protocol == "https" {
			nhost = "https://" + host
		}
	}
	return
}

func WriteXlsx(fResult map[string][][]string, output string) {

	logger.Info(fmt.Sprintf("结果输出 %v , 等待...", output))

	f := excelize.NewFile()
	firstSheet := true

	for q, r := range fResult {
		sheetName := cleanSheetName(q)

		if firstSheet == true {
			f.SetSheetName("Sheet1", sheetName)
			f.SetActiveSheet(0)
			firstSheet = false
		} else {
			index := f.NewSheet(sheetName)
			f.SetActiveSheet(index)
		}

		for k, v := range colWidth {
			f.SetColWidth(sheetName, k, k, v)
		}

		for k, v := range fields {
			f.SetCellValue(sheetName, k, v)
		}

		for i, line := range r {
			row := i + 2
			for j, cell := range line {
				if j == 10 { // 协议
					continue
				}
				if j == 0 {
					cell = formatHost(cell, line[10])
				} else if j == 5 {
					cell = strings.TrimSpace(cell)
				}
				col := byte(65) + byte(j)
				axis := string([]byte{col}) + fmt.Sprint(row)
				f.SetCellValue(sheetName, axis, cell)
			}
		}
	}

	f.SetActiveSheet(0)

	if err := f.SaveAs(output); err != nil {
		logger.Warn(err.Error())
	}

	logger.Success(fmt.Sprintf("成功写入到 %v 。", output))
}
