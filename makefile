.PHONY: build run docker
build:
	GOFLAGS='' CGO_ENABLED=0 go build -o bin/ai-homelab-tools ./cmd/server

run:
	ADDR=:7070 go run ./cmd/server

docker:
	docker build -t ai-homelab-tools:local .

