package option

import (
	"errors"
	"fofa/xlsx"
)

var (
	PrintUsage        = errors.New("")
	PrintGrammar      = errors.New("Grammar")
	IconHash          = errors.New("IconHash")
	FofaTip           = errors.New("FofaTip")
	errorNoQueryFound = errors.New("未找到查询语句或规则文件，必须指定 -q 或 -f 参数。")
	errorNolevelFound = errors.New("-x 需和 -l 配合使用")
	Is_honeypot       = "true"
	errorOver         = errors.New("查询结束")
)

func ParseCli(args []string) (
	email string,
	apiKey string,
	query string,
	ruleFile string,
	size string,
	output string,
	xlsxs string,
	level string,
	err error) {
	email, apiKey, size, output = Config()
	if len(args) == 0 {
		err = PrintUsage
		return
	}
	for pos := 0; pos < len(args); pos++ {
		switch args[pos] {
		case "-e", "--is_honeypot":
			Is_honeypot = "false"
		case "-k", "--key":
			if pos+1 < len(args) {
				apiKey = args[pos+1]
				pos++
			}
		case "-q", "--query":
			if pos+1 < len(args) {
				query = args[pos+1]
				pos++
			}
		case "-f", "--file":
			if pos+1 < len(args) {
				ruleFile = args[pos+1]
				pos++
			}
		case "-s", "--size":
			if pos+1 < len(args) {
				size = args[pos+1]
				pos++
			}
		case "-o", "--output":
			if pos+1 < len(args) {
				output = args[pos+1]
				pos++
			}
		case "-t", "--tip":
			if pos+1 < len(args) {
				FofaTip = errors.New(args[pos+1])
				err = FofaTip
				return
			}
		case "-ih", "--iconhash":
			if pos+1 < len(args) {
				IconHash = errors.New(args[pos+1])
				err = IconHash
				return
			}
		case "-h", "--help":
			err = PrintUsage
			return
		case "-g", "--grammar":
			err = PrintGrammar
			return
		case "-x", "--xlsx":
			if pos+1 < len(args) {
				xlsxs = args[pos+1]
				pos++
			}
		case "-l", "--level":
			if pos+1 < len(args) {
				level = args[pos+1]
				pos++
			}
		}
	}
	//全部为空时
	if xlsxs == "" && level == "" {
		if query == "" && ruleFile == "" {
			err = errorNoQueryFound
			return
		}
		return
	}
	//读取文件不不全时
	if xlsxs == "" || level == "" {
		err = errorNolevelFound
		return
	}
	//读取文件成立时
	if xlsxs != "" && level != "" {
		xlsx.Xlsx(xlsxs, level)
		err = errorOver
		return
	}

	return
}

func Honeypot() string {
	return Is_honeypot
}
