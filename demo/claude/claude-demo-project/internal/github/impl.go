package github

import (
	"context"
	"fmt"
	"net/http"
	"time"

	githubv60 "github.com/google/go-github/v60/github"
)

// githubClient 实现 GitHubClient 接口
type githubClient struct {
	client  *githubv60.Client
	token   string
	baseURL string // 用于测试时指向 Mock Server
}

// NewGitHubClientWithHTTPClient 使用自定义 HTTP 客户端创建 GitHub 客户端
func NewGitHubClientWithHTTPClient(token string, httpClient *http.Client) GitHubClient {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}

	client := githubv60.NewClient(httpClient)
	if token != "" {
		client = client.WithAuthToken(token)
	}

	return &githubClient{
		client:  client,
		token:   token,
		baseURL: "",
	}
}

// NewGitHubClientWithBaseURL 创建使用自定义 Base URL 的 GitHub 客户端（用于测试）
func NewGitHubClientWithBaseURL(token string, baseURL string) GitHubClient {
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	client := githubv60.NewClient(httpClient)
	if token != "" {
		client = client.WithAuthToken(token)
	}

	// 设置自定义 Base URL（用于测试）
	if baseURL != "" {
		client, _ = client.WithEnterpriseURLs(baseURL, baseURL)
	}

	return &githubClient{
		client:  client,
		token:   token,
		baseURL: baseURL,
	}
}

// validateParams 验证公共参数
func validateParams(owner, repo string, number int) error {
	if owner == "" {
		return fmt.Errorf("owner cannot be empty")
	}
	if repo == "" {
		return fmt.Errorf("repo cannot be empty")
	}
	if number <= 0 {
		return fmt.Errorf("number must be positive")
	}
	return nil
}

// FetchIssue 获取 Issue 详细信息
func (c *githubClient) FetchIssue(ctx context.Context, owner, repo string, number int) (*IssueData, error) {
	if err := validateParams(owner, repo, number); err != nil {
		return nil, err
	}

	// 获取 Issue 基本信息
	issue, resp, err := c.client.Issues.Get(ctx, owner, repo, number)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("issue not found: %s/%s#%d", owner, repo, number)
		}
		if resp != nil && resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("unauthorized: invalid or missing token")
		}
		return nil, fmt.Errorf("failed to fetch issue: %w", err)
	}

	// 获取评论
	comments, err := c.fetchIssueComments(ctx, owner, repo, number)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}

	return &IssueData{
		Title:     issue.GetTitle(),
		URL:       issue.GetHTMLURL(),
		Author:    userFromGitHub(issue.GetUser()),
		CreatedAt: issue.GetCreatedAt().Time,
		UpdatedAt: issue.GetUpdatedAt().Time,
		State:     issue.GetState(),
		Type:      IssueTypeIssue,
		Body:      issue.GetBody(),
		Comments:  comments,
	}, nil
}

// FetchPullRequest 获取 Pull Request 详细信息
func (c *githubClient) FetchPullRequest(ctx context.Context, owner, repo string, number int) (*IssueData, error) {
	if err := validateParams(owner, repo, number); err != nil {
		return nil, err
	}

	// 获取 PR 基本信息
	pr, resp, err := c.client.PullRequests.Get(ctx, owner, repo, number)
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("pull request not found: %s/%s#%d", owner, repo, number)
		}
		if resp != nil && resp.StatusCode == http.StatusUnauthorized {
			return nil, fmt.Errorf("unauthorized: invalid or missing token")
		}
		return nil, fmt.Errorf("failed to fetch pull request: %w", err)
	}

	// 确定 PR 状态
	state := pr.GetState()
	if pr.GetMerged() {
		state = "merged"
	}

	// 获取 PR 评论（包括 Review Comments）
	comments, err := c.fetchPRComments(ctx, owner, repo, number)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}

	return &IssueData{
		Title:     pr.GetTitle(),
		URL:       pr.GetHTMLURL(),
		Author:    userFromGitHub(pr.GetUser()),
		CreatedAt: pr.GetCreatedAt().Time,
		UpdatedAt: pr.GetUpdatedAt().Time,
		State:     state,
		Type:      IssueTypePR,
		Body:      pr.GetBody(),
		Comments:  comments,
	}, nil
}

// FetchDiscussion 获取 Discussion 详细信息
// 注意：GitHub REST API 不直接支持 Discussions，需要使用 GraphQL API
// 这里提供一个简化实现，实际生产环境应使用 GraphQL
func (c *githubClient) FetchDiscussion(ctx context.Context, owner, repo string, number int) (*IssueData, error) {
	if err := validateParams(owner, repo, number); err != nil {
		return nil, err
	}

	// GitHub REST API 不支持 Discussions
	// 返回一个提示信息，建议使用 GraphQL API
	return nil, fmt.Errorf("discussions require GraphQL API (not implemented in this version)")
}

// fetchIssueComments 获取 Issue 的所有评论
func (c *githubClient) fetchIssueComments(ctx context.Context, owner, repo string, number int) ([]Comment, error) {
	opts := &githubv60.IssueListCommentsOptions{
		ListOptions: githubv60.ListOptions{
			PerPage: 100,
		},
	}

	var allComments []Comment

	for {
		comments, resp, err := c.client.Issues.ListComments(ctx, owner, repo, number, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list comments: %w", err)
		}

		for _, comment := range comments {
			allComments = append(allComments, Comment{
				Author:    userFromGitHub(comment.GetUser()),
				CreatedAt: comment.GetCreatedAt().Time,
				Body:      comment.GetBody(),
				Reactions: reactionsFromGitHub(comment.GetReactions()),
			})
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allComments, nil
}

// fetchPRComments 获取 PR 的所有评论（包括普通评论和 Review Comments）
func (c *githubClient) fetchPRComments(ctx context.Context, owner, repo string, number int) ([]Comment, error) {
	var allComments []Comment

	// 获取 Issue 类型的评论
	issueComments, err := c.fetchIssueComments(ctx, owner, repo, number)
	if err != nil {
		return nil, err
	}
	allComments = append(allComments, issueComments...)

	// 获取 Review Comments
	opts := &githubv60.PullRequestListCommentsOptions{
		ListOptions: githubv60.ListOptions{
			PerPage: 100,
		},
	}

	for {
		reviewComments, resp, err := c.client.PullRequests.ListComments(ctx, owner, repo, number, opts)
		if err != nil {
			return nil, fmt.Errorf("failed to list review comments: %w", err)
		}

		for _, comment := range reviewComments {
			allComments = append(allComments, Comment{
				Author:    userFromGitHub(comment.GetUser()),
				CreatedAt: comment.GetCreatedAt().Time,
				Body:      comment.GetBody(),
				Reactions: reactionsFromGitHub(comment.GetReactions()),
			})
		}

		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allComments, nil
}

// userFromGitHub 从 GitHub User 转换为内部 User
func userFromGitHub(u *githubv60.User) User {
	if u == nil {
		return User{}
	}
	return User{
		Login: u.GetLogin(),
		URL:   u.GetHTMLURL(),
	}
}

// reactionsFromGitHub 从 GitHub Reactions 转换为内部 Reactions
func reactionsFromGitHub(r *githubv60.Reactions) *Reactions {
	if r == nil {
		return nil
	}
	return &Reactions{
		TotalCount: r.GetTotalCount(),
		PlusOne:    r.GetPlusOne(),
		MinusOne:   r.GetMinusOne(),
		Laugh:      r.GetLaugh(),
		Hooray:     r.GetHooray(),
		Confused:   r.GetConfused(),
		Heart:      r.GetHeart(),
		Rocket:     r.GetRocket(),
		Eyes:       r.GetEyes(),
	}
}