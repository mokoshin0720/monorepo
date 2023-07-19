.SHELL := /bin/bash

init: ## setup docker build, network, and databases
	bash hack/docker-run.sh

server: init migrate
	docker-compose up monorepo_api

.PHONY: migrate
migrate: init
	docker-compose run --rm monorepo_migration

down:
	docker-compose down --remove-orphans