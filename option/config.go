package option

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//解析yml文件
type BaseInfo struct {
	DefaultEmail  string `yaml:"DefaultEmail"`
	DefaultAPIKey string `yaml:"DefaultAPIKey"`
	DefaultSize   string `yaml:"DefaultSize"`
	DefaultOutput string `yaml:"DefaultOutput"`
}

func (c *BaseInfo) GetConf() *BaseInfo {
	////对文件进行判断
	//if _, err := os.Stat(".././config.yml"); err != nil {
	//	if os.IsNotExist(err) {
	//		// 如果文件不存在
	//		fp, _ := os.Create(".././config.yml")
	//		fp.Close()
	//	}
	//}
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}

func Config() (DefaultEmail string, DefaultAPIKey string, DefaultSize string, DefaultOutput string) {
	info := BaseInfo{}
	conf := info.GetConf()
	//fmt.Println(conf.DefaultEmail)
	return conf.DefaultEmail, conf.DefaultAPIKey, conf.DefaultSize, conf.DefaultOutput
}
