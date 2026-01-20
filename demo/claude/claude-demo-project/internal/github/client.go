package github

import (
	"context"
)

// GitHubClient 定义了与 GitHub API 交互的接口
type GitHubClient interface {
	// FetchIssue 获取 Issue 详细信息
	FetchIssue(ctx context.Context, owner, repo string, number int) (*IssueData, error)

	// FetchPullRequest 获取 Pull Request 详细信息
	FetchPullRequest(ctx context.Context, owner, repo string, number int) (*IssueData, error)

	// FetchDiscussion 获取 Discussion 详细信息
	FetchDiscussion(ctx context.Context, owner, repo string, number int) (*IssueData, error)
}

// NewGitHubClient 创建新的 GitHub 客户端
func NewGitHubClient(token string) GitHubClient {
	return NewGitHubClientWithHTTPClient(token, nil)
}