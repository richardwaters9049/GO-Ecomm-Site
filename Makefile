build:
	@go build -o bin/go-eccom

run: build
	@./bin/go-eccom

test: 
	@go test -v ./...

