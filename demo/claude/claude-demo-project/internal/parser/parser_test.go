package parser

import (
	"testing"

	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/github"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		want        *ResourceInfo
		wantErr     bool
		errContains string
	}{
		// 合法的 Issue URL
		{
			name: "valid issue url",
			url:  "https://github.com/owner/repo/issues/123",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "repo",
				Number: 123,
				Type:   github.IssueTypeIssue,
			},
			wantErr: false,
		},
		// 合法的 Pull Request URL
		{
			name: "valid pull request url",
			url:  "https://github.com/owner/repo/pull/456",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "repo",
				Number: 456,
				Type:   github.IssueTypePR,
			},
			wantErr: false,
		},
		// 合法的 Discussion URL
		{
			name: "valid discussion url",
			url:  "https://github.com/owner/repo/discussions/789",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "repo",
				Number: 789,
				Type:   github.IssueTypeDiscussion,
			},
			wantErr: false,
		},
		// 无效的 URL（格式错误）
		{
			name:        "invalid url format",
			url:         "not-a-valid-url",
			want:        nil,
			wantErr:     true,
			errContains: "only GitHub URLs are supported",
		},
		// 不支持的 URL 类型（仓库主页）
		{
			name:        "unsupported url type - repository homepage",
			url:         "https://github.com/owner/repo",
			want:        nil,
			wantErr:     true,
			errContains: "invalid GitHub URL format",
		},
		// 不支持的 URL 类型（分支页面）
		{
			name:        "unsupported url type - branch page",
			url:         "https://github.com/owner/repo/tree/main",
			want:        nil,
			wantErr:     true,
			errContains: "unsupported resource type",
		},
		// 不支持的 URL 类型（提交页面）
		{
			name:        "unsupported url type - commit page",
			url:         "https://github.com/owner/repo/commit/abc123",
			want:        nil,
			wantErr:     true,
			errContains: "unsupported resource type",
		},
		// 空 URL
		{
			name:        "empty url",
			url:         "",
			want:        nil,
			wantErr:     true,
			errContains: "URL cannot be empty",
		},
		// 非 GitHub 域名
		{
			name:        "non-github domain",
			url:         "https://gitlab.com/owner/repo/issues/123",
			want:        nil,
			wantErr:     true,
			errContains: "only GitHub URLs are supported",
		},
		// 无效的资源编号（非数字）
		{
			name:        "invalid resource number - not a number",
			url:         "https://github.com/owner/repo/issues/abc",
			want:        nil,
			wantErr:     true,
			errContains: "invalid resource number",
		},
		// 无效的资源编号（负数）
		{
			name:        "invalid resource number - negative",
			url:         "https://github.com/owner/repo/issues/-123",
			want:        nil,
			wantErr:     true,
			errContains: "resource number must be positive",
		},
		// 无效的资源编号（零）
		{
			name:        "invalid resource number - zero",
			url:         "https://github.com/owner/repo/issues/0",
			want:        nil,
			wantErr:     true,
			errContains: "resource number must be positive",
		},
		// 不支持的资源类型
		{
			name:        "unsupported resource type",
			url:         "https://github.com/owner/repo/releases/1.0.0",
			want:        nil,
			wantErr:     true,
			errContains: "unsupported resource type",
		},
		// URL 包含查询参数（应该忽略）
		{
			name: "url with query parameters",
			url:  "https://github.com/owner/repo/issues/123?param=value",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "repo",
				Number: 123,
				Type:   github.IssueTypeIssue,
			},
			wantErr: false,
		},
		// URL 包含锚点（应该忽略）
		{
			name: "url with fragment",
			url:  "https://github.com/owner/repo/issues/123#some-section",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "repo",
				Number: 123,
				Type:   github.IssueTypeIssue,
			},
			wantErr: false,
		},
		// 带下划线的仓库名
		{
			name: "repo name with underscore",
			url:  "https://github.com/owner/my_repo/issues/123",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "my_repo",
				Number: 123,
				Type:   github.IssueTypeIssue,
			},
			wantErr: false,
		},
		// 带连字符的仓库名
		{
			name: "repo name with hyphen",
			url:  "https://github.com/owner/my-repo/issues/123",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "my-repo",
				Number: 123,
				Type:   github.IssueTypeIssue,
			},
			wantErr: false,
		},
		// 带数字的仓库名
		{
			name: "repo name with numbers",
			url:  "https://github.com/owner/repo123/issues/456",
			want: &ResourceInfo{
				Owner:  "owner",
				Repo:   "repo123",
				Number: 456,
				Type:   github.IssueTypeIssue,
			},
			wantErr: false,
		},
	}

	parser := NewParser()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.ParseURL(tt.url)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if err == nil {
					t.Error("ParseURL() expected error but got none")
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("ParseURL() error = %v, should contain %v", err.Error(), tt.errContains)
				}
				return
			}

			if got == nil && tt.want != nil {
				t.Error("ParseURL() returned nil when expected result")
				return
			}

			if got != nil && tt.want == nil {
				t.Error("ParseURL() returned result when expected nil")
				return
			}

			if got != nil && tt.want != nil {
				if got.Owner != tt.want.Owner {
					t.Errorf("ParseURL() Owner = %v, want %v", got.Owner, tt.want.Owner)
				}
				if got.Repo != tt.want.Repo {
					t.Errorf("ParseURL() Repo = %v, want %v", got.Repo, tt.want.Repo)
				}
				if got.Number != tt.want.Number {
					t.Errorf("ParseURL() Number = %v, want %v", got.Number, tt.want.Number)
				}
				if got.Type != tt.want.Type {
					t.Errorf("ParseURL() Type = %v, want %v", got.Type, tt.want.Type)
				}
			}
		})
	}
}

// contains 检查字符串是否包含子串
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && (s[:len(substr)] == substr || contains(s[1:], substr)))
}

func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Error("NewParser() returned nil")
	}

	// 验证返回的对象实现了 Parser 接口
	if _, ok := parser.(Parser); !ok {
		t.Error("NewParser() returned object does not implement Parser interface")
	}
}