NAME = server

run: build
	@./bin/$(NAME)

build: clean
	@go build -o bin/$(NAME) ./cmd/app

clean:
	@rm -rf bin

unit-test:
	@go clean -testcache
	@go test `go list ./... | grep -v ./cmd/app | grep -v ./internals/database`

coverage:
	@go clean -testcache
	@go test `go list ./... | grep -v ./cmd/app | grep -v ./internals/database` -coverprofile=coverage.out
	@go tool cover -html=coverage.out
