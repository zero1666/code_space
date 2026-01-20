package config

import (
	"errors"
	"os"
	"strings"
)

// Config 表示命令行配置
type Config struct {
	URL             string
	OutputFile      string
	IncludeReactions bool
	IncludeUserLinks bool
	GitHubToken     string
}

// Validate 验证配置的有效性
func (c *Config) Validate() error {
	if c.URL == "" {
		return errors.New("URL is required")
	}

	if !isValidGitHubURL(c.URL) {
		return errors.New("invalid GitHub URL format")
	}

	if c.GitHubToken == "" {
		return errors.New("GitHub token is required (set GITHUB_TOKEN environment variable)")
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

// LoadFromEnvironment 从环境变量加载配置
func LoadFromEnvironment() *Config {
	return &Config{
		GitHubToken: os.Getenv("GITHUB_TOKEN"),
	}
}