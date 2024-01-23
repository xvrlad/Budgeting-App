build:
	@go build -o bin/budgeting-app

run: build
	@./bin/budgeting-app

test:
	@go test -v ./...