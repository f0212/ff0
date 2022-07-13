package main

import (
	"fmt"
	"fofa/fetch"
	"fofa/logger"
	"fofa/option"
	"fofa/report"
	"fofa/utils"
	"os"
)

func main() {
	logger.InitPlatform()
	logger.AsciiBanner()

	email, apiKey, query, ruleFile, size, output, err := option.ParseCli(os.Args[1:])

	if err != nil {
		if err == option.PrintUsage {
			logger.Usage()
		} else if err == option.PrintGrammar {
			logger.FofaGrammar()
		} else if err == option.IconHash {
			fetch.GetIconHash(err.Error())
		} else if err == option.FofaTip {
			fetch.GetFofaTip(err.Error())
		} else {
			logger.Warn(err.Error())
		}
		return
	}
	//fmt.Println(email)
	logger.Success(fmt.Sprintf("Email: %v", email))
	//logger.Success(fmt.Sprintf("Key: %v", apiKey))

	var querys []string

	if query != "" {
		logger.Success(fmt.Sprintf("查询语句: %v", query))
		querys = append(querys, query)
	} else {
		logger.Success(fmt.Sprintf("规则文件: %v", ruleFile))
		querys = utils.ScanFile(ruleFile)
	}

	logger.Success(fmt.Sprintf("提取数量: %v", size))
	logger.Success(fmt.Sprintf("输出路径: %v", output))

	clt := fetch.NewFofaClient(email, apiKey, size)

	vaild, err := clt.Auth()
	if vaild != true {
		if err != nil {
			logger.Warn(err.Error())
		} else {
			logger.Warn("空的")
		}
		return
	} else {
		logger.Success("账号认证成功！")
	}

	clt.QueryAllT(querys)

	report.WriteXlsx(fetch.FetchResultT.M, "../output/"+output)

	return
}
