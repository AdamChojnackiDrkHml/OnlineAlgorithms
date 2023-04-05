.PHONY: build basic, build, scenario, omg

all: build basic

build:
	@go build -o build/main cmd/main.go

basic: build
	@go run cmd/main.go data/configs/pagingAllUni

scenario: build
	@go run cmd/main.go data/configs/$(SCENARIO)

TESTS_DIR = ./data/configs

omg:
	$(foreach file, $(wildcard $(TESTS_DIR)/*.yml), go run cmd/main.go -f $(file);)