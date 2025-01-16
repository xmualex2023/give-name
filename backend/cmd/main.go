package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/give-names/backend/internal/api/handler"
	"github.com/give-names/backend/internal/config"
	"github.com/give-names/backend/internal/service/name"
	"github.com/give-names/backend/pkg/gemini"
)

var (
	configPath string
	help       bool
)

func init() {
	// 定义命令行参数
	flag.StringVar(&configPath, "config", "", "配置文件路径 (默认: ./configs/config.yaml)")
	flag.BoolVar(&help, "help", false, "显示帮助信息")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `名字生成服务

用法: %s [选项]

选项:
`, os.Args[0])
	flag.PrintDefaults()
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 显示帮助信息
	if help {
		flag.Usage()
		os.Exit(0)
	}

	// 加载配置
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建 Gemini 客户端
	geminiClient, err := gemini.NewClient(&cfg.Gemini, &cfg.Proxy)
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer geminiClient.Close()

	// 创建名字生成服务
	nameService := name.NewService(geminiClient)

	// 创建 HTTP 处理器
	nameHandler := handler.NewNameHandler(nameService)

	// 设置路由
	router := gin.Default()

	// 配置 CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 注册路由
	nameHandler.RegisterRoutes(router)

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		// 启动服务器
		if err := router.Run(":" + cfg.Server.Port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 等待中断信号
	<-quit
	log.Println("Shutting down server...")
}
