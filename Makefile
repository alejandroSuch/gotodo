
.PHONY: prepare
prepare:
	go fmt ./...
	go vet ./...
	go mod tidy

build-server: prepare
	rm -rf dist/server
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/server/darwin-x86_64/todo-server cmd/server/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/server/darwin-arm64/todo-server cmd/server/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o dist/server/linux-i386/todo-server cmd/server/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/server/linux-x86_64/todo-server cmd/server/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dist/server/windows-i386/todo-server.exe cmd/server/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/server/windows-x86_64/todo-server.exe cmd/server/main.go

build-cli: prepare
	rm -rf dist/cli
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o dist/cli/darwin-x86_64/todo-server cmd/cli/main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o dist/cli/darwin-arm64/todo-server cmd/cli/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o dist/cli/linux-i386/todo-server cmd/cli/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/cli/linux-x86_64/todo-server cmd/cli/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o dist/cli/windows-i386/todo-server.exe cmd/cli/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o dist/cli/windows-x86_64/todo-server.exe cmd/cli/main.go

build-ui: prepare
	rm -rf dist/ui
	GOOS=darwin GOARCH=amd64 go build -o dist/ui/darwin-x86_64/todo-server cmd/ui/main.go

build-all: build-server build-cli build-ui
