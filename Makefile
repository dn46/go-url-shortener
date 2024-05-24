build:
	@go build -o bin/shortener

run: build
	@./bin/shortener

test:
	@go test -v ./...