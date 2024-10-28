app_name := api

run: build
	@./bin/$(app_name)

build:
	@go build -o ./bin/$(app_name) ./cmd/$(app_name)/main.go

dev:
	air

install:
	@go install github.com/air-verse/air@latest
