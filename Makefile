#  __    __   __________   ___      ___       _______   ______   .__   __.      ___       __
# |  |  |  | |   ____\  \ /  /     /   \     /  _____| /  __  \  |  \ |  |     /   \     |  |
# |  |__|  | |  |__   \  V  /     /  ^  \   |  |  __  |  |  |  | |   \|  |    /  ^  \    |  |
# |   __   | |   __|   >   <     /  /_\  \  |  | |_ | |  |  |  | |  . `  |   /  /_\  \   |  |
# |  |  |  | |  |____ /  .  \   /  _____  \ |  |__| | |  `--'  | |  |\   |  /  _____  \  |  `----.
# |__|  |__| |_______/__/ \__\ /__/     \__\ \______|  \______/  |__| \__| /__/     \__\ |_______|
#
#                _______   ______       __  ___      ___   .___________.    ___
#               /  _____| /  __  \     |  |/  /     /   \  |           |   /   \
#              |  |  __  |  |  |  |    |  '  /     /  ^  \ `---|  |----`  /  ^  \
#              |  | |_ | |  |  |  |    |    <     /  /_\  \    |  |      /  /_\  \
#              |  |__| | |  `--'  |    |  .  \   /  _____  \   |  |     /  _____  \
#               \______|  \______/     |__|\__\ /__/     \__\  |__|    /__/     \__\
#
USER := $(shell id -u)
GROUP := $(shell id -g)
DOCKER_COMPOSE := docker compose

.DEFAULT_GOAL:=help

##@ Help

.PHONY: help

help: ## Show this help screen.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Build
.PHONY: start stop down destroy full-wipe

status: ## Show containers status
	@docker compose ps

build: ## Build the project
	@-$(DOCKER_COMPOSE) up --build -d

start: ## Starts the containers
	@-$(DOCKER_COMPOSE) up -d --remove-orphans

stop: ## Stops the containers
	@-$(DOCKER_COMPOSE) stop

down: ## Go down the containers
	@-$(DOCKER_COMPOSE) down

destroy: ## Destroy containers and its volumes (all data will be lost)
	@-$(DOCKER_COMPOSE) down -v

full-wipe: destroy start ## Wipe all containers


