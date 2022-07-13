package option

import (
	"errors"
)

var (
	PrintUsage        = errors.New("")
	PrintGrammar      = errors.New("Grammar")
	IconHash          = errors.New("IconHash")
	FofaTip           = errors.New("FofaTip")
	errorNoQueryFound = errors.New("未找到查询语句或规则文件，必须指定 -q 或 -f 参数。")
	Is_honeypot       = "true"
)

func ParseCli(args []string) (
	email string,
	apiKey string,
	query string,
	ruleFile string,
	size string,
	output string,
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
		}
	}

	if query == "" && ruleFile == "" {
		err = errorNoQueryFound
		return
	}

	return
}

func Honeypot() string {
	return Is_honeypot
}
