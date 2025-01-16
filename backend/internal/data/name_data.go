package data

import (
	"encoding/json"
	"os"
)

type NameData struct {
	Characters map[string]CharacterInfo `json:"characters"`
	Phonemes   map[string][]string      `json:"phonemes"`
}

type CharacterInfo struct {
	Pinyin  string   `json:"pinyin"`
	Meaning string   `json:"meaning"`
	Usage   []string `json:"usage"`
	Score   float64  `json:"score"`
}

func NewNameData() *NameData {
	data := &NameData{}
	data.loadData()
	return data
}

func (d *NameData) loadData() error {
	// 加载字库数据
	chars, err := os.ReadFile("configs/characters.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(chars, &d.Characters); err != nil {
		return err
	}

	// 加载音素映射数据
	phonemes, err := os.ReadFile("configs/phonemes.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(phonemes, &d.Phonemes)
}
