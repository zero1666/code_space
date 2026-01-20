package converter

import (
	"fmt"
	"strings"
	"time"

	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/github"
)

// Converter 定义了将 GitHub 数据转换为 Markdown 的接口
type Converter interface {
	// Convert 将 GitHub 数据转换为 Markdown
	Convert(data *github.IssueData, opts ConvertOptions) (string, error)
}

// ConvertOptions 转换选项
type ConvertOptions struct {
	IncludeReactions bool
	IncludeUserLinks bool
}

// markdownConverter 实现 Converter 接口
type markdownConverter struct{}

// NewConverter 创建新的转换器
func NewConverter() Converter {
	return &markdownConverter{}
}

// Convert 将 GitHub 数据转换为 Markdown
func (c *markdownConverter) Convert(data *github.IssueData, opts ConvertOptions) (string, error) {
	if data == nil {
		return "", fmt.Errorf("data cannot be nil")
	}

	var sb strings.Builder

	// 生成 YAML Frontmatter
	sb.WriteString(generateFrontmatter(data))

	// 生成标题
	sb.WriteString(fmt.Sprintf("# %s\n\n", data.Title))

	// 生成元信息行
	sb.WriteString(generateMetaLine(data, opts))

	// 生成正文
	if data.Body != "" {
		sb.WriteString(data.Body)
		sb.WriteString("\n")
	}

	// 生成评论区
	if len(data.Comments) > 0 {
		sb.WriteString("\n---\n\n")
		sb.WriteString("## 评论\n\n")

		for _, comment := range data.Comments {
			sb.WriteString(generateComment(comment, opts))
		}
	}

	return sb.String(), nil
}

// generateFrontmatter 生成 YAML Frontmatter
func generateFrontmatter(data *github.IssueData) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	sb.WriteString(fmt.Sprintf("title: %q\n", data.Title))
	sb.WriteString(fmt.Sprintf("url: %q\n", data.URL))
	sb.WriteString(fmt.Sprintf("author: %q\n", data.Author.Login))
	sb.WriteString(fmt.Sprintf("created_at: %q\n", data.CreatedAt.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("updated_at: %q\n", data.UpdatedAt.Format(time.RFC3339)))
	sb.WriteString(fmt.Sprintf("state: %q\n", data.State))
	sb.WriteString(fmt.Sprintf("type: %q\n", data.Type))
	sb.WriteString("---\n\n")

	return sb.String()
}

// generateMetaLine 生成元信息行
func generateMetaLine(data *github.IssueData, opts ConvertOptions) string {
	author := data.Author.Login
	if opts.IncludeUserLinks && data.Author.URL != "" {
		author = fmt.Sprintf("[%s](%s)", data.Author.Login, data.Author.URL)
	}

	state := capitalizeFirst(data.State)

	return fmt.Sprintf("**作者**: %s • **创建时间**: %s • **状态**: %s\n\n",
		author,
		FormatTimestamp(data.CreatedAt),
		state,
	)
}

// generateComment 生成单条评论
func generateComment(comment github.Comment, opts ConvertOptions) string {
	var sb strings.Builder

	// 评论者和时间
	author := comment.Author.Login
	if opts.IncludeUserLinks && comment.Author.URL != "" {
		author = fmt.Sprintf("[%s](%s)", comment.Author.Login, comment.Author.URL)
	}

	sb.WriteString(fmt.Sprintf("### %s • %s\n\n", author, FormatTimestamp(comment.CreatedAt)))

	// 评论内容
	sb.WriteString(comment.Body)
	sb.WriteString("\n")

	// Reactions（如果启用）
	if opts.IncludeReactions && comment.Reactions != nil && comment.Reactions.TotalCount > 0 {
		sb.WriteString("\n")
		sb.WriteString(generateReactions(comment.Reactions))
	}

	// Answer 标记
	if comment.IsAnswer {
		sb.WriteString("\n💡 **Accepted Answer**\n")
	}

	sb.WriteString("\n---\n\n")

	return sb.String()
}

// generateReactions 生成 reactions 统计
func generateReactions(r *github.Reactions) string {
	var parts []string

	if r.PlusOne > 0 {
		parts = append(parts, fmt.Sprintf("+1: %d", r.PlusOne))
	}
	if r.MinusOne > 0 {
		parts = append(parts, fmt.Sprintf("-1: %d", r.MinusOne))
	}
	if r.Laugh > 0 {
		parts = append(parts, fmt.Sprintf("laugh: %d", r.Laugh))
	}
	if r.Hooray > 0 {
		parts = append(parts, fmt.Sprintf("hooray: %d", r.Hooray))
	}
	if r.Confused > 0 {
		parts = append(parts, fmt.Sprintf("confused: %d", r.Confused))
	}
	if r.Heart > 0 {
		parts = append(parts, fmt.Sprintf("heart: %d", r.Heart))
	}
	if r.Rocket > 0 {
		parts = append(parts, fmt.Sprintf("rocket: %d", r.Rocket))
	}
	if r.Eyes > 0 {
		parts = append(parts, fmt.Sprintf("eyes: %d", r.Eyes))
	}

	if len(parts) == 0 {
		return ""
	}

	return fmt.Sprintf("**Reactions**: %s\n", strings.Join(parts, " | "))
}

// FormatTimestamp 格式化时间戳为可读格式
func FormatTimestamp(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05 UTC")
}

// capitalizeFirst 首字母大写
func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}