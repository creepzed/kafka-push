DOCKER_COMPOSE = docker-compose

.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo ""
	@echo "Caution: those targets are optimized for docker 19+"
	@echo ""
	@echo "Please add your custom Makefile in the directory "make-file". They will be automatically loaded!"
	@echo ""

.PHONY: start
start:
	APP_ENV=dev $(MAKE) up

.PHONY: stop
stop:
	$(MAKE) down

.PHONY: logs
logs:
	$(MAKE) show-log

.PHONY: up
up:
	$(DOCKER_COMPOSE) up -d --remove-orphan ${C}

.PHONY: down
down:
	$(DOCKER_COMPOSE) down -v

.PHONY: show-log
show-log:
	$(DOCKER_COMPOSE) logs -f