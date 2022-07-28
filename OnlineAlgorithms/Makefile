all: build basic

build:
	@go build -o build/main cmd/main.go

basic: build
	@go run cmd/main.go data/configs/pagingAllUni

scenario: build
	@go run cmd/main.go  data/configs/$(SCENARIO)
