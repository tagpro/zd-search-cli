

.PHONY: install/dependencies
install/dependencies:
	# binary will be $(go env GOPATH)/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.41.1
	golangci-lint --version
	# installing gomock mockgen
	go install github.com/golang/mock/mockgen@v1.6.0

.PHONY: build
build:
	@go build -o ./bin/search.out ./cmd/search

.PHONY: run
run: build
	@./bin/search.out

.PHONY: test
test:
	go test ./...

.PHONY: test/generate
test/generate:  ## Generate go mocks
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run