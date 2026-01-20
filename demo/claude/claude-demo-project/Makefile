.PHONY: build test web clean

# 构建主程序
build:
	go build -o bin/issue2md ./cmd/issue2md

# 运行所有测试
test:
	go test ./... -v

# 构建Web版本（未来扩展）
web:
	@echo "Web版本将在未来版本中实现"

# 清理构建产物
clean:
	rm -rf bin/

# 下载依赖
deps:
	go mod tidy

# 运行代码格式检查
fmt:
	go fmt ./...

# 运行代码检查
lint:
	go vet ./...