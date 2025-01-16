package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/give-names/backend/internal/config"
	"github.com/give-names/backend/internal/model"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Client Gemini API 客户端
type Client struct {
	client *genai.Client
	model  *genai.GenerativeModel
	config *config.GeminiConfig
}

// NewClient 创建新的 Gemini 客户端
func NewClient(cfg *config.GeminiConfig, proxyCfg *config.ProxyConfig) (*Client, error) {
	ctx := context.Background()

	// 验证 API Key
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("gemini API key is required")
	}

	// 创建自定义 HTTP 客户端
	httpClient := &http.Client{
		Timeout: cfg.Timeout,
	}

	// 配置代理
	if cfg.UseProxy && proxyCfg.Enabled && !cfg.DirectMode {
		// 设置环境变量
		if err := setProxyEnv(proxyCfg); err != nil {
			return nil, fmt.Errorf("failed to set proxy environment: %v", err)
		}

		// 创建代理传输
		transport, err := createProxyTransport(proxyCfg)
		if err != nil {
			return nil, fmt.Errorf("failed to create proxy transport: %v", err)
		}

		// 设置自定义传输
		httpClient.Transport = transport
	} else {
		// 清除代理环境变量
		clearProxyEnv()
		// 使用默认传输
		httpClient.Transport = &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			ResponseHeaderTimeout: cfg.Timeout,
			TLSHandshakeTimeout:   10 * time.Second,
			DisableKeepAlives:     false,
			MaxIdleConnsPerHost:   10,
		}
	}

	// 创建带认证头的 RoundTripper
	authTransport := &authRoundTripper{
		base:   httpClient.Transport,
		apiKey: cfg.APIKey,
	}
	httpClient.Transport = authTransport

	// 创建客户端选项
	opts := []option.ClientOption{
		option.WithoutAuthentication(), // 禁用默认认证
		option.WithHTTPClient(httpClient),
	}

	// 创建客户端
	client, err := genai.NewClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %v", err)
	}

	// 创建生成模型
	model := client.GenerativeModel(cfg.Model)
	model.SetTemperature(0.7)      // 设置温度
	model.SetTopK(40)              // 设置 top-k
	model.SetTopP(0.95)            // 设置 top-p
	model.SetMaxOutputTokens(2048) // 设置最大输出 token

	return &Client{
		client: client,
		model:  model,
		config: cfg,
	}, nil
}

// authRoundTripper 是一个自定义的 RoundTripper，用于添加认证头
type authRoundTripper struct {
	base   http.RoundTripper
	apiKey string
}

// RoundTrip 实现 http.RoundTripper 接口
func (a *authRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// 添加 API Key 到 URL 查询参数
	q := req.URL.Query()
	q.Set("key", a.apiKey)
	req.URL.RawQuery = q.Encode()

	// 打印请求信息（调试用）
	fmt.Printf("Debug - Request URL: %s\n", req.URL.String())
	fmt.Printf("Debug - API Key: %s\n", a.apiKey)
	fmt.Printf("Debug - Headers: %v\n", req.Header)

	// 使用基础 RoundTripper 发送请求
	resp, err := a.base.RoundTrip(req)
	if err != nil {
		fmt.Printf("Debug - Request Error: %v\n", err)
		return nil, err
	}

	// 打印响应状态（调试用）
	fmt.Printf("Debug - Response Status: %s\n", resp.Status)
	return resp, nil
}

// createProxyTransport 创建代理传输
func createProxyTransport(cfg *config.ProxyConfig) (*http.Transport, error) {
	proxyURL, err := url.Parse(cfg.HTTPProxy)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy URL: %v", err)
	}

	return &http.Transport{
		Proxy:                 http.ProxyURL(proxyURL),
		ResponseHeaderTimeout: 30 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		DisableKeepAlives:     false,
		MaxIdleConnsPerHost:   10,
	}, nil
}

// setProxyEnv 设置代理环境变量
func setProxyEnv(cfg *config.ProxyConfig) error {
	if err := os.Setenv("HTTP_PROXY", cfg.HTTPProxy); err != nil {
		return fmt.Errorf("failed to set HTTP_PROXY: %v", err)
	}
	if err := os.Setenv("HTTPS_PROXY", cfg.HTTPSProxy); err != nil {
		return fmt.Errorf("failed to set HTTPS_PROXY: %v", err)
	}
	// 设置不代理的地址
	if err := os.Setenv("NO_PROXY", "localhost,127.0.0.1"); err != nil {
		return fmt.Errorf("failed to set NO_PROXY: %v", err)
	}
	return nil
}

// clearProxyEnv 清除代理环境变量
func clearProxyEnv() {
	os.Unsetenv("HTTP_PROXY")
	os.Unsetenv("HTTPS_PROXY")
	os.Unsetenv("NO_PROXY")
}

// GenerateNames 生成中文名字建议
func (c *Client) GenerateNames(ctx context.Context, req *model.NameRequest) (*model.NameResponse, error) {
	if req.EnglishName == "" {
		return nil, fmt.Errorf("english name is required")
	}

	prompt := fmt.Sprintf(`Generate 3 Chinese names for the English name: %s

Requirements:
1. Names should be phonetically similar to the English name
2. Each name should have deep cultural meaning
3. Each name should be 2-3 characters long

IMPORTANT: Return ONLY a JSON object without any markdown formatting or code blocks. The response must strictly follow this structure:
{
    "suggestions": [
        {
            "chinese_name": "米凯尔",
            "pinyin": "mi kai er",
            "characters": [
                {
                    "character": "米",
                    "pinyin": "mi"
                }
            ],
            "meaning": "Brief meaning",
            "cultural_notes": "Brief cultural significance",
            "personality": "Key personality traits",
            "english_intro": "Brief English explanation"
        }
    ]
}

Rules:
1. Each character in the name must have its own entry in the "characters" array
2. Pinyin must be in lowercase with spaces between syllables
3. All text fields must be concise and clear
4. Do not include any explanation or comments outside the JSON structure
5. Ensure the JSON is valid and properly formatted`, req.EnglishName)

	// 设置超时上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	// 生成内容
	resp, err := c.model.GenerateContent(timeoutCtx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %v", err)
	}

	// 增加响应验证
	if resp == nil || len(resp.Candidates) == 0 {
		return nil, fmt.Errorf("empty response from Gemini")
	}

	candidate := resp.Candidates[0]
	if candidate.Content == nil || len(candidate.Content.Parts) == 0 {
		return nil, fmt.Errorf("invalid response format from Gemini")
	}

	// 获取文本内容
	text := candidate.Content.Parts[0].(genai.Text)
	if text == "" {
		return nil, fmt.Errorf("empty text response from Gemini")
	}

	fmt.Printf("Debug - Raw response:\n%s\n", text)

	// 清理响应文本，移除可能的 markdown 格式
	cleanText := cleanResponse(string(text))
	fmt.Printf("Debug - Cleaned response:\n%s\n", cleanText)

	// 解析 JSON
	var response model.NameResponse
	if err := json.Unmarshal([]byte(cleanText), &response); err != nil {
		// 尝试格式化 JSON 以便于调试
		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, []byte(cleanText), "", "    "); err == nil {
			fmt.Printf("Debug - Formatted JSON:\n%s\n", prettyJSON.String())
		}
		return nil, fmt.Errorf("failed to parse response: %v\nraw response: %s", err, cleanText)
	}

	// 验证响应内容
	if len(response.Suggestions) == 0 {
		return nil, fmt.Errorf("no name suggestions in response")
	}

	fmt.Printf("Debug - Successfully parsed %d suggestions\n", len(response.Suggestions))

	return &response, nil
}

// cleanResponse 清理响应文本，移除 markdown 格式
func cleanResponse(text string) string {
	// 移除可能的代码块标记
	text = removeMarkdownCodeBlocks(text)
	// 移除多余的空白字符
	text = removeExtraWhitespace(text)
	return text
}

// removeMarkdownCodeBlocks 移除 markdown 代码块
func removeMarkdownCodeBlocks(text string) string {
	// 如果文本以 ``` 开头和结尾，移除它们
	if len(text) > 6 && text[:3] == "```" {
		if idx := len(text) - 3; idx > 3 && text[idx:] == "```" {
			// 移除开头的 ```language 标记
			firstNewline := 3
			for i := 3; i < len(text); i++ {
				if text[i] == '\n' {
					firstNewline = i + 1
					break
				}
			}
			return text[firstNewline:idx]
		}
	}
	return text
}

// removeExtraWhitespace 移除多余的空白字符
func removeExtraWhitespace(text string) string {
	// 移除开头和结尾的空白字符
	text = strings.TrimSpace(text)
	return text
}

// Close 关闭客户端
func (c *Client) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
