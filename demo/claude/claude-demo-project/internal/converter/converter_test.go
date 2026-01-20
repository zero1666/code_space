package converter

import (
	"strings"
	"testing"
	"time"

	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/github"
)

func TestConvert(t *testing.T) {
	createdAt := time.Date(2024, 1, 19, 14, 30, 0, 0, time.UTC)
	updatedAt := time.Date(2024, 1, 20, 10, 15, 0, 0, time.UTC)

	tests := []struct {
		name            string
		data            *github.IssueData
		opts            ConvertOptions
		wantContains    []string
		wantNotContains []string
		wantErr         bool
	}{
		{
			name: "convert issue with comments",
			data: &github.IssueData{
				Title: "Add dark mode support",
				URL:   "https://github.com/owner/repo/issues/123",
				Author: github.User{
					Login: "alice",
					URL:   "https://github.com/alice",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				State:     "open",
				Type:      github.IssueTypeIssue,
				Body:      "We should add dark mode support.",
				Comments: []github.Comment{
					{
						Author: github.User{
							Login: "bob",
							URL:   "https://github.com/bob",
						},
						CreatedAt: createdAt.Add(1 * time.Hour),
						Body:      "Great idea!",
					},
				},
			},
			opts: ConvertOptions{},
			wantContains: []string{
				"title: \"Add dark mode support\"",
				"type: \"issue\"",
				"state: \"open\"",
				"author: \"alice\"",
				"# Add dark mode support",
				"We should add dark mode support.",
				"bob",
				"Great idea!",
			},
			wantNotContains: []string{
				"https://github.com/alice", // 默认不显示用户链接
			},
			wantErr: false,
		},
		{
			name: "convert PR with merged state",
			data: &github.IssueData{
				Title: "feat: implement dark mode",
				URL:   "https://github.com/owner/repo/pull/456",
				Author: github.User{
					Login: "alice",
					URL:   "https://github.com/alice",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				State:     "merged",
				Type:      github.IssueTypePR,
				Body:      "Implements dark mode toggle.",
				Comments:  []github.Comment{},
			},
			opts: ConvertOptions{},
			wantContains: []string{
				"type: \"pr\"",
				"state: \"merged\"",
				"feat: implement dark mode",
			},
			wantErr: false,
		},
		{
			name: "convert with user links enabled",
			data: &github.IssueData{
				Title: "Test Issue",
				URL:   "https://github.com/owner/repo/issues/789",
				Author: github.User{
					Login: "alice",
					URL:   "https://github.com/alice",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				State:     "open",
				Type:      github.IssueTypeIssue,
				Body:      "Test body",
				Comments:  []github.Comment{},
			},
			opts: ConvertOptions{
				IncludeUserLinks: true,
			},
			wantContains: []string{
				"[alice](https://github.com/alice)",
			},
			wantErr: false,
		},
		{
			name: "convert with reactions enabled",
			data: &github.IssueData{
				Title: "Test Issue",
				URL:   "https://github.com/owner/repo/issues/789",
				Author: github.User{
					Login: "alice",
					URL:   "https://github.com/alice",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				State:     "open",
				Type:      github.IssueTypeIssue,
				Body:      "Test body",
				Comments: []github.Comment{
					{
						Author: github.User{
							Login: "bob",
							URL:   "https://github.com/bob",
						},
						CreatedAt: createdAt,
						Body:      "Nice!",
						Reactions: &github.Reactions{
							TotalCount: 5,
							PlusOne:    3,
							Heart:      2,
						},
					},
				},
			},
			opts: ConvertOptions{
				IncludeReactions: true,
			},
			wantContains: []string{
				"+1: 3",
				"heart: 2",
			},
			wantErr: false,
		},
		{
			name: "convert discussion with answer",
			data: &github.IssueData{
				Title: "How to implement feature X?",
				URL:   "https://github.com/owner/repo/discussions/999",
				Author: github.User{
					Login: "alice",
					URL:   "https://github.com/alice",
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				State:     "open",
				Type:      github.IssueTypeDiscussion,
				Body:      "I need help with feature X.",
				Comments: []github.Comment{
					{
						Author: github.User{
							Login: "bob",
							URL:   "https://github.com/bob",
						},
						CreatedAt: createdAt,
						Body:      "Here is how you do it...",
						IsAnswer:  true,
					},
				},
			},
			opts: ConvertOptions{},
			wantContains: []string{
				"type: \"discussion\"",
				"Accepted Answer",
			},
			wantErr: false,
		},
		{
			name:    "nil data",
			data:    nil,
			opts:    ConvertOptions{},
			wantErr: true,
		},
	}

	converter := NewConverter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := converter.Convert(tt.data, tt.opts)

			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			for _, want := range tt.wantContains {
				if !strings.Contains(got, want) {
					t.Errorf("Convert() output should contain %q, got:\n%s", want, got)
				}
			}

			for _, notWant := range tt.wantNotContains {
				if strings.Contains(got, notWant) {
					t.Errorf("Convert() output should NOT contain %q, got:\n%s", notWant, got)
				}
			}
		})
	}
}

func TestFormatTimestamp(t *testing.T) {
	tests := []struct {
		name string
		time time.Time
		want string
	}{
		{
			name: "format UTC time",
			time: time.Date(2024, 1, 19, 14, 30, 0, 0, time.UTC),
			want: "2024-01-19 14:30:00 UTC",
		},
		{
			name: "format zero time",
			time: time.Time{},
			want: "0001-01-01 00:00:00 UTC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatTimestamp(tt.time)
			if got != tt.want {
				t.Errorf("FormatTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConverter(t *testing.T) {
	converter := NewConverter()
	if converter == nil {
		t.Error("NewConverter() returned nil")
	}
}