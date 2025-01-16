package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Gemini GeminiConfig `mapstructure:"gemini"`
	CORS   CORSConfig   `mapstructure:"cors"`
	Proxy  ProxyConfig  `mapstructure:"proxy"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port    string        `mapstructure:"port"`
	Timeout time.Duration `mapstructure:"timeout"`
}

// GeminiConfig Gemini API 配置
type GeminiConfig struct {
	APIKey     string        `mapstructure:"api_key"`
	Model      string        `mapstructure:"model"`
	Timeout    time.Duration `mapstructure:"timeout"`
	UseProxy   bool          `mapstructure:"use_proxy"`   // 是否使用代理
	DirectMode bool          `mapstructure:"direct_mode"` // 是否强制直连模式
}

// ProxyConfig 代理配置
type ProxyConfig struct {
	HTTPProxy  string `mapstructure:"http_proxy"`
	HTTPSProxy string `mapstructure:"https_proxy"`
	Enabled    bool   `mapstructure:"enabled"`
}

// CORSConfig CORS 配置
type CORSConfig struct {
	AllowedOrigins []string `mapstructure:"allowed_origins"`
	AllowedMethods []string `mapstructure:"allowed_methods"`
}

// LoadConfig 加载配置
// configPath 可以是配置文件的路径，如果为空则使用默认路径
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// 设置默认值
	setDefaults(v)

	// 如果提供了配置文件路径
	if configPath != "" {
		// 获取配置文件的目录和文件名
		dir, file := filepath.Split(configPath)
		ext := filepath.Ext(file)
		name := file[:len(file)-len(ext)]

		// 设置配置文件信息
		v.SetConfigName(name)
		v.SetConfigType(ext[1:]) // 去掉点号
		v.AddConfigPath(dir)
	} else {
		// 使用默认配置路径
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath("./configs")
		v.AddConfigPath(".")
	}

	// 设置环境变量前缀
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// 解析配置到结构体
	config := &Config{}
	if err := v.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	// 验证配置
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	return config, nil
}

// setDefaults 设置默认配置值
func setDefaults(v *viper.Viper) {
	v.SetDefault("server.port", "8080")
	v.SetDefault("server.timeout", "30s")
	v.SetDefault("gemini.model", "gemini-pro")
	v.SetDefault("gemini.timeout", "20s")
	v.SetDefault("gemini.use_proxy", true)
	v.SetDefault("gemini.direct_mode", false)
	v.SetDefault("proxy.enabled", true)
	v.SetDefault("proxy.http_proxy", "http://127.0.0.1:8118")
	v.SetDefault("proxy.https_proxy", "http://127.0.0.1:8118")
	v.SetDefault("cors.allowed_origins", []string{"http://localhost:3000"})
	v.SetDefault("cors.allowed_methods", []string{"GET", "POST", "OPTIONS"})
}

// validateConfig 验证配置是否有效
func validateConfig(cfg *Config) error {
	if cfg.Gemini.APIKey == "" {
		return fmt.Errorf("gemini API key is required")
	}
	if cfg.Server.Port == "" {
		return fmt.Errorf("server port is required")
	}
	if cfg.Server.Timeout <= 0 {
		return fmt.Errorf("invalid server timeout")
	}
	if cfg.Gemini.Timeout <= 0 {
		return fmt.Errorf("invalid gemini timeout")
	}
	if cfg.Proxy.Enabled && (cfg.Proxy.HTTPProxy == "" || cfg.Proxy.HTTPSProxy == "") {
		return fmt.Errorf("proxy URLs are required when proxy is enabled")
	}
	return nil
}
