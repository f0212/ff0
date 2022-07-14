package xlsx

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/fatih/color"
	"github.com/liushuochen/gotable"
	"os"
)

const (
	max = 50
)

func Xlsx(Pathxlsx, level string) {
	switch level {
	case "1":
		Level1(Pathxlsx)
	case "2":
		Level2(Pathxlsx)
	case "3":
		Level3(Pathxlsx)
	}
}
func Level1(excelPath string) {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}
	rows, err := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex())) //读取全部数据
	table, err := gotable.Create("Host", "Ip", "Port")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
	}
	i := 0
	for _, row := range rows {
		if i == 0 {
			i = i + 1
			continue
		}
		err := table.AddRow([]string{row[0], row[1], row[2]})
		if err != nil {
			return
		}
	}
	blue := color.New(color.FgBlue)
	boldblue := blue.Add(color.Bold)
	boldblue.Println(table)
	return
}

func Level2(excelPath string) {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}
	rows, err := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex())) //读取全部数据
	table, err := gotable.Create("Host", "Ip", "Port", "Domain", "Title")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
	}
	i := 0
	for _, row := range rows {
		if i == 0 {
			i = i + 1
			continue
		}
		err := table.AddRow([]string{row[0], row[1], row[2], row[4], row[5]})
		if err != nil {
			return
		}
	}
	blue := color.New(color.FgBlue)
	boldblue := blue.Add(color.Bold)
	boldblue.Println(table)
	return
}

func Level3(excelPath string) {
	xlsx, err := excelize.OpenFile(excelPath)
	if err != nil {
		fmt.Println("open excel error,", err.Error())
		os.Exit(1)
	}
	rows, err := xlsx.GetRows(xlsx.GetSheetName(xlsx.GetActiveSheetIndex())) //读取全部数据
	table, err := gotable.Create("Host", "Ip", "Port", "Server", "Domain", "Title", "Country")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
	}
	i := 0
	for _, row := range rows {
		if i == 0 {
			i = i + 1
			continue
		}
		err := table.AddRow([]string{row[0], row[1], row[2], row[3], row[4], row[5], row[6]})
		if err != nil {
			return
		}
	}
	blue := color.New(color.FgBlue)
	boldblue := blue.Add(color.Bold)
	boldblue.Println(table)
	return
}
