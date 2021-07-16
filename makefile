

.PHONY: build
build:
	@go build -o ./bin/search.out ./cmd/search

.PHONY: run
run: build
	@./bin/search.out