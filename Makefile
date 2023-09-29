# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## run/nats: run the cmd/nats application
.PHONY: run/nats
run/nats:
	go run ./cmd/nats

## test: run all tests
.PHONY: test
test:
	go test -v -race -buildvcs ./...

## test/cover: run all tests and display coverage
.PHONY: test/cover
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# ==================================================================================== #
# BUILD
# ==================================================================================== #
## build/nats: build the cmd/nats application
.PHONY: build/nats
build/nats:
	@echo 'Building cmd/nats...'
	go build -ldflags='-s' -o=./bin/nats ./cmd/nats