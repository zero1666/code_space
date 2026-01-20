package github

import (
	"context"
	"strings"
	"testing"
)

func TestFetchIssue_Validation(t *testing.T) {
	tests := []struct {
		name        string
		owner       string
		repo        string
		number      int
		wantErr     bool
		errContains string
	}{
		{
			name:        "empty owner",
			owner:       "",
			repo:        "repo",
			number:      123,
			wantErr:     true,
			errContains: "owner cannot be empty",
		},
		{
			name:        "empty repo",
			owner:       "owner",
			repo:        "",
			number:      123,
			wantErr:     true,
			errContains: "repo cannot be empty",
		},
		{
			name:        "invalid issue number - zero",
			owner:       "owner",
			repo:        "repo",
			number:      0,
			wantErr:     true,
			errContains: "number must be positive",
		},
		{
			name:        "invalid issue number - negative",
			owner:       "owner",
			repo:        "repo",
			number:      -1,
			wantErr:     true,
			errContains: "number must be positive",
		},
	}

	client := NewGitHubClient("test-token")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := client.FetchIssue(ctx, tt.owner, tt.repo, tt.number)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchIssue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("FetchIssue() error = %v, should contain %v", err.Error(), tt.errContains)
				}
			}
		})
	}
}

func TestFetchPullRequest_Validation(t *testing.T) {
	tests := []struct {
		name        string
		owner       string
		repo        string
		number      int
		wantErr     bool
		errContains string
	}{
		{
			name:        "empty owner",
			owner:       "",
			repo:        "repo",
			number:      123,
			wantErr:     true,
			errContains: "owner cannot be empty",
		},
		{
			name:        "empty repo",
			owner:       "owner",
			repo:        "",
			number:      123,
			wantErr:     true,
			errContains: "repo cannot be empty",
		},
		{
			name:        "invalid number",
			owner:       "owner",
			repo:        "repo",
			number:      0,
			wantErr:     true,
			errContains: "number must be positive",
		},
	}

	client := NewGitHubClient("test-token")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := client.FetchPullRequest(ctx, tt.owner, tt.repo, tt.number)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchPullRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("FetchPullRequest() error = %v, should contain %v", err.Error(), tt.errContains)
				}
			}
		})
	}
}

func TestFetchDiscussion_Validation(t *testing.T) {
	tests := []struct {
		name        string
		owner       string
		repo        string
		number      int
		wantErr     bool
		errContains string
	}{
		{
			name:        "empty owner",
			owner:       "",
			repo:        "repo",
			number:      123,
			wantErr:     true,
			errContains: "owner cannot be empty",
		},
		{
			name:        "valid params but not implemented",
			owner:       "owner",
			repo:        "repo",
			number:      123,
			wantErr:     true,
			errContains: "GraphQL API",
		},
	}

	client := NewGitHubClient("test-token")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			_, err := client.FetchDiscussion(ctx, tt.owner, tt.repo, tt.number)

			if (err != nil) != tt.wantErr {
				t.Errorf("FetchDiscussion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("FetchDiscussion() error = %v, should contain %v", err.Error(), tt.errContains)
				}
			}
		})
	}
}

func TestNewGitHubClient(t *testing.T) {
	tests := []struct {
		name  string
		token string
	}{
		{
			name:  "create client with token",
			token: "test-token",
		},
		{
			name:  "create client without token",
			token: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewGitHubClient(tt.token)
			if client == nil {
				t.Error("NewGitHubClient() returned nil")
			}
		})
	}
}