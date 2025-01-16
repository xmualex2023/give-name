package model

// NameRequest 表示名字生成请求
type NameRequest struct {
	EnglishName string `json:"english_name"`
	Language    string `json:"language"` // en/zh
}

// NameSuggestion 表示单个名字建议
type NameSuggestion struct {
	ChineseName   string `json:"chinese_name"`   // 中文名
	Pinyin        string `json:"pinyin"`         // 拼音
	Meaning       string `json:"meaning"`        // 字义
	Characters    []Char `json:"characters"`     // 单字解释
	CulturalNotes string `json:"cultural_notes"` // 文化内涵
	Personality   string `json:"personality"`    // 个性特征
	EnglishIntro  string `json:"english_intro"`  // 英文说明
}

// Char 表示单个汉字的信息
type Char struct {
	Character string `json:"character"` // 汉字
	Pinyin    string `json:"pinyin"`    // 拼音
}

// NameResponse 表示名字生成响应
type NameResponse struct {
	Suggestions []NameSuggestion `json:"suggestions"`
}
