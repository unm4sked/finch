run-dev:
	go run ./cmd/api/main.go
format:
	go fmt ./...
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/api cmd/api/main.go