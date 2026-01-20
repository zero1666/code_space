package cli

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/converter"
	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/github"
	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/parser"
)

// CLI 命令行接口
type CLI interface {
	// Run 执行命令行工具
	Run(ctx context.Context, args []string) error
}

// Config 配置
type Config struct {
	URL              string
	OutputFile       string
	IncludeReactions bool
	IncludeUserLinks bool
	GitHubToken      string
}

// Validate 验证配置的有效性
func (c *Config) Validate() error {
	if c.URL == "" {
		return fmt.Errorf("URL is required")
	}

	if !isValidGitHubURL(c.URL) {
		return fmt.Errorf("invalid GitHub URL format")
	}

	if c.GitHubToken == "" {
		return fmt.Errorf("GitHub token is required (set GITHUB_TOKEN environment variable)")
	}

	return nil
}

// isValidGitHubURL 检查 URL 是否为有效的 GitHub Issue/PR/Discussion URL
func isValidGitHubURL(url string) bool {
	if !strings.HasPrefix(url, "https://github.com/") {
		return false
	}

	// 检查是否包含 issues、pull 或 discussions 路径
	paths := []string{"/issues/", "/pull/", "/discussions/"}
	for _, path := range paths {
		if strings.Contains(url, path) {
			return true
		}
	}

	return false
}

// cliImpl 实现 CLI 接口
type cliImpl struct {
	githubClient github.GitHubClient
	parser       parser.Parser
	converter    converter.Converter
	stdout       io.Writer
}

// NewCLI 创建新的 CLI 实例
func NewCLI(githubClient github.GitHubClient, p parser.Parser, conv converter.Converter) CLI {
	return &cliImpl{
		githubClient: githubClient,
		parser:       p,
		converter:    conv,
		stdout:       os.Stdout,
	}
}

// ParseFlags 解析命令行参数
func ParseFlags(args []string) (*Config, error) {
	flags := flag.NewFlagSet("issue2md", flag.ContinueOnError)

	var (
		includeReactions bool
		includeUserLinks bool
	)

	flags.BoolVar(&includeReactions, "enable-reactions", false, "Include reactions statistics")
	flags.BoolVar(&includeUserLinks, "enable-user-links", false, "Render usernames as GitHub profile links")

	if err := flags.Parse(args); err != nil {
		return nil, fmt.Errorf("failed to parse flags: %w", err)
	}

	// 获取位置参数
	positionalArgs := flags.Args()
	var url, outputFile string

	if len(positionalArgs) > 0 {
		url = positionalArgs[0]
	}
	if len(positionalArgs) > 1 {
		outputFile = positionalArgs[1]
	}

	return &Config{
		URL:              url,
		OutputFile:       outputFile,
		IncludeReactions: includeReactions,
		IncludeUserLinks: includeUserLinks,
		GitHubToken:      os.Getenv("GITHUB_TOKEN"),
	}, nil
}

// Run 执行命令行工具
func (c *cliImpl) Run(ctx context.Context, args []string) error {
	config, err := ParseFlags(args)
	if err != nil {
		return fmt.Errorf("configuration error: %w", err)
	}

	if err := config.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	// 解析 URL
	resourceInfo, err := c.parser.ParseURL(config.URL)
	if err != nil {
		return fmt.Errorf("failed to parse URL: %w", err)
	}

	// 获取数据
	var data *github.IssueData
	switch resourceInfo.Type {
	case github.IssueTypeIssue:
		data, err = c.githubClient.FetchIssue(ctx, resourceInfo.Owner, resourceInfo.Repo, resourceInfo.Number)
	case github.IssueTypePR:
		data, err = c.githubClient.FetchPullRequest(ctx, resourceInfo.Owner, resourceInfo.Repo, resourceInfo.Number)
	case github.IssueTypeDiscussion:
		data, err = c.githubClient.FetchDiscussion(ctx, resourceInfo.Owner, resourceInfo.Repo, resourceInfo.Number)
	default:
		return fmt.Errorf("unsupported resource type: %s", resourceInfo.Type)
	}

	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}

	// 转换为 Markdown
	opts := converter.ConvertOptions{
		IncludeReactions: config.IncludeReactions,
		IncludeUserLinks: config.IncludeUserLinks,
	}

	markdown, err := c.converter.Convert(data, opts)
	if err != nil {
		return fmt.Errorf("failed to convert to markdown: %w", err)
	}

	// 输出结果
	if config.OutputFile != "" {
		if err := os.WriteFile(config.OutputFile, []byte(markdown), 0644); err != nil {
			return fmt.Errorf("failed to write output file: %w", err)
		}
	} else {
		fmt.Fprint(c.stdout, markdown)
	}

	return nil
}