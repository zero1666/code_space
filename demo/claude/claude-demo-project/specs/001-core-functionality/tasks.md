# issue2md 核心功能任务列表

**分支**: `001-core-functionality` | **日期**: 2026-01-20 | **规范**: `specs/001-core-functionality/spec.md`

## 任务组织原则

- **TDD强制**: 每个实现任务前必须有对应的测试任务
- **原子化**: 每个任务只涉及一个主要文件的修改或创建
- **依赖关系**: 明确标注任务间的依赖关系
- **并行标记**: `[P]` 表示可并行执行的任务

---

## Phase 1: Foundation (数据结构定义)

### 1.1 核心数据结构定义

**任务 1.1.1** [P] - 创建 `internal/github/types.go` 测试文件
- **文件**: `internal/github/types_test.go`
- **内容**: 定义 IssueData、User、Comment、Reactions 结构体的表格驱动测试
- **依赖**: 无

**任务 1.1.2** - 实现 `internal/github/types.go`
- **文件**: `internal/github/types.go`
- **内容**: 定义 IssueData、IssueType、User、Comment、Reactions 结构体
- **依赖**: 任务 1.1.1

### 1.2 配置结构定义

**任务 1.2.1** [P] - 创建 `internal/config/config_test.go`
- **文件**: `internal/config/config_test.go`
- **内容**: 定义 Config 结构体的表格驱动测试
- **依赖**: 无

**任务 1.2.2** - 实现 `internal/config/config.go`
- **文件**: `internal/config/config.go`
- **内容**: 定义 Config 结构体和相关验证函数
- **依赖**: 任务 1.2.1

---

## Phase 2: GitHub Fetcher (API交互逻辑，TDD)

### 2.1 URL 解析器

**任务 2.1.1** [P] - 创建 `internal/parser/parser_test.go`
- **文件**: `internal/parser/parser_test.go`
- **内容**: 表格驱动测试，覆盖有效/无效 URL 解析场景
- **测试用例**:
  - 有效的 Issue/PR/Discussion URL
  - 无效的 URL 格式
  - 不支持的 GitHub URL 类型
- **依赖**: 任务 1.1.2 (IssueType 定义)

**任务 2.1.2** - 实现 `internal/parser/parser.go`
- **文件**: `internal/parser/parser.go`
- **内容**: 实现 Parser 接口和 ParseURL 函数
- **依赖**: 任务 2.1.1

### 2.2 GitHub API 客户端接口定义

**任务 2.2.1** [P] - 创建 `internal/github/client_test.go` (接口测试)
- **文件**: `internal/github/client_test.go`
- **内容**: 定义 GitHubClient 接口的测试用例
- **依赖**: 任务 1.1.2

**任务 2.2.2** - 定义 `internal/github/client.go` 接口
- **文件**: `internal/github/client.go`
- **内容**: 定义 GitHubClient 接口和工厂函数签名
- **依赖**: 任务 2.2.1

### 2.3 GitHub API 客户端实现

**任务 2.3.1** [P] - 创建 `internal/github/impl_test.go`
- **文件**: `internal/github/impl_test.go`
- **内容**: 实现具体的 API 调用测试（可配置跳过真实 API 调用）
- **依赖**: 任务 2.2.2

**任务 2.3.2** - 实现 `internal/github/impl.go`
- **文件**: `internal/github/impl.go`
- **内容**: 实现 GitHubClient 接口的具体逻辑
- **功能**:
  - FetchIssue: 获取 Issue 数据
  - FetchPullRequest: 获取 PR 数据
  - FetchDiscussion: 获取 Discussion 数据
  - 错误处理和超时控制
- **依赖**: 任务 2.3.1

**任务 2.3.3** - 更新 `internal/github/client.go` 工厂函数
- **文件**: `internal/github/client.go`
- **内容**: 实现 NewGitHubClient 和 NewGitHubClientWithHTTPClient 函数
- **依赖**: 任务 2.3.2

---

## Phase 3: Markdown Converter (转换逻辑，TDD)

### 3.1 转换器接口定义

**任务 3.1.1** [P] - 创建 `internal/converter/converter_test.go` (接口测试)
- **文件**: `internal/converter/converter_test.go`
- **内容**: 定义 Converter 接口的表格驱动测试
- **测试用例**:
  - Issue 转换测试
  - PR 转换测试
  - Discussion 转换测试
  - 可选功能测试（reactions、用户链接）
- **依赖**: 任务 1.1.2

**任务 3.1.2** - 定义 `internal/converter/converter.go` 接口
- **文件**: `internal/converter/converter.go`
- **内容**: 定义 Converter 接口和 ConvertOptions 结构体
- **依赖**: 任务 3.1.1

### 3.2 模板处理

**任务 3.2.1** [P] - 创建 `internal/converter/template_test.go`
- **文件**: `internal/converter/template_test.go`
- **内容**: 测试模板生成函数和工具函数
- **测试用例**:
  - YAML Frontmatter 生成
  - 时间戳格式化
  - Markdown 转义
- **依赖**: 任务 3.1.2

**任务 3.2.2** - 实现 `internal/converter/template.go`
- **文件**: `internal/converter/template.go`
- **内容**: 实现模板数据结构和模板生成函数
- **功能**:
  - TemplateData 结构体
  - FormatTimestamp 函数
  - EscapeMarkdown 函数
- **依赖**: 任务 3.2.1

### 3.3 转换器实现

**任务 3.3.1** [P] - 创建 `internal/converter/impl_test.go`
- **文件**: `internal/converter/impl_test.go`
- **内容**: 实现具体的转换逻辑测试
- **依赖**: 任务 3.2.2

**任务 3.3.2** - 实现 `internal/converter/impl.go`
- **文件**: `internal/converter/impl.go`
- **内容**: 实现 Converter 接口的具体逻辑
- **功能**:
  - Convert 函数实现
  - YAML Frontmatter 生成
  - 评论排序和时间格式化
  - 可选功能实现（reactions、用户链接）
- **依赖**: 任务 3.3.1

**任务 3.3.3** - 更新 `internal/converter/converter.go` 工厂函数
- **文件**: `internal/converter/converter.go`
- **内容**: 实现 NewConverter 工厂函数
- **依赖**: 任务 3.3.2

---

## Phase 4: CLI Assembly (命令行入口集成)

### 4.1 CLI 接口定义

**任务 4.1.1** [P] - 创建 `internal/cli/cli_test.go` (接口测试)
- **文件**: `internal/cli/cli_test.go`
- **内容**: 定义 CLI 接口和参数解析的表格驱动测试
- **测试用例**:
  - 参数解析测试
  - 配置验证测试
  - 错误处理测试
- **依赖**: 任务 1.2.2

**任务 4.1.2** - 定义 `internal/cli/cli.go` 接口
- **文件**: `internal/cli/cli.go`
- **内容**: 定义 CLI 接口和 Config 结构体
- **依赖**: 任务 4.1.1

### 4.2 CLI 实现

**任务 4.2.1** [P] - 创建 `internal/cli/impl_test.go`
- **文件**: `internal/cli/impl_test.go`
- **内容**: 实现 CLI 运行逻辑的测试
- **依赖**: 任务 4.1.2, 任务 2.1.2, 任务 2.3.3, 任务 3.3.3

**任务 4.2.2** - 实现 `internal/cli/impl.go`
- **文件**: `internal/cli/impl.go`
- **内容**: 实现 CLI 接口的具体逻辑
- **功能**:
  - ParseFlags 函数实现
  - Run 函数实现（集成解析器、GitHub客户端、转换器）
  - 错误处理和退出码设置
- **依赖**: 任务 4.2.1

**任务 4.2.3** - 更新 `internal/cli/cli.go` 工厂函数
- **文件**: `internal/cli/cli.go`
- **内容**: 实现 NewCLI 工厂函数
- **依赖**: 任务 4.2.2

### 4.3 主程序入口

**任务 4.3.1** [P] - 创建 `cmd/issue2md/main_test.go`
- **文件**: `cmd/issue2md/main_test.go`
- **内容**: 测试主程序的错误处理和退出码
- **依赖**: 任务 4.2.3

**任务 4.3.2** - 实现 `cmd/issue2md/main.go`
- **文件**: `cmd/issue2md/main.go`
- **内容**: 实现主程序入口点
- **功能**:
  - 环境变量读取（GITHUB_TOKEN）
  - CLI 实例创建和运行
  - 信号处理和优雅退出
- **依赖**: 任务 4.3.1

### 4.4 集成测试和构建

**任务 4.4.1** [P] - 创建端到端测试文件
- **文件**: `internal/integration/integration_test.go`
- **内容**: 完整的端到端集成测试
- **依赖**: 任务 4.3.2

**任务 4.4.2** - 更新 Makefile 构建脚本
- **文件**: `Makefile`
- **内容**: 添加构建、测试、运行命令
- **依赖**: 任务 4.4.1

**任务 4.4.3** - 创建示例和文档
- **文件**: `examples/` 目录和 README.md
- **内容**: 使用示例和项目文档
- **依赖**: 任务 4.4.2

---

## 任务依赖关系图

```
Phase 1 (Foundation)
├── 1.1.1 → 1.1.2 (types)
� 1.2.1 → 1.2.2 (config)

Phase 2 (GitHub Fetcher)
├── 2.1.1 → 2.1.2 (parser)
├── 2.2.1 → 2.2.2 (client interface)
� 2.3.1 → 2.3.2 → 2.3.3 (client impl)

Phase 3 (Markdown Converter)
├── 3.1.1 → 3.1.2 (converter interface)
├── 3.2.1 → 3.2.2 (template)
�3.3.1 → 3.3.2 → 3.3.3 (converter impl)

Phase 4 (CLI Assembly)
├── 4.1.1 → 4.1.2 (cli interface)
├── 4.2.1 → 4.2.2 → 4.2.3 (cli impl)
├── 4.3.1 → 4.3.2 (main)
└── 4.4.1 → 4.4.2 → 4.4.3 (integration)

跨阶段依赖:
2.1.1 → 1.1.2
2.2.1 → 1.1.2
3.1.1 → 1.1.2
4.1.1 → 1.2.2
4.2.1 → 2.1.2 + 2.3.3 + 3.3.3
```

## 并行执行指南

### 可并行执行的组别

**组 A** [完全并行]:
- 任务 1.1.1, 1.2.1 (Foundation 测试)
- 任务 2.1.1, 2.2.1 (GitHub Fetcher 接口测试)
- 任务 3.1.1, 3.2.1 (Converter 接口和模板测试)
- 任务 4.1.1 (CLI 接口测试)

**组 B** [Foundation 完成后并行]:
- 任务 2.3.1, 3.3.1 (实现测试)
- 任务 4.2.1, 4.3.1, 4.4.1 (集成测试)

### 开发流程建议

1. **首先执行所有标记为 [P] 的测试任务**
2. **按阶段顺序实现功能**，确保每个实现任务前测试已存在
3. **定期运行集成测试**，确保各模块协同工作
4. **完成每个阶段后运行完整测试套件**

---

*本任务列表严格遵循 TDD 原则和宪法要求，确保代码质量和可维护性。*