build:
	@go build -o bin/proceso

run: build
	@./bin/proceso

test:
	@go test -v ./...

lint:
	@golangci-lint run --enable-all

format:
	@gofmt -s -l -w .
