package ai

import (
	"context"
	"fmt"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiService struct {
	client *genai.Client
	model  *genai.Model
}

func NewGeminiService(apiKey string) (*GeminiService, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %v", err)
	}

	model := client.GenerativeModel("gemini-pro")
	return &GeminiService{
		client: client,
		model:  model,
	}, nil
}

func (s *GeminiService) GenerateChineseNames(ctx context.Context, firstName, lastName string) (string, error) {
	prompt := fmt.Sprintf(`作为一个专业的中文姓名翻译专家，请为英文名"%s %s"生成3个富有文化内涵的中文名字。

要求：
1. 每个名字需要包含：
   - 汉字（2-3个字）
   - 拼音（带声调）
   - 字义解释
   - 文化内涵
   - 个性特征描述

2. 名字要求：
   - 发音要尽可能接近英文原名
   - 字义优美，寓意深远
   - 符合中国传统文化
   - 避免生僻字和不雅含义

请按照以下JSON格式返回结果：
{
  "names": [
    {
      "characters": [
        {
          "char": "米",
          "pinyin": "mǐ",
          "meaning": "米粒；珍贵"
        }
      ],
      "pinyin": "mǐ kǎi lè",
      "meaning": "寓意描述",
      "cultural": "文化内涵解释"
    }
  ]
}`, firstName, lastName)

	resp, err := s.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %v", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	return resp.Candidates[0].Content.Parts[0].Text(), nil
}
