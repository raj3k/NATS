# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## run/nats: run the cmd/nats application
.PHONY: run/nats
run/nats:
	go run ./cmd/nats

# ==================================================================================== #
# BUILD
# ==================================================================================== #
## build/nats: build the cmd/nats application
.PHONY: build/nats
build/nats:
	@echo 'Building cmd/nats...'
	go build -ldflags='-s' -o=./bin/nats ./cmd/nats