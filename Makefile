build:
	@go build -o ~/Users/mac/Personal/golang/goapp cmd/main.go

run: build
	@~/Users/mac/Personal/golang/goapp