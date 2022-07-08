.PHONY: migrate up build run test docker-build

migrate:
	@scripts/migrate.sh

up:
	@docker-compose -f docker-compose.yml up

run:
	@scripts/run-all.sh

build:
	@go build -ldflags="-s -w" -o bin/url_shortener cmd/grpc/main.go

test:
	@scripts/run-integration-tests.sh

docker-build-push:
	@scripts/build-and-push-docker-images.sh
