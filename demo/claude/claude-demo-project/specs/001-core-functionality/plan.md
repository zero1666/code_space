# Implementation Plan: issue2md 核心功能

**Branch**: `001-core-functionality` | **Date**: 2026-01-20 | **Spec**: `specs/001-core-functionality/spec.md`
**Input**: Feature specification from `/specs/001-core-functionality/spec.md`

## Summary

issue2md 是一个命令行工具，用于将 GitHub Issue、Pull Request 和 Discussion 转换为 Markdown 格式。核心功能包括：
- 自动识别和解析 GitHub URL（Issue/PR/Discussion）
- 通过 GitHub API 获取完整内容数据
- 转换为标准 Markdown 格式，包含 YAML Frontmatter
- 支持可选功能（reactions 统计、用户链接）

技术方案基于 Go 标准库和 `google/go-github` 库，严格遵循项目宪法原则。

## Technical Context

**Language/Version**: Go 1.23.2
**Primary Dependencies**: `github.com/google/go-github/v60` (GitHub API 客户端)
**Storage**: 无数据库，通过 API 实时获取数据
**Testing**: Go 标准测试框架，表格驱动测试优先
**Target Platform**: 跨平台命令行工具 (Linux/macOS/Windows)
**Project Type**: 单项目 CLI 工具
**Performance Goals**: 单个转换在 10 秒内完成
**Constraints**: 内存使用合理，无缓存机制
**Scale/Scope**: 单个 URL 处理，支持公开仓库

## Constitution Check

*GATE: 通过宪法审查，方案符合所有宪法条款*

### 第一条：简单性原则 ✅
- **1.1 (YAGNI)**: 仅实现 spec.md 中明确要求的功能，无过度设计
- **1.2 (标准库优先)**: 使用 `net/http` 作为 HTTP 客户端，仅引入必要的 `go-github` 库
- **1.3 (反过度工程)**: 简单的函数和数据结构，避免复杂接口体系

### 第二条：测试先行铁律 ✅
- **2.1 (TDD循环)**: 所有功能从编写失败的测试开始
- **2.2 (表格驱动)**: 单元测试优先采用表格驱动测试风格
- **2.3 (拒绝Mocks)**: 优先编写集成测试，使用真实依赖

### 第三条：明确性原则 ✅
- **3.1 (错误处理)**: 所有错误显式处理，使用 `fmt.Errorf("...: %w", err)` 包装
- **3.2 (无全局变量)**: 所有依赖通过函数参数或结构体成员显式注入

## Project Structure

### Documentation (this feature)

```text
specs/001-core-functionality/
├── plan.md              # 本文件
├── spec.md              # 功能规格说明书
└── api-sketch.md        # API 接口设计草图
```

### Source Code (repository root)

```text
cmd/
├── issue2md/            # CLI 入口点
│   └── main.go
└── issue2mdweb/         # Web 版本（未来扩展）

internal/
├── cli/                 # 命令行接口处理
│   ├── cli.go          # CLI 接口和配置
│   └── cli_test.go     # CLI 测试
├── config/              # 配置管理
│   ├── config.go       # 配置结构
│   └── config_test.go  # 配置测试
├── converter/           # Markdown 转换器
│   ├── converter.go    # 转换器接口和实现
│   ├── template.go     # 模板处理
│   └── converter_test.go
├── github/              # GitHub API 客户端
│   ├── client.go       # GitHub 客户端接口和实现
│   ├── types.go        # 数据结构定义
│   └── client_test.go  # 客户端测试
└── parser/              # URL 解析器
    ├── parser.go       # URL 解析接口和实现
    └── parser_test.go  # 解析器测试

go.mod                  # Go 模块定义
go.sum                  # 依赖锁定
Makefile                # 构建和测试脚本
```

**Structure Decision**: 采用标准的 Go 项目结构，`cmd/` 包含入口点，`internal/` 包含核心业务逻辑，包之间依赖关系清晰。

## Core Data Structures

### GitHub 数据结构 (`internal/github/types.go`)

```go
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

type IssueType string

const (
    IssueTypeIssue      IssueType = "issue"
    IssueTypePR         IssueType = "pr"
    IssueTypeDiscussion IssueType = "discussion"
)

type User struct {
    Login string `json:"login"`
    URL   string `json:"html_url"`
}

type Comment struct {
    Author    User      `json:"author"`
    CreatedAt time.Time `json:"created_at"`
    Body      string    `json:"body"`
    Reactions *Reactions `json:"reactions,omitempty"`
    IsAnswer  bool      `json:"is_answer,omitempty"`
}

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
```

## Interface Design

### GitHub Client Interface (`internal/github/client.go`)

```go
type GitHubClient interface {
    FetchIssue(ctx context.Context, owner, repo string, number int) (*IssueData, error)
    FetchPullRequest(ctx context.Context, owner, repo string, number int) (*IssueData, error)
    FetchDiscussion(ctx context.Context, owner, repo string, number int) (*IssueData, error)
}

func NewGitHubClient(token string) GitHubClient
func NewGitHubClientWithHTTPClient(token string, httpClient *http.Client) GitHubClient
```

### Parser Interface (`internal/parser/parser.go`)

```go
type Parser interface {
    ParseURL(url string) (*ResourceInfo, error)
}

type ResourceInfo struct {
    Owner  string
    Repo   string
    Number int
    Type   github.IssueType
}

func NewParser() Parser
```

### Converter Interface (`internal/converter/converter.go`)

```go
type Converter interface {
    Convert(data *github.IssueData, opts ConvertOptions) (string, error)
}

type ConvertOptions struct {
    IncludeReactions bool
    IncludeUserLinks bool
}

func NewConverter() Converter
```

### CLI Interface (`internal/cli/cli.go`)

```go
type CLI interface {
    Run(ctx context.Context, args []string) error
}

type Config struct {
    URL             string
    OutputFile      string
    IncludeReactions bool
    IncludeUserLinks bool
    GitHubToken     string
}

func NewCLI(githubClient github.GitHubClient, parser parser.Parser, converter converter.Converter) CLI
func ParseFlags(args []string) (*Config, error)
```

## Package Dependencies

```
cmd/issue2md → internal/cli → internal/parser → internal/github
                                      ↓
                              internal/converter
```

## Implementation Phases

### Phase 1: 核心基础设施 (Week 1)
1. 实现 `internal/parser` - URL 解析器
2. 实现 `internal/github` - GitHub API 客户端
3. 编写表格驱动测试

### Phase 2: 转换引擎 (Week 2)
1. 实现 `internal/converter` - Markdown 转换器
2. 实现 YAML Frontmatter 生成
3. 实现可选功能（reactions、用户链接）

### Phase 3: CLI 集成 (Week 3)
1. 实现 `internal/cli` - 命令行接口
2. 实现 `cmd/issue2md` - 主程序入口
3. 集成测试和错误处理

### Phase 4: 完善和优化 (Week 4)
1. 性能优化和错误处理完善
2. 文档和示例
3. 最终测试和发布准备

## Testing Strategy

### 单元测试 (表格驱动)
- **Parser**: URL 解析测试，包含有效/无效 URL 场景
- **Converter**: Markdown 转换测试，包含各种内容类型
- **CLI**: 参数解析和配置测试

### 集成测试
- **GitHub Client**: 真实 API 调用测试（可配置跳过）
- **端到端测试**: 完整流程测试，从 URL 输入到 Markdown 输出

### 错误处理测试
- 网络错误、认证错误、资源不存在等场景
- 错误信息清晰度和退出码验证

## Risk Assessment

### 技术风险
- **GitHub API 限流**: 实现合理的重试机制和错误处理
- **网络稳定性**: 设置合理的超时时间（30秒）
- **内存使用**: 处理大型讨论时监控内存使用

### 缓解措施
- 使用 context 进行超时控制
- 实现渐进式错误处理
- 添加性能监控和日志记录

## Success Metrics

### 功能完成度
- [ ] URL 识别和验证功能
- [ ] GitHub API 集成
- [ ] Markdown 转换功能
- [ ] CLI 接口
- [ ] 可选功能实现

### 质量指标
- [ ] 100% 测试覆盖率（单元测试）
- [ ] 集成测试通过
- [ ] 错误处理完备
- [ ] 性能目标达成（10秒内完成转换）

---

*本技术方案已通过宪法审查，符合所有开发原则，可进入实施阶段。*