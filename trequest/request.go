package trequest

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/FeatherStark/ttutils/trandom"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// HttpRequestConfig HTTP 请求配置
type HttpRequestConfig struct {
	Method         string   // Method 请求方法
	URI            string   // URI 请求Uri路径
	Data           string   // Data 请求数据
	Following      int      // Following 重定向次数
	FollowRedirect bool     // FollowRedirect 是否跟随重定向
	Header         sync.Map // Header 请求头
	Timeout        int      // Timeout 请求超时时间
	VerifyTls      bool     // VerifyTls 是否验证证书
	Proxy          string   // Proxy 代理地址
}

// HttpResponse HTTP 响应
type HttpResponse struct {
	URL            string // URL 请求地址
	Utf8Html       string // Utf8Html UTF8 解码后的HTML
	HeaderString   string // HeaderString 响应头字符串
	*http.Response        // HTTP响应结构体
}

// HttpRequestGetConfig GET 请求。
// Args: uri string 请求地址。
// Returns: *HttpRequestConfig HTTP 请求配置。
func HttpRequestGetConfig(uri string) *HttpRequestConfig {
	return HttpNewRequestConfig("GET", uri, "")
}

// HttpRequestPostConfig POST 请求  默认的 Content-Type 是 application/x-www-form-urlencoded。
// Args: uri string 请求地址, data string 请求数据。
// Returns: *HttpRequestConfig HTTP 请求配置。
func HttpRequestPostConfig(uri string) *HttpRequestConfig {
	cfg := HttpNewRequestConfig("POST", uri, "")
	cfg.Header.Store("Content-Type", "application/x-www-form-urlencoded")
	return cfg
}

// HttpNewRequestConfig 创建请求配置。
// Args: method string 请求方法, uri string 请求地址, data string 请求数据。
// Returns: *HttpRequestConfig HTTP 请求配置。
func HttpNewRequestConfig(method, uri, data string) *HttpRequestConfig {
	cfg := &HttpRequestConfig{
		FollowRedirect: false,
		Following:      3,
		Timeout:        15,
		VerifyTls:      false,
	}
	cfg.Header.Store("User-Agent", GetRandomUserAgent(nil))
	cfg.Data = data
	cfg.Method = method
	cfg.URI = uri
	return cfg
}

// createHttpClient 创建 HTTP 客户端。
// Args: config *HttpRequestConfig 请求配置。
// Returns: *http.Client HTTP 客户端, error 错误。
func createHttpClient(config *HttpRequestConfig) (*http.Client, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !config.VerifyTls},
	}

	if config.Proxy != "" {
		proxyURL, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, fmt.Errorf("解析代理地址失败: %v", err)
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	client := &http.Client{
		Timeout:   time.Duration(config.Timeout) * time.Second,
		Transport: transport,
	}

	if !config.FollowRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	return client, nil
}

// setHttpRequestHeaders 设置请求头。
// Args: req *http.Request 请求, header sync.Map 请求头。
func setHttpRequestHeaders(req *http.Request, header sync.Map) {
	header.Range(func(key, value interface{}) bool {
		req.Header.Set(key.(string), value.(string))
		return true
	})
}

// buildHttpResponse 构造 HttpResponse。
// Args: resp *http.Response HTTP 响应。
// Returns: *HttpResponse 构造后的 HttpResponse, error 错误。
func buildHttpResponse(resp *http.Response) (*HttpResponse, error) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %v", err)
	}

	var headerBuilder strings.Builder
	for key, values := range resp.Header {
		for _, value := range values {
			headerBuilder.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}
	}

	return &HttpResponse{
		URL:          resp.Request.URL.String(),
		Utf8Html:     string(body),
		HeaderString: headerBuilder.String(),
		Response:     resp,
	}, nil
}

// DoHttpRequest 执行 HTTP 请求。
// Args: host string 主机地址, config *HttpRequestConfig 请求配置。
// Returns: *HttpResponse HTTP 响应, error 错误。
func DoHttpRequest(host string, config *HttpRequestConfig) (*HttpResponse, error) {
	host = GetHost(host)
	client, err := createHttpClient(config)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	if config.Method == http.MethodPost || config.Method == http.MethodPut {
		req, err = http.NewRequest(config.Method, host+config.URI, bytes.NewBufferString(config.Data))
	} else {
		req, err = http.NewRequest(config.Method, host+config.URI, nil)
	}
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}
	// 设置自定义请求头
	setHttpRequestHeaders(req, config.Header)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("发送请求失败: %v", err)
	}
	return buildHttpResponse(resp)
}

// GetRandomUserAgent 传入UserAgent列表，获取随机的 User-Agent;如果传入为nil，则使用内置随机UA头。
// Args: userAgents []string 自定义UserAgent列表。
// Returns: string 随机的User-Agent。
func GetRandomUserAgent(userAgents []string) string {
	if len(userAgents) == 0 {
		userAgents = []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.75 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.1.2.86",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:65.0) Gecko/20100101 Firefox/65.0",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15",
		}
	}
	return userAgents[trandom.RandomNumber(0, len(userAgents)-1)]
}

// GetHost 获取主机地址,协议+域名或IP。
// Args: host string 主机地址。
// Returns: string 主机地址。
func GetHost(host string) string {
	u, err := url.Parse(host)
	if err != nil {
		return ""
	}
	return u.Scheme + "://" + u.Host
}
