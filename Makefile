#===================#
#== Env Variables ==#
#===================#
include .env
DOCKER_COMPOSE_FILE ?= deployment/docker-compose.yml

#===================#
#======= Boot ======#
#===================#
# Run with whole stack
run:
	@docker compose -f ${DOCKER_COMPOSE_FILE} --env-file .env --profile app --profile dependencies up --build

# Run dependencies only
pre-run:
	@docker compose --profile dependencies -f ${DOCKER_COMPOSE_FILE} --env-file .env up --build

# Run dependencies only
clear:
	@docker compose -f ${DOCKER_COMPOSE_FILE} --env-file .env down --rmi local -v --remove-orphans

# Install Golang dependencies
install:
	@go mod tidy

# Run in Prod mode
.PHONY: services run-product-service run-order-service
services: run-product-service run-order-service

run-product-service:
	@go run ./cmd/product_service/main.go

run-order-service:
	@go run ./cmd/order_service/main.go

# Run in dev mode
.PHONY: services-dev run-order-service-dev run-order-service-dev
services-dev: run-product-service-dev run-order-service-dev

run-product-service-dev:
	@~/go/bin/air -c .air.product.toml

run-order-service-dev:
	@~/go/bin/air -c .air.order.toml

# Build binary
build:
	@go build -o ./bin/server ./cmd/server/main.go

# Test
test:
	@go test -v ./...


#===================#
#=== Migrations ====#
#===================#
DATABASE_CONNECTION ?= "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable"
MIGRATION_PATH ?= ${CURDIR}/migrations
CONTAINER_MIGRATION_PATH ?= /migrations/
MIGRATION_VOLUME ?= ${MIGRATION_PATH}:${CONTAINER_MIGRATION_PATH}
# Migration down with migration_msg = msg for all

migration-create:
	@docker run \
	--rm \
	-v ${MIGRATION_VOLUME} \
	--network host migrate/migrate \
	-path=${CONTAINER_MIGRATION_PATH} \
	-database ${DATABASE_CONNECTION} \
	create -ext sql -dir /migrations $(migration_msg)

# Migration down with step = N / {{space}} for all
migration-up:
	@docker run \
	--rm \
	-v ${MIGRATION_VOLUME} \
	--network host migrate/migrate \
	-path=${CONTAINER_MIGRATION_PATH} \
	-database ${DATABASE_CONNECTION} \
	up ${step}

# Migration down with step = N / -all
migration-down:
	@echo y | docker run \
	--rm \
	-v ${MIGRATION_VOLUME} \
	--network host migrate/migrate \
	-path=${CONTAINER_MIGRATION_PATH} \
	-database ${DATABASE_CONNECTION} \
	down ${step}