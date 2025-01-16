# 中文名字生成服务

这是一个基于 Go 和 Gemini API 的中文名字生成服务，可以根据英文名生成具有文化内涵的中文名字建议。

## 功能特点

- 支持英文名到中文名的智能转换
- 提供每个字的详细解释和文化内涵
- 包含拼音、字义和个性特征分析
- RESTful API 接口
- 支持跨域请求

## 系统要求

- Go 1.21 或更高版本
- Gemini API 密钥

## 安装和运行

1. 克隆代码库：

   ```bash
   git clone https://github.com/your-username/give-names.git
   cd give-names/backend
   ```

2. 安装依赖：

   ```bash
   go mod download
   ```

3. 配置环境变量：

   ```bash
   export APP_GEMINI_API_KEY=your-api-key
   ```

4. 运行服务：
   ```bash
   go run cmd/main.go
   ```

服务将在 http://localhost:8080 启动。

## API 文档

### 生成名字建议

**请求**

```http
POST /api/generate
Content-Type: application/json

{
    "english_name": "Michael",
    "language": "en"  // 可选，默认为 "en"
}
```

**响应**

```json
{
  "suggestions": [
    {
      "chinese_name": "米凯乐",
      "pinyin": "mi kai le",
      "meaning": "凯旋欢乐",
      "characters": [
        {
          "character": "米",
          "meaning": "米，谷物",
          "pinyin": "mi"
        },
        {
          "character": "凯",
          "meaning": "凯旋，胜利",
          "pinyin": "kai"
        },
        {
          "character": "乐",
          "meaning": "快乐，音乐",
          "pinyin": "le"
        }
      ],
      "cultural_notes": "体现中国传统文化中对欢乐与成功的追求",
      "personality": "开朗、乐观、追求成功",
      "english_intro": "A name that represents joy and triumph"
    }
  ]
}
```

## 配置说明

配置文件位于 `configs/config.yaml`，支持以下配置项：

```yaml
server:
  port: 8080
  timeout: 30s

gemini:
  api_key: ${GEMINI_API_KEY}
  model: gemini-pro
  timeout: 20s

cors:
  allowed_origins: ["http://localhost:3000"]
  allowed_methods: ["GET", "POST", "OPTIONS"]
```

## 开发

### 项目结构

```
backend/
├── cmd/                    # 应用程序入口
│   └── main.go            # 主程序入口
├── internal/              # 内部包
│   ├── api/              # API 处理层
│   │   └── handler/      # HTTP 处理器
│   ├── service/          # 业务逻辑层
│   │   └── name/        # 名字生成服务
│   ├── model/           # 数据模型
│   └── config/          # 配置管理
├── pkg/                 # 可共享的包
│   ├── gemini/         # Gemini API 客户端
│   └── response/       # 统一响应处理
└── configs/            # 配置文件
    └── config.yaml
```

### 测试

运行单元测试：

```bash
go test ./...
```

### API 测试

使用 curl 测试 API：

```bash
curl -X POST http://localhost:8080/api/generate \
  -H "Content-Type: application/json" \
  -d '{"english_name": "Michael"}'
```
