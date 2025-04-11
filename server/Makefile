.DEFAULT_GOAL := test

DOCKER_IMAGE = olycall-server
DOCKER_TAG = latest
PORT = 8080

.PHONY: fmt
fmt:
	@gofumpt -w -l .
	@goimports -w -l .
	@golangci-lint run --fix

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@gotestsum -f testdox

.PHONY: build
build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: run
run:
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE):$(DOCKER_TAG) start

.PHONY: swag
swag:
	@swag init -g cmd/server/server.go
