# Makefile for the Go Project

# 애플리케이션 이름과 Go 컴파일러 경로를 변수로 정의합니다.
APP_NAME=my-graphql-server
GO_CMD=go

# .PHONY는 파일 이름이 아닌 타겟(명령)임을 명시합니다.
.PHONY: all run build test lint tidy clean

all: build

# 'make run': 애플리케이션을 실행합니다.
run:
	@echo "🔥 Starting server..."
	@$(GO_CMD) run ./cmd/server

# 'make build': 프로덕션용 바이너리를 빌드합니다.
build: tidy
	@echo "🔨 Building binary..."
	@$(GO_CMD) build -o bin/$(APP_NAME) ./cmd/server

# 'make test': 모든 테스트를 실행합니다.
test:
	@echo "🧪 Running tests..."
	@$(GO_CMD) test -v ./...

# 'make lint': golangci-lint를 실행하여 코드 품질을 검사합니다.
lint:
	@echo "🧹 Running linter..."
	@golangci-lint run

# 'make tidy': go.mod 파일을 정리합니다.
tidy:
	@echo "🧹 Tidying go.mod..."
	@$(GO_CMD) mod tidy

generate: tidy
	@echo "📄 Generating GraphQL code..."
	@go get github.com/99designs/gqlgen
	@go run github.com/99designs/gqlgen generate

# 'make clean': 빌드 결과물과 생성된 gqlgen 코드를 모두 삭제합니다.
clean: clean-bin clean-graph
	@echo "🗑️ All generated files and binaries are cleaned."

# 'make clean-bin': 빌드된 바이너리만 삭제합니다.
clean-bin:
	@echo "🗑️ Cleaning up binaries..."
	@rm -rf bin

# 'make clean-graph': gqlgen으로 생성된 코드만 삭제합니다.
clean-graph:
	@echo "🧹 Cleaning up generated GraphQL files..."
	# rm -rf: 지정된 폴더와 그 안의 모든 파일을 강제로 삭제하는 명령어입니다.
	# internal/graph 폴더 전체를 삭제하여 깨끗한 상태로 만듭니다.
	@rm -rf internal/graph