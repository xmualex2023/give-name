# 英文转中文姓名推荐系统

## 项目简介

这是一个智能化的英文转中文姓名推荐系统，专门为外国友人提供富有文化内涵的中文名字选择。系统会根据用户输入的英文名，推荐 3 个独特的中文名字方案，并提供详细的文化解读。

## 功能特点

- 智能中文名字生成
- 详细的文化解读说明
- 简洁直观的用户界面
- 名字收藏与导出功能

## 项目架构

```
project/
├── frontend/           # 前端代码
│   ├── index.html     # 主页面
│   ├── styles.css     # 样式文件
│   ├── js/           # JavaScript 文件
│   │   └── main.js   # 主要逻辑
│   └── assets/       # 静态资源
│       └── logo.svg  # Logo
│
├── backend/           # 后端代码
│   ├── cmd/          # 主程序入口
│   │   └── main.go   # 服务启动
│   ├── internal/     # 内部包
│   │   ├── api/     # API 处理
│   │   ├── service/ # 业务逻辑
│   │   ├── ai/      # AI 服务集成
│   │   └── data/    # 数据层
│   └── configs/     # 配置文件
│       ├── characters.json  # 汉字库
│       └── phonemes.json    # 音素映射
└── README.md
```

## 技术栈

### 前端

- HTML5 + CSS3
- 原生 JavaScript
- 响应式设计
- 多语言支持 (中/英)

### 后端

- Go 1.21+
- Gin Web 框架
- Google Gemini AI
- RESTful API

## 如何运行

### 前置要求

- Go 1.21 或更高版本
- Node.js 18 或更高版本（用于开发）
- Google Gemini API Key

### 后端启动

```bash
# 1. 克隆项目
git clone https://github.com/your-username/chinese-name-generator.git
cd chinese-name-generator

# 2. 安装依赖
go mod download

# 3. 配置环境变量
export GEMINI_API_KEY=your_api_key

# 4. 启动后端服务
cd backend/cmd
go run main.go
```

### 前端启动

```bash
# 1. 进入前端目录
cd frontend

# 2. 安装 http-server（如果没有）
npm install -g http-server

# 3. 启动前端服务
http-server -p 8080
```

访问 http://localhost:8080 即可使用系统

## API 文档

### 生成名字

**请求**

```
POST /api/generate
Content-Type: application/json

{
    "firstName": "string",
    "lastName": "string"
}
```

**响应**

```json
{
  "names": [
    {
      "characters": [
        {
          "char": "string",
          "pinyin": "string",
          "meaning": "string"
        }
      ],
      "pinyin": "string",
      "meaning": "string",
      "cultural": "string"
    }
  ]
}
```

## 页面结构

### 1. 主页 (index.html)

- 顶部导航栏：logo、语言切换
- 中心搜索区：英文名输入框
- 推荐结果展示区
- 底部：版权信息

### 2. 样式设计

- 采用简约现代风格
- 主色调：中国红(#E54B4B)、青瓷蓝(#89C4F4)
- 字体：英文使用 Roboto，中文使用思源黑体

## 开发规范

### 代码风格

- 遵循 Go 官方代码规范
- 使用 ESLint 进行 JavaScript 代码检查
- 使用 Prettier 进行代码格式化

### Git 提交规范

```
feat: 新功能
fix: 修复问题
docs: 文档修改
style: 代码格式修改
refactor: 代码重构
test: 测试用例修改
chore: 其他修改
```

## 部署说明

### Docker 部署

```bash
# 构建镜像
docker build -t chinese-name-generator .

# 运行容器
docker run -p 8000:8000 -e GEMINI_API_KEY=your_api_key chinese-name-generator
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 技术架构

- 前端：HTML5 + CSS3
- 响应式设计：适配移动端和桌面端
- 字库系统：JSON 格式存储中文字库

## 联系方式

- 项目维护者：[Your Name]
- Email：contact@chinesename.com
- 项目主页：https://github.com/your-username/chinese-name-generator
