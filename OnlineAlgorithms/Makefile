.PHONY: build basic

all: build basic

build:
	@go build -o build/main cmd/main.go

basic: build
	@go run cmd/main.go -f data/configs/pagingAllUni

scenario: build
	@go run cmd/main.go  -f data/configs/$(SCENARIO)

handdebug: build
	@go run cmd/main.go