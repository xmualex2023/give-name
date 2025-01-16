package service

import (
	"context"
	"encoding/json"
	"github.com/xmualex2023/give-name/internal/ai"
	"github.com/xmualex2023/give-name/internal/data"
)

type NameService struct {
	nameData *data.NameData
	gemini   *ai.GeminiService
}

type ChineseName struct {
	Characters []Character `json:"characters"`
	Pinyin     string      `json:"pinyin"`
	Meaning    string      `json:"meaning"`
	Cultural   string      `json:"cultural"`
}

type Character struct {
	Char    string `json:"char"`
	Pinyin  string `json:"pinyin"`
	Meaning string `json:"meaning"`
}

func NewNameService() *NameService {
	return &NameService{
		nameData: data.NewNameData(),
	}
}

func (s *NameService) Generate(firstName, lastName string) []ChineseName {
	ctx := context.Background()
	jsonStr, err := s.gemini.GenerateChineseNames(ctx, firstName, lastName)
	if err != nil {
		// 处理错误，返回空结果
		return []ChineseName{}
	}

	var result struct {
		Names []ChineseName `json:"names"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return []ChineseName{}
	}

	return result.Names
}

func (s *NameService) convertToPhonemes(name string) []string {
	// 实现英文到音素的转换逻辑
	// 这里需要使用音素库和转换规则
	return []string{}
}

func (s *NameService) generateCandidates(phonemes []string) []ChineseName {
	// 根据音素生成候选中文名字
	// 使用字库和组合规则
	return []ChineseName{}
}

func (s *NameService) rankNames(names []ChineseName) []ChineseName {
	// 对候选名字进行评分和排序
	// 考虑发音相似度、文化内涵等因素
	return names
}
