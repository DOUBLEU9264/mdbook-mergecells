package main

import (
	"encoding/json"
)

// 定义顶层JSON结构（包含两个元素的数组）
type BookData struct {
	Context Context `json:"config"`
	Book    Book    `json:"book"`
}

func NewBookDataFromJson(data []byte) (*BookData, error) {
	// 解析JSON到结构体
	var topLevel []json.RawMessage
	if err := json.Unmarshal(data, &topLevel); err != nil {
		return nil, err
	}

	var bookData BookData
	if err := json.Unmarshal(topLevel[0], &bookData.Context); err != nil {
		return nil, err
	}
	if err := json.Unmarshal(topLevel[1], &bookData.Book); err != nil {
		return nil, err
	}

	return &bookData, nil
}

func (bd *BookData) ToJsonIndent() ([]byte, error) {
	// 将Book对象编码为JSON
	return json.MarshalIndent(bd.Book, "", "  ")
}

func (bd *BookData) ProcessBook() {
	sectionsPtr := &bd.Book.Sections
	for i := range *sectionsPtr {
		chapterContentPtr := &(*sectionsPtr)[i].Chapter.Content
		// 原地修改
		*chapterContentPtr = ProcChapterContent(*chapterContentPtr)
	}
}

// 第一部分的配置结构
type Context struct {
	Root          string `json:"root"`
	Renderer      string `json:"renderer"`
	MdbookVersion string `json:"mdbook_version"`
	Config        struct {
		Book struct {
			Authors  []string `json:"authors"`
			Language string   `json:"language"`
			Src      string   `json:"src"`
			Title    string   `json:"title"`
		} `json:"book"`
		Preprocessor map[string]map[string]string `json:"preprocessor"`
	} `json:"config"`
}

// 第二部分的书籍结构
type Book struct {
	Sections      []Section `json:"sections"`
	NonExhaustive any       `json:"__non_exhaustive"`
}

// 章节结构
type Section struct {
	Chapter Chapter `json:"Chapter"`
}

// 章节详情
type Chapter struct {
	Name        string   `json:"name"`
	Content     string   `json:"content"`
	Number      []int    `json:"number"`
	SubItems    []any    `json:"sub_items"`
	Path        string   `json:"path"`
	SourcePath  string   `json:"source_path"`
	ParentNames []string `json:"parent_names"`
}
