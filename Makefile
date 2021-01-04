.PHONY:build
build:
	go build -v ./cmd/gateway/main.go

.DEFAULT_GOAL:=build