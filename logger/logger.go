package logger

import (
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	pWARN        = "\033[1;31m[!]\033[0m "
	pINFO        = "\033[1;33m[+]\033[0m "
	pDEBUG       = "\033[1;34m[-]\033[0m "
	pSUCCESS     = "\033[1;32m[*]\033[0m "
	isWin        = false
	usageCommand = "Usage:\n  [!]./fofa  -q 'host=\"sdxiehe.edu.cn\"' -s 10000 -o data.xlsx\n" +
		"  [!]./fofa  -f query_rules_file.txt -s 10000 -o data.xlsx\n"
	usageOptions = `Options:
  -h, --help
  -q, --query       [string]          参数字符串 (默认: '')
  -f, --file        [filePath]        批量查询规则文件 (默认: '')
  -s, --size        [int]             到处数据量 (默认: 10000)
  -e, --is_honeypot                   排除蜜罐数据  (仅限FOFA高级会员使用 ) 
  -o, --output      [string]          输出文件名字 / 绝对路径 (默认: data.xlsx)
  -g, --grammar                       fofa搜索语法帮助表
  -t, --tip         [string]          fofa 搜索关键字提示列表
  -x, --xlsx        [文件路径]	      读取生成xlsx文件 (配合 level 参数使用)
  -l, --level       [1/2/3]    		  读取文件细粒度 (配合 xlsx 参数使用)
  -ih, --iconhash   [string]          计算指定 URL favicon icon_hash      
`
	fofaGrammar = `
[+]                Rule                                   Mark               
 ------------------------------------------- -------------------------------- 
  domain="qq.com"                             搜索根域名带有qq.com的网站                
  host=".gov.cn"                              从host中搜索".gov.cn"               
  ip="1.1.1.1"                                从ip中搜索包含"1.1.1.1"的网站            
  ip="220.181.111.1/24"                       查询IP为"220.181.111.1"的C段资产       
  port="6379"                                 查找对应"6379"端口的资产                 
  title="beijing"                             从标题中搜索"北京"                      
  status_code="402"                           查询服务器状态为"402"的资产                
  protocol="quic"                             查询quic协议资产                      
  header="elastic"                            从http头中搜索"elastic"              
  body="网络空间测绘"                            从html正文中搜索"网络空间测绘"              
  os="centos"                                 搜索CentOS资产                      
  server=="Microsoft-IIS/10"                  搜索IIS10服务器                      
  app="Microsoft-Exchange"                    搜索Microsoft-Exchange设备          
  base_protocol="udp"                         搜索指定udp协议的资产                    
  banner=users && protocol=ftp                搜索FTP协议中带有users文本的资产            
  icp="京ICP证030173号"                        查找备案号为"京ICP证030173号"的网站         
  icon_hash="-247388890"                      搜索使用此icon的资产(VIP)               
  js_name="js/jquery.js"                      查找网站正文中包含js/jquery.js的资产        
  js_md5="82ac3f14327a8b7ba49baa208d4eaa15"   查找js源码与之匹配的资产                   
  type=service                                搜索所有协议资产，支持subdomain和service两种  
  is_domain=true                              搜索域名的资产                         
  ip_ports="80,161"                           搜索同时开放80和161端口的ip               
  port_size="6"                               查询开放端口数量等于"6"的资产(VIP)           
  port_size_gt="6"                            查询开放端口数量大于"6"的资产(VIP)           
  port_size_lt="12"                           查询开放端口数量小于"12"的资产(VIP)          
  is_ipv6=true                                搜索ipv6的资产                       
  is_fraud=false                              排除仿冒/欺诈数据                       
  is_honeypot=false                           排除蜜罐数据(VIP)                     
  country="CN"                                搜索指定国家(编码)的资产                   
  region="Xinjiang"                           搜索指定行政区的资产                      
  city="Ürümqi"                               搜索指定城市的资产                       
  asn="19551"                                 搜索指定asn的资产                      
  org="Amazon.com,Inc."                       搜索指定org(组织)的资产                  
  cert="baidu"                                搜索证书(https或者imaps等)中带有baidu的资产  
  cert.subject="Oracle Corporation"           搜索证书持有者是OracleCorporation的资产    
  cert.issuer="DigiCert"                      搜索证书颁发者为DigiCertInc的资产          
  cert.is_valid=true                          验证证书是否有效,true有效,false无效(VIP)    
  after="2017" && before="2017-10-01"         限定时间范围
  
[+] 高级搜索：可以使用括号 () / 和 && / 或 || / 完全匹配 == / 不为 != 等逻辑运算符。
`
)

func InitPlatform() {
	if strings.HasSuffix(strings.ToLower(os.Args[0]), ".exe") {
		pWARN = "[!] "
		pINFO = "[+] "
		pSUCCESS = "[*] "
		pDEBUG = "[-] "
		isWin = true
	}
	usageCommand = strings.Replace(usageCommand, "./fofa", os.Args[0], 2)
	return
}

func randColor() string {
	rand.Seed(time.Now().UnixNano())
	colorNum := rand.Intn(6) + 31
	return strconv.Itoa(colorNum)
}

func AsciiBanner() {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)

	if isWin != true {
		boldRed.Println("\033[1;" + randColor() + "m" + asciiBanner2 + "\033[0m")
	} else {
		boldRed.Println(asciiBanner2)
	}
}

func Usage() {
	yellow := color.New(color.FgHiYellow)
	boldyellow := yellow.Add(color.Bold)
	boldyellow.Println(usageCommand)
	boldyellow.Println(usageOptions)
}

func FofaGrammar() {
	blue := color.New(color.FgBlue)
	boldblue := blue.Add(color.Bold)
	boldblue.Println(fofaGrammar)
}

func Warn(format string, args ...interface{}) {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRed.Fprintf(os.Stderr, pWARN+format+"\n", args...)
}

func Info(format string, args ...interface{}) {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRed.Fprintf(os.Stdout, pINFO+format+"\n", args...)
}

func Success(format string, args ...interface{}) {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRed.Fprintf(os.Stdout, pSUCCESS+format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
	boldRed.Fprintf(os.Stdout, pDEBUG+format+"\n", args...)
}
