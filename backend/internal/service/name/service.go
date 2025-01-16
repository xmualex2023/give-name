package name

import (
    "context"

    "github.com/give-names/backend/internal/model"
    "github.com/give-names/backend/pkg/gemini"
)

// Service 名字生成服务接口
type Service interface {
    GenerateNames(ctx context.Context, req *model.NameRequest) (*model.NameResponse, error)
}

// service 实现 Service 接口
type service struct {
    geminiClient *gemini.Client
}

// NewService 创建新的名字生成服务
func NewService(geminiClient *gemini.Client) Service {
    return &service{
        geminiClient: geminiClient,
    }
}

// GenerateNames 生成名字建议
func (s *service) GenerateNames(ctx context.Context, req *model.NameRequest) (*model.NameResponse, error) {
    return s.geminiClient.GenerateNames(ctx, req)
} 