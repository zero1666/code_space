# issue2md API 接口设计草图

## 概述

本文档定义了 `internal/github` 和 `internal/converter` 包的主要接口，作为后续开发的参考。

---

## internal/github 包

### 核心接口

#### GitHubClient 接口
```go
// GitHubClient 定义了与 GitHub API 交互的接口
type GitHubClient interface {
    // FetchIssue 获取 Issue 详细信息
    FetchIssue(ctx context.Context, owner, repo string, number int) (*Issue, error)

    // FetchPullRequest 获取 Pull Request 详细信息
    FetchPullRequest(ctx context.Context, owner, repo string, number int) (*PullRequest, error)

    // FetchDiscussion 获取 Discussion 详细信息
    FetchDiscussion(ctx context.Context, owner, repo string, number int) (*Discussion, error)

    // FetchComments 获取评论列表（支持分页）
    FetchComments(ctx context.Context, owner, repo string, number int, issueType IssueType) ([]Comment, error)
}
```

#### 数据结构
```go
// IssueType 表示 GitHub 资源类型
type IssueType string

const (
    IssueTypeIssue      IssueType = "issue"
    IssueTypePR         IssueType = "pr"
    IssueTypeDiscussion IssueType = "discussion"
)

// Issue 表示 GitHub Issue
type Issue struct {
    Number    int       `json:"number"`
    Title     string    `json:"title"`
    Body      string    `json:"body"`
    Author    User      `json:"user"`
    State     string    `json:"state"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    URL       string    `json:"html_url"`
}

// PullRequest 表示 GitHub Pull Request
type PullRequest struct {
    Number    int       `json:"number"`
    Title     string    `json:"title"`
    Body      string    `json:"body"`
    Author    User      `json:"user"`
    State     string    `json:"state"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    URL       string    `json:"html_url"`
    Merged    bool      `json:"merged"`
}

// Discussion 表示 GitHub Discussion
type Discussion struct {
    Number    int       `json:"number"`
    Title     string    `json:"title"`
    Body      string    `json:"body"`
    Author    User      `json:"user"`
    State     string    `json:"state"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    URL       string    `json:"html_url"`
    Answer    *Comment  `json:"answer,omitempty"`
}

// Comment 表示评论
type Comment struct {
    ID        int          `json:"id"`
    Body      string       `json:"body"`
    Author    User         `json:"user"`
    CreatedAt time.Time    `json:"created_at"`
    UpdatedAt time.Time    `json:"updated_at"`
    Reactions *Reactions   `json:"reactions,omitempty"`
    IsAnswer  bool         `json:"is_answer,omitempty"`
}

// User 表示 GitHub 用户
type User struct {
    Login string `json:"login"`
    URL   string `json:"html_url"`
}

// Reactions 表示 reactions 统计
type Reactions struct {
    TotalCount int `json:"total_count"`
    // 可根据需要添加具体 reaction 类型统计
}
```

#### 工厂函数
```go
// NewGitHubClient 创建新的 GitHub 客户端
func NewGitHubClient(token string) GitHubClient

// NewGitHubClientWithHTTPClient 使用自定义 HTTP 客户端创建 GitHub 客户端
func NewGitHubClientWithHTTPClient(token string, httpClient *http.Client) GitHubClient
```

---

## internal/converter 包

### 核心接口

#### Converter 接口
```go
// Converter 定义了将 GitHub 数据转换为 Markdown 的接口
type Converter interface {
    // ConvertIssue 将 Issue 转换为 Markdown
    ConvertIssue(issue *github.Issue, comments []github.Comment, opts ConvertOptions) (string, error)

    // ConvertPullRequest 将 Pull Request 转换为 Markdown
    ConvertPullRequest(pr *github.PullRequest, comments []github.Comment, opts ConvertOptions) (string, error)

    // ConvertDiscussion 将 Discussion 转换为 Markdown
    ConvertDiscussion(discussion *github.Discussion, comments []github.Comment, opts ConvertOptions) (string, error)
}
```

#### 配置选项
```go
// ConvertOptions 转换选项
type ConvertOptions struct {
    IncludeReactions bool
    IncludeUserLinks bool
}
```

#### 模板结构
```go
// TemplateData 模板数据
type TemplateData struct {
    Title     string
    URL       string
    Author    github.User
    CreatedAt time.Time
    UpdatedAt time.Time
    State     string
    Type      github.IssueType
    Body      string
    Comments  []CommentData
}

// CommentData 评论数据
type CommentData struct {
    Author    github.User
    CreatedAt time.Time
    Body      string
    Reactions *github.Reactions
    IsAnswer  bool
}
```

#### 工厂函数和工具函数
```go
// NewConverter 创建新的转换器
func NewConverter() Converter

// FormatTimestamp 格式化时间戳为可读格式
func FormatTimestamp(t time.Time) string

// EscapeMarkdown 转义 Markdown 特殊字符
func EscapeMarkdown(text string) string
```

---

## internal/parser 包

### 核心接口

```go
// Parser 解析 GitHub URL
type Parser interface {
    // ParseURL 解析 GitHub URL，返回资源信息
    ParseURL(url string) (*ResourceInfo, error)
}

// ResourceInfo 资源信息
type ResourceInfo struct {
    Owner   string
    Repo    string
    Number  int
    Type    github.IssueType
}

// NewParser 创建新的解析器
func NewParser() Parser
```

---

## internal/cli 包

### 核心接口

```go
// CLI 命令行接口
type CLI interface {
    // Run 执行命令行工具
    Run(ctx context.Context, args []string) error
}

// Config 配置
type Config struct {
    URL            string
    OutputFile     string
    IncludeReactions bool
    IncludeUserLinks bool
    GitHubToken    string
}

// NewCLI 创建新的 CLI 实例
func NewCLI() CLI

// ParseFlags 解析命令行参数
func ParseFlags(args []string) (*Config, error)
```

---

## 包依赖关系

```
cmd/issue2md → internal/cli → internal/parser → internal/github
                                      ↓
                              internal/converter
```

## 设计原则

1. **接口隔离**: 每个包只暴露必要的接口，隐藏实现细节
2. **依赖注入**: 通过接口依赖，便于测试和替换
3. **错误处理**: 所有错误都使用 `fmt.Errorf("...: %w", err)` 包装
4. **单一职责**: 每个包和函数只负责一个明确的职责

## 测试策略

- **表格驱动测试**: 所有包都使用表格驱动测试
- **集成测试**: `internal/github` 包需要集成测试（可配置为跳过）
- **Mock 测试**: 使用接口进行单元测试，避免真实 API 调用