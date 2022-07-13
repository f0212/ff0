package fetch

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"fofa/logger"
	"fofa/option"
	"github.com/buger/jsonparser"
	"github.com/projectdiscovery/retryablehttp-go"
	"github.com/twmb/murmur3"
	"hash"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	errorUnauthorized = errors.New("401 未经授权，请确保电子邮件和apikey正确无误。")
	errorForbidden    = errors.New("403 禁止，无法正常访问Fofa服务。")
	error503          = errors.New("503 服务暂时不可用。")
	error502          = errors.New("502 网关失效。")
	fields            = "host,ip,port,server,domain,title,country,province,city,icp,protocol"
	FetchResult       = make(map[string]([][]string))
)

var FetchResultT = struct {
	M map[string]([][]string)
	sync.Mutex
}{M: make(map[string]([][]string))}

type Fofa struct {
	email  string
	apiKey string
	size   string
	*retryablehttp.Client
}

func FofaRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		//不要在上下文上重试
		if ctx.Err() != nil {
			return false, ctx.Err()
		}

		if resp.StatusCode == 503 || resp.StatusCode == 409 {
			return true, nil
		}

		return false, nil
	}
}

func RequestLogHook(req *http.Request, i int) {
	logger.Debug("--> REQ %v %v", i, req.URL)
}

func ResponseLogHook(resp *http.Response) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn(err.Error())
	}
	logger.Debug("<-- RESP %v", resp.StatusCode, string(content))
}

func NewFofaClient(email, apiKey, size string) *Fofa {

	Client := retryablehttp.NewClient(retryablehttp.Options{
		// 等待重试的最短时间
		RetryWaitMin: 700 * time.Millisecond,
		// 等待重试的最长时间
		RetryWaitMax: 3500 * time.Millisecond,
		// 超时等待请求的最长时间
		Timeout: 14000 * time.Millisecond,
		// 最大重试次数
		RetryMax: 6,
		// 读取的最大HTTP响应大小
		// 重用连接
		RespReadLimit: 4096,
		// 详细指定是否应打印调试消息
		Verbose: true,
		// 指定是否杀死所有保持活动的连接
		KillIdleConn: true,
	})
	Client.CheckRetry = FofaRetryPolicy()
	//Client.RequestLogHook = RequestLogHook
	//Client.ResponseLogHook = ResponseLogHook

	return &Fofa{
		email:  email,
		apiKey: apiKey,
		size:   size,
		Client: Client,
	}
}

func (ff *Fofa) Get(u string) ([]byte, error) {

	body, err := ff.Client.Get(u)
	if err != nil {
		return nil, err
	}
	defer body.Body.Close()

	if body.StatusCode == 503 {
		return nil, error503
	} else if body.StatusCode == 502 {
		return nil, error502
	} else if body.StatusCode == 403 {
		return nil, errorForbidden
	}

	content, err := ioutil.ReadAll(body.Body)

	if err != nil {
		return nil, err
	}

	return content, nil
}

//单挑参数查询
func (ff *Fofa) Query(query string) (resultArray [][]string) {

	base64Query := base64.StdEncoding.EncodeToString([]byte(query))
	queryUrl := fmt.Sprintf("https://fofa.info/api/v1/search/all?email=%s&key=%s&qbase64=%s&size=%s&fields=%s&is_honeypot=%s",
		ff.email, ff.apiKey, base64Query, ff.size, fields, option.Is_honeypot)
	body, err := ff.Get(queryUrl)

	if err != nil {
		logger.Warn("<< " + query + " " + err.Error())
		return
	}

	errmsg, err := jsonparser.GetString(body, "errmsg")

	if errmsg != "" {
		logger.Warn("<< " + query + " " + errors.New(errmsg).Error())
		return
	}

	results, _, _, err := jsonparser.Get(body, "results")

	if err != nil {
		logger.Warn("<< " + query + " " + err.Error())
		return
	}

	err = json.Unmarshal(results, &resultArray)
	if err != nil {
		logger.Warn("<< " + query + " " + err.Error())
		return
	}

	return
}

//进行账号认证
func (ff *Fofa) Auth() (valid bool, err error) {

	valid = false
	err = nil

	authUrl := fmt.Sprintf("https://fofa.info/api/v1/info/my?email=%s&key=%s", ff.email, ff.apiKey)

	body, err := ff.Get(authUrl)

	if err != nil {
		return
	}

	body_str := string(body)

	if strings.Contains(body_str, "fofa_server") {
		valid = true
		return
	} else {
		if strings.Contains(body_str, "401") {
			err = errorUnauthorized
		} else if strings.Contains(body_str, "403") {
			err = errorForbidden
		}
		return
	}
}

//tip查询
func (ff *Fofa) QueryAll(querys []string) {

	logger.Info(fmt.Sprintf("共有 %v 条规则进行查询。", len(querys)))

	for i, q := range querys {
		logger.Info(fmt.Sprintf("现在查询 (%v, %v)。", i, q))
		FetchResult[q] = ff.Query(q)
	}

	logger.Success(fmt.Sprintf("全部 (%v) 条查询完成。", len(querys)))
}

//入口查询
func (ff *Fofa) QueryAllT(querys []string) {

	logger.Info(fmt.Sprintf("共有 %v 条规则进行查询。", len(querys)))
	//线程
	maxNoGoroutines := 10
	concurrentGoroutines := make(chan struct{}, maxNoGoroutines)
	for i := 0; i < maxNoGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}

	done := make(chan bool)
	waitForAllJobs := make(chan bool)

	go func() {
		for i := 0; i < len(querys); i++ {
			<-done
			concurrentGoroutines <- struct{}{}
		}
		waitForAllJobs <- true
	}()
	//进行查询
	for i, q := range querys {
		time.Sleep(300 * time.Millisecond)
		<-concurrentGoroutines

		go func(pos int, query string) {
			logger.Info(fmt.Sprintf("-> %v, %v", pos, query))
			qresult := ff.Query(query)
			FetchResultT.Lock()
			FetchResultT.M[query] = qresult
			FetchResultT.Unlock()
			if len(qresult) == 0 {
				logger.Warn("<- %v, %v 没有找到有关信息", pos, query)
			} else {
				logger.Success(fmt.Sprintf("<- %v, %v", pos, query))
			}
			done <- true
		}(i, q)
	}

	<-waitForAllJobs
	logger.Success(fmt.Sprintf("全部 (%v) 条参数查询完成。", len(querys)))

}

func (ff *Fofa) IconHash(url string) (iconHash string) {

	content, err := ff.Get(url)

	if err != nil {
		logger.Warn(err.Error())
		return
	}

	b64s := base64.StdEncoding.EncodeToString(content)

	var buffer bytes.Buffer

	for i := 0; i < len(b64s); i++ {
		ch := b64s[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}

	buffer.WriteByte('\n')

	var h32 hash.Hash32 = murmur3.New32()
	h32.Write(buffer.Bytes())
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}

func GetIconHash(url string) {

	logger.Info(url)

	iconHash := NewFofaClient("", "", "").IconHash(url)

	logger.Success("icon_hash=\"%v\"", iconHash)
}

//tip查询
func (ff *Fofa) FofaTip(keyword string) (dropList []string) {

	tipUrl := "https://api.fofa.info/v1/search/tip?q=" + keyword

	content, err := ff.Get(tipUrl)

	if err != nil {
		logger.Warn(err.Error())
		return nil
	}

	regexp := regexp.MustCompile(`"name":"[^"]+`)
	results := regexp.FindAllStringSubmatch(string(content), -1)

	for _, v := range results {
		dropList = append(dropList, fmt.Sprintf("%v", v[0][8:]))
	}

	return
}

func GetFofaTip(keyword string) {

	logger.Info(fmt.Sprintf("Keyword: %v", keyword))

	dropList := NewFofaClient("", "", "").FofaTip(keyword)

	if dropList == nil {
		logger.Warn("未找到任何内容，请尝试其他关键字。")
		return
	}

	for _, v := range dropList {
		logger.Success(v)
	}
}
