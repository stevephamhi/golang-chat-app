build:
	@go build -o ~/Users/mac/Personal/golang/chat-app cmd/main.go

run: build
	@~/Users/mac/Personal/golang/chat-app