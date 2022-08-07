# Common variables
VERSION := 0.0.1
BUILD_INFO := Manual build 

# Things you don't want to change
REPO_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
# Tools
GOLINT_PATH := $(REPO_DIR)/bin/golangci-lint              # Remove if not using Go
AIR_PATH := $(REPO_DIR)/bin/air                           # Remove if not using Go

.PHONY: help image push build run lint lint-fix
.DEFAULT_GOAL := help

help: ## 💬 This help message :)
	@figlet $@ || true
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

install-tools: ## 🔮 Install dev tools into project bin directory
	@figlet $@ || true
	@$(GOLINT_PATH) > /dev/null 2>&1 || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin/
	@$(AIR_PATH) -v > /dev/null 2>&1 || curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh
	
lint: ## 🌟 Lint & format, will not fix but sets exit code on error
	@figlet $@ || true
	$(GOLINT_PATH) run --modules-download-mode=mod ./...

lint-fix: ## 🔍 Lint & format, will try to fix errors and modify code
	@figlet $@ || true
	golangci-lint run --modules-download-mode=mod --fix ./...

build: ## 🔨 Run a local build without a container
	@figlet $@ || true
	go build -o ./bin/game ./...
	go build -o ./bin/game.exe ./...

run: ## 🏃‍♂️ Run application, used for local development
	@figlet $@ || true
	$(AIR_PATH) -c .air.toml