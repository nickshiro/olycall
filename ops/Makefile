# Load only COMPOSE_ENV from .env
export COMPOSE_ENV := $(shell grep ^COMPOSE_ENV .env | cut -d '=' -f2-)
COMPOSE_ENV ?= dev

COMPOSE_FILE=compose.$(COMPOSE_ENV).yml

.PHONY: down
down:
	docker compose -f ${COMPOSE_FILE} down

.PHONY: clean
clean:
	docker compose -f ${COMPOSE_FILE} down -v --remove-orphans

.PHONY: up
up:
	docker compose -f ${COMPOSE_FILE} up

.PHONY: upd
upd:
	docker compose -f ${COMPOSE_FILE} up -d

.PHONY: build
build:
	docker compose -f ${COMPOSE_FILE} build

.PHONY: run
run:
	docker compose -f ${COMPOSE_FILE} up --build

.PHONY: rund
rund:
	docker compose -f ${COMPOSE_FILE} up --build -d

.PHONY: pull
pull:
	docker compose -f ${COMPOSE_FILE} pull
	docker compose -f ${COMPOSE_FILE} up -d --force-recreate
	docker image prune -a -f
