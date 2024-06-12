.PHONY: default run build test docs clean

APP_NAME=gooportunities

default: run-with-docs

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@swag init
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs