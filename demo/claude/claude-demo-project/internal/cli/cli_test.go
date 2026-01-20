package cli

import (
	"strings"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		wantURL          string
		wantOutputFile   string
		wantReactions    bool
		wantUserLinks    bool
		wantErr          bool
	}{
		{
			name:           "url as positional argument",
			args:           []string{"https://github.com/owner/repo/issues/123"},
			wantURL:        "https://github.com/owner/repo/issues/123",
			wantOutputFile: "",
			wantReactions:  false,
			wantUserLinks:  false,
			wantErr:        false,
		},
		{
			name:           "url and output file as positional arguments",
			args:           []string{"https://github.com/owner/repo/issues/123", "output.md"},
			wantURL:        "https://github.com/owner/repo/issues/123",
			wantOutputFile: "output.md",
			wantReactions:  false,
			wantUserLinks:  false,
			wantErr:        false,
		},
		{
			name:           "with enable-reactions flag",
			args:           []string{"-enable-reactions", "https://github.com/owner/repo/issues/123"},
			wantURL:        "https://github.com/owner/repo/issues/123",
			wantOutputFile: "",
			wantReactions:  true,
			wantUserLinks:  false,
			wantErr:        false,
		},
		{
			name:           "with enable-user-links flag",
			args:           []string{"-enable-user-links", "https://github.com/owner/repo/issues/123"},
			wantURL:        "https://github.com/owner/repo/issues/123",
			wantOutputFile: "",
			wantReactions:  false,
			wantUserLinks:  true,
			wantErr:        false,
		},
		{
			name:           "with all flags",
			args:           []string{"-enable-reactions", "-enable-user-links", "https://github.com/owner/repo/issues/123", "output.md"},
			wantURL:        "https://github.com/owner/repo/issues/123",
			wantOutputFile: "output.md",
			wantReactions:  true,
			wantUserLinks:  true,
			wantErr:        false,
		},
		{
			name:           "no arguments",
			args:           []string{},
			wantURL:        "",
			wantOutputFile: "",
			wantReactions:  false,
			wantUserLinks:  false,
			wantErr:        false, // ParseFlags 本身不报错，Validate 会报错
		},
		{
			name:    "invalid flag",
			args:    []string{"-invalid-flag", "https://github.com/owner/repo/issues/123"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config, err := ParseFlags(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if config.URL != tt.wantURL {
				t.Errorf("ParseFlags() URL = %v, want %v", config.URL, tt.wantURL)
			}
			if config.OutputFile != tt.wantOutputFile {
				t.Errorf("ParseFlags() OutputFile = %v, want %v", config.OutputFile, tt.wantOutputFile)
			}
			if config.IncludeReactions != tt.wantReactions {
				t.Errorf("ParseFlags() IncludeReactions = %v, want %v", config.IncludeReactions, tt.wantReactions)
			}
			if config.IncludeUserLinks != tt.wantUserLinks {
				t.Errorf("ParseFlags() IncludeUserLinks = %v, want %v", config.IncludeUserLinks, tt.wantUserLinks)
			}
		})
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		token       string
		wantErr     bool
		errContains string
	}{
		{
			name:    "valid config",
			url:     "https://github.com/owner/repo/issues/123",
			token:   "test-token",
			wantErr: false,
		},
		{
			name:        "empty url",
			url:         "",
			token:       "test-token",
			wantErr:     true,
			errContains: "URL is required",
		},
		{
			name:        "invalid url format",
			url:         "not-a-github-url",
			token:       "test-token",
			wantErr:     true,
			errContains: "invalid GitHub URL",
		},
		{
			name:        "empty token",
			url:         "https://github.com/owner/repo/issues/123",
			token:       "",
			wantErr:     true,
			errContains: "token is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := &Config{
				URL:         tt.url,
				GitHubToken: tt.token,
			}

			err := config.Validate()

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err != nil {
				if !strings.Contains(err.Error(), tt.errContains) {
					t.Errorf("Validate() error = %v, should contain %v", err.Error(), tt.errContains)
				}
			}
		})
	}
}

func TestNewCLI(t *testing.T) {
	cli := NewCLI(nil, nil, nil)
	if cli == nil {
		t.Error("NewCLI() returned nil")
	}
}