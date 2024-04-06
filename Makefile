#----
# Config
#----

ifneq (,$(wildcard .env))
	include .env
	export
	ENV_FILE_PARAM := --env-file .env
endif

#----
# Variables
#----

PROJECT_NAME := note

COMMON_COMPOSE := -f docker/compose.yaml
COMPOSE_PROJECT_NAME := --project-name $(PROJECT_NAME)

DEV_COMPOSE := $(COMMON_COMPOSE) -f docker/compose.dev.yaml $(COMPOSE_PROJECT_NAME)
PRE_COMPOSE := $(COMMON_COMPOSE) -f docker/compose.pre.yaml $(COMPOSE_PROJECT_NAME)
PRO_COMPOSE := $(COMMON_COMPOSE) -f docker/compose.prod.yaml $(COMPOSE_PROJECT_NAME)

#----
# Docker Compose
#----

.PHONY: start/dev
start/dev: checkEnv ## Start development environment
	@echo "Starting development environment..."
	docker-compose $(DEV_COMPOSE) up --build

.PHONY: start/pre
start/pre: checkEnv ## Start pre-production environment
	@echo "Starting pre-production environment..."
	docker-compose $(PRE_COMPOSE) up --build

#----
# Scripts
#----

.PHONY: checkEnv
checkEnv: ## checks for empty or unset env variables
	go run scripts/main.go

#----
# Others
#----

.DEFAULT_GOAL := help

.PHONY: help
help:
	@awk -F ':|##' '/^[^\t].+:.*##/ { printf "\033[36mmake %-28s\033[0m -%s\n", $$1, $$NF }' $(MAKEFILE_LIST) | sort
