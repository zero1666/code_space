# issue2md 核心功能规格说明书

## 概述

issue2md 是一个命令行工具，用于将 GitHub Issue、Pull Request 和 Discussion 转换为 Markdown 格式，便于文档归档和离线阅读。

**版本**: 1.0.0
**状态**: 草案
**创建日期**: 2026-01-19

---

## 用户故事

### CLI 版本 (MVP)

**作为** 开发者
**我希望** 通过命令行将 GitHub Issue/PR/Discussion 转换为 Markdown
**以便于** 离线阅读、文档归档和知识管理

**验收标准**:
- 支持 GitHub Issue、Pull Request、Discussion 三种类型
- 自动识别 URL 类型，无需手动指定
- 支持通过环境变量进行 GitHub API 认证
- 输出格式为标准 Markdown，包含 YAML Frontmatter
- 支持输出到 stdout 或指定文件

### Web 版本 (未来扩展)

**作为** 非技术用户
**我希望** 通过 Web 界面将 GitHub Issue/PR/Discussion 转换为 Markdown
**以便于** 无需安装命令行工具即可使用

---

## 功能性需求

### 1. 输入处理

#### 1.1 URL 识别与验证
- **支持格式**:
  - Issue: `https://github.com/{owner}/{repo}/issues/{number}`
  - PR: `https://github.com/{owner}/{repo}/pull/{number}`
  - Discussion: `https://github.com/{owner}/{repo}/discussions/{number}`
- **自动识别**: 工具必须自动解析 URL 结构判断类型
- **验证**: 检查 URL 格式正确性，资源不存在时返回清晰错误信息

#### 1.2 认证机制
- **认证方式**: 仅通过环境变量 `GITHUB_TOKEN` 获取 Personal Access Token
- **权限范围**: 仅支持公开仓库，无需特殊权限
- **错误处理**: 认证失败时提供清晰的错误提示

### 2. 数据获取与处理

#### 2.1 内容范围
- **必须包含**:
  - 标题、作者、创建时间、状态（Open/Closed/Merged）
  - 主楼/描述内容
  - 所有评论内容（按时间正序排列）
- **PR 特殊处理**: 仅包含描述和 Review Comments，不包含 diff 信息
- **Discussion 特殊处理**: 标记被采纳的答案（Answer）

#### 2.2 可选功能（通过 Flags 控制）
- **Reactions 统计**: 在主楼和评论下方显示 reactions 统计
- **用户链接**: 将用户名渲染为指向 GitHub 主页的链接

### 3. 输出格式

#### 3.1 Markdown 结构
```markdown
---
title: "Issue/PR/Discussion 标题"
url: "原始 GitHub URL"
author: "创建者用户名"
created_at: "2024-01-19T14:30:00Z"
updated_at: "2024-01-20T10:15:00Z"
state: "open|closed|merged"
type: "issue|pr|discussion"
---

# [标题]

**作者**: [作者] • **创建时间**: [时间] • **状态**: [状态]

[主楼/描述内容]

---

## 评论

### [评论者1] • [时间]

[评论内容]

---

### [评论者2] • [时间]

[评论内容]

[如果是 Answer，添加标记: 💡 **Accepted Answer**]
```

#### 3.2 特殊内容处理
- **图片/附件**: 保留原始 GitHub 链接，不下载到本地
- **代码块**: 保持原有的语法高亮标记
- **用户提及**: `@username` 保持原样或根据 Flag 转换为链接
- **表格/任务列表**: 保持 GitHub Flavored Markdown 格式

### 4. 命令行接口

#### 4.1 基本语法
```bash
issue2md [flags] <url> [output_file]
```

#### 4.2 参数说明
- **`<url>`** (必需): GitHub Issue/PR/Discussion URL
- **`[output_file]`** (可选): 输出文件路径。如未指定，输出到 stdout

#### 4.3 Flags
- **`-enable-reactions`**: 包含 reactions 统计信息
- **`-enable-user-links`**: 将用户名渲染为 GitHub 主页链接

#### 4.4 示例用法
```bash
# 输出到 stdout
issue2md https://github.com/owner/repo/issues/123

# 输出到文件
issue2md https://github.com/owner/repo/pull/456 output.md

# 启用所有可选功能
issue2md -enable-reactions -enable-user-links https://github.com/owner/repo/discussions/789 discussion.md
```

---

## 非功能性需求

### 1. 架构原则
- **简单性**: 遵循 Go 语言哲学，避免过度工程
- **标准库优先**: 优先使用 Go 标准库（如 `net/http`）
- **解耦设计**: API 客户端、解析器、格式化器分离

### 2. 错误处理
- **清晰错误信息**: 所有错误信息必须清晰易懂
- **适当退出码**: 使用标准 Unix 退出码（0=成功，非0=错误）
- **网络错误**: 设置合理超时（建议 30 秒）
- **API 限流**: 透传 GitHub API 错误信息

### 3. 性能要求
- **响应时间**: 单个转换应在 10 秒内完成
- **内存使用**: 处理大型讨论时内存使用应合理
- **无缓存**: MVP 阶段不实现缓存机制

---

## 验收标准

### 测试用例

#### 1. URL 验证测试
- [ ] 有效的 GitHub URL 被正确识别
- [ ] 无效的 URL 格式返回错误
- [ ] 不存在的资源返回 404 错误

#### 2. 类型识别测试
- [ ] Issue URL 正确识别为 issue 类型
- [ ] PR URL 正确识别为 pr 类型
- [ ] Discussion URL 正确识别为 discussion 类型

#### 3. 内容转换测试
- [ ] 标题、作者、时间戳正确转换
- [ ] 主楼内容完整保留
- [ ] 评论按时间正序排列
- [ ] PR Review Comments 正确包含
- [ ] Discussion Answer 正确标记

#### 4. Flags 功能测试
- [ ] `-enable-reactions` 正确显示 reactions 统计
- [ ] `-enable-user-links` 正确生成用户链接
- [ ] 默认情况下不显示可选内容

#### 5. 输出测试
- [ ] 输出到 stdout 功能正常
- [ ] 输出到文件功能正常
- [ ] YAML Frontmatter 格式正确
- [ ] Markdown 语法正确

#### 6. 错误处理测试
- [ ] 网络错误正确处理
- [ ] API 认证错误正确处理
- [ ] 资源不存在错误正确处理

---

## 输出格式示例

### Issue 转换示例

```markdown
---
title: "Add dark mode support"
url: "https://github.com/owner/repo/issues/123"
author: "alice"
created_at: "2024-01-19T14:30:00Z"
updated_at: "2024-01-20T10:15:00Z"
state: "open"
type: "issue"
---

# Add dark mode support

**作者**: alice • **创建时间**: 2024-01-19 14:30:00 UTC • **状态**: Open

We should add dark mode support to improve user experience during nighttime usage.

## Requirements
- [ ] Toggle switch in settings
- [ ] CSS variables for themes
- [ ] Persist user preference

---

## 评论

### bob • 2024-01-19 15:20:00 UTC

Great idea! I can help with the CSS implementation.

### charlie • 2024-01-20 09:45:00 UTC

We should also consider accessibility - ensure proper contrast ratios.

💡 **Accepted Answer**
```

### PR 转换示例

```markdown
---
title: "feat: implement dark mode toggle"
url: "https://github.com/owner/repo/pull/456"
author: "alice"
created_at: "2024-01-21T11:00:00Z"
updated_at: "2024-01-22T16:30:00Z"
state: "merged"
type: "pr"
---

# feat: implement dark mode toggle

**作者**: alice • **创建时间**: 2024-01-21 11:00:00 UTC • **状态**: Merged

Implements dark mode toggle with CSS variables and local storage persistence.

---

## 评论

### bob • 2024-01-21 14:30:00 UTC

Looks good! Just one small suggestion about the color scheme.

### alice • 2024-01-22 10:15:00 UTC

Fixed the color scheme as suggested.
```

---

## 技术约束

- **Go 版本**: ≥ 1.24
- **依赖管理**: 使用 Go Modules
- **测试要求**: 表格驱动测试优先
- **错误处理**: 所有错误必须显式处理，使用 `fmt.Errorf("...: %w", err)` 包装

## 未来扩展

- Web 界面支持
- 批量处理多个 URL
- 支持其他 Git 平台（GitLab、Bitbucket）
- 自定义模板支持
- 缓存机制优化