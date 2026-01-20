package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/cli"
	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/converter"
	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/github"
	"github.com/zero1666/code_space/demo/claude/claude-demo-project/internal/parser"
)

const (
	exitCodeSuccess = 0
	exitCodeError   = 1
)

func main() {
	os.Exit(run())
}

func run() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 设置信号处理，支持优雅退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()

	// 从环境变量获取 GitHub Token
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		fmt.Fprintln(os.Stderr, "Error: GITHUB_TOKEN environment variable is required")
		printUsage()
		return exitCodeError
	}

	// 检查是否有参数
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: URL is required")
		printUsage()
		return exitCodeError
	}

	// 创建依赖实例
	githubClient := github.NewGitHubClient(githubToken)
	urlParser := parser.NewParser()
	markdownConverter := converter.NewConverter()

	// 创建 CLI 实例
	cliInstance := cli.NewCLI(githubClient, urlParser, markdownConverter)

	// 运行 CLI
	if err := cliInstance.Run(ctx, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return exitCodeError
	}

	return exitCodeSuccess
}

func printUsage() {
	fmt.Fprint(os.Stderr, `
Usage: issue2md [flags] <url> [output_file]

Arguments:
  <url>           GitHub Issue/PR/Discussion URL (required)
  [output_file]   Output file path (optional, defaults to stdout)

Flags:
  -enable-reactions    Include reactions statistics
  -enable-user-links   Render usernames as GitHub profile links

Environment Variables:
  GITHUB_TOKEN    GitHub Personal Access Token (required)

Examples:
  # Output to stdout
  issue2md https://github.com/owner/repo/issues/123

  # Output to file
  issue2md https://github.com/owner/repo/pull/456 output.md

  # Enable all features
  issue2md -enable-reactions -enable-user-links https://github.com/owner/repo/discussions/789 discussion.md
`)
}