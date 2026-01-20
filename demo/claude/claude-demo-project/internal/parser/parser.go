package parser

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/github"
)

// Parser 解析 GitHub URL
type Parser interface {
	ParseURL(url string) (*ResourceInfo, error)
}

// ResourceInfo 资源信息
type ResourceInfo struct {
	Owner  string
	Repo   string
	Number int
	Type   github.IssueType
}

// githubParser 实现 Parser 接口
type githubParser struct{}

// NewParser 创建新的解析器
func NewParser() Parser {
	return &githubParser{}
}

// pathComponents 存储解析后的路径组件
type pathComponents struct {
	owner        string
	repo         string
	resourceType string
	numberStr    string
}

// splitAndValidatePath 分割并验证 URL 路径
// 返回路径组件或错误
func splitAndValidatePath(parsedURL *url.URL) (*pathComponents, error) {
	if parsedURL.Host != "github.com" {
		return nil, fmt.Errorf("only GitHub URLs are supported")
	}

	pathParts := strings.Split(strings.Trim(parsedURL.Path, "/"), "/")
	if len(pathParts) < 4 {
		return nil, fmt.Errorf("invalid GitHub URL format")
	}

	return &pathComponents{
		owner:        pathParts[0],
		repo:         pathParts[1],
		resourceType: pathParts[2],
		numberStr:    pathParts[3],
	}, nil
}

// parseResourceType 解析资源类型
func parseResourceType(resourceType string) (github.IssueType, error) {
	switch resourceType {
	case "issues":
		return github.IssueTypeIssue, nil
	case "pull":
		return github.IssueTypePR, nil
	case "discussions":
		return github.IssueTypeDiscussion, nil
	default:
		return "", fmt.Errorf("unsupported resource type: %s", resourceType)
	}
}

// parseResourceNumber 解析资源编号
func parseResourceNumber(numberStr string) (int, error) {
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("invalid resource number: %w", err)
	}

	if number <= 0 {
		return 0, fmt.Errorf("resource number must be positive")
	}

	return number, nil
}

// ParseURL 解析 GitHub URL，返回资源信息
func (p *githubParser) ParseURL(urlStr string) (*ResourceInfo, error) {
	if urlStr == "" {
		return nil, fmt.Errorf("URL cannot be empty")
	}

	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid URL format: %w", err)
	}

	// 分割并验证路径
	components, err := splitAndValidatePath(parsedURL)
	if err != nil {
		return nil, err
	}

	// 解析资源类型
	issueType, err := parseResourceType(components.resourceType)
	if err != nil {
		return nil, err
	}

	// 解析资源编号
	number, err := parseResourceNumber(components.numberStr)
	if err != nil {
		return nil, err
	}

	return &ResourceInfo{
		Owner:  components.owner,
		Repo:   components.repo,
		Number: number,
		Type:   issueType,
	}, nil
}