package github

import "time"

// IssueType 表示 GitHub 资源类型
type IssueType string

const (
	IssueTypeIssue      IssueType = "issue"
	IssueTypePR         IssueType = "pr"
	IssueTypeDiscussion IssueType = "discussion"
)

// IssueData 统一的数据结构，在模块间流转
type IssueData struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Author    User      `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	State     string    `json:"state"`
	Type      IssueType `json:"type"`
	Body      string    `json:"body"`
	Comments  []Comment `json:"comments"`
}

// User 表示 GitHub 用户
type User struct {
	Login string `json:"login"`
	URL   string `json:"html_url"`
}

// Comment 表示评论
type Comment struct {
	Author    User      `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
	Reactions *Reactions `json:"reactions,omitempty"`
	IsAnswer  bool      `json:"is_answer,omitempty"`
}

// Reactions 表示 reactions 统计
type Reactions struct {
	TotalCount int `json:"total_count"`
	// 支持具体 reaction 类型统计
	PlusOne    int `json:"+1"`
	MinusOne   int `json:"-1"`
	Laugh      int `json:"laugh"`
	Hooray     int `json:"hooray"`
	Confused   int `json:"confused"`
	Heart      int `json:"heart"`
	Rocket     int `json:"rocket"`
	Eyes       int `json:"eyes"`
}