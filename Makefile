NAME = server

run: build
	@./bin/$(NAME)

build: clean
	@go build -o bin/$(NAME) ./cmd/app

clean:
	@rm -rf bin
