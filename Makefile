build: 
	@go build -o bin/management-api ./cmd/main.go

run: build
	@./bin/management-api