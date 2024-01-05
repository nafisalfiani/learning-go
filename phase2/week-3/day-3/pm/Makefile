.PHONY: swag
swag:
	@`go env GOPATH`/bin/swag init

.PHONY: build
build:
	@go build -o ./build/app ./main.go

.PHONY: run
run: swag build
	@./build/app