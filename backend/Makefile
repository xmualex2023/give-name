.PHONY: all build run test clean lint help api-test

# 默认配置文件路径
CONFIG_FILE ?= ./configs/config.yaml

# API 测试参数
API_HOST ?= localhost:8080
TEST_NAME ?= Michael

# 默认目标
all: lint test build

# 构建可执行文件
build:
	@echo "Building..."
	@go build -o bin/server cmd/main.go

# 运行服务
run:
	@echo "Running server..."
	@go run cmd/main.go -config $(CONFIG_FILE)

# API 测试命令
api-test:
	@echo "Testing API with name: $(TEST_NAME)"
	@echo "Generate name API test:"
	@curl -X POST "http://$(API_HOST)/api/generate" \
		-H "Content-Type: application/json" \
		-d '{"english_name":"$(TEST_NAME)"}' | jq '.' || echo "Failed to parse JSON (jq not installed)"

# API 测试帮助
api-test-help:
	@echo "API Test Commands:"
	@echo "  make api-test                      - Test API with default name (Michael)"
	@echo "  make api-test TEST_NAME=John       - Test API with specific name"
	@echo "  make api-test API_HOST=localhost:9090 - Test API with different host"
	@echo "Example:"
	@echo "  make api-test TEST_NAME='John Smith'"

# 运行测试
test:
	@echo "Running tests..."
	@go test -v ./...

# 运行测试并生成覆盖率报告
cover:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html

# 代码格式化
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# 代码检查
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint is not installed. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run ./...; \
	fi

# 清理构建产物
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

# 安装依赖
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make build                    - Build the application"
	@echo "  make run                      - Run the application with default config"
	@echo "  make run CONFIG_FILE=<path>   - Run with specific config file"
	@echo "  make test                     - Run tests"
	@echo "  make cover                    - Run tests with coverage"
	@echo "  make fmt                      - Format code"
	@echo "  make lint                     - Run linter"
	@echo "  make clean                    - Clean build artifacts"
	@echo "  make deps                     - Install dependencies"
	@echo "  make api-test                 - Test API endpoints"
	@echo "  make api-test-help            - Show API test help"
	@echo "  make help                     - Show this help message"

# 设置默认 Go 代理
export GOPROXY=https://goproxy.cn 

# 添加测试脚本支持
api-test-full:
	@echo "Running full API test suite..."
	@chmod +x scripts/test-api.sh
	@./scripts/test-api.sh $(API_HOST) 