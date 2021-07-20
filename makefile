

.PHONY: install-linter
install-linter:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.41.1
	golangci-lint --version

.PHONY: build
build:
	@go build -o ./bin/search.out ./cmd/search

.PHONY: run
run: build
	@./bin/search.out

.PHONY: lint
lint:
	golangci-lint run