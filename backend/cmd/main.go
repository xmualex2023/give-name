package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/xmualex2023/give-name/internal/ai"
	"github.com/xmualex2023/give-name/internal/handler"
	"github.com/xmualex2023/give-name/internal/service"
)

func main() {
	// 初始化 Gemini 服务
	geminiService, err := ai.NewGeminiService("AIzaSyA3_z2C2AVBEUXYs6Fn6vTnspwf_dd2h8E")
	if err != nil {
		log.Fatalf("Failed to initialize Gemini service: %v", err)
	}

	// 初始化名字服务
	nameService := service.NewNameService()

	// 初始化处理器
	nameHandler := handler.NewNameHandler(nameService, geminiService)

	// 创建 Gin 引擎
	r := gin.Default()

	// 配置 CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	// 设置路由
	setupRoutes(r, nameHandler)

	// 启动服务器
	log.Fatal(r.Run(":8000"))
}

func setupRoutes(r *gin.Engine, nameHandler *handler.NameHandler) {
	api := r.Group("/api")
	{
		api.POST("/generate", nameHandler.GenerateName)
	}
}
