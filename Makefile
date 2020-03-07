.DEFAULT_GOAL := explain
.PHONY: explain
explain:
	### Welcome
	#
	#        __          __        _    ____              _    
	#        \ \        / /       | |  |  _ \            | |   
	#       __\ \  /\  / /__  _ __| | _| |_) | ___   ___ | | __
	#      / _ \ \/  \/ / _ \| '__| |/ /  _ < / _ \ / _ \| |/ /
	#     |  __/\  /\  / (_) | |  |   <| |_) | (_) | (_) |   < 
	#      \___| \/  \/ \___/|_|  |_|\_\____/ \___/ \___/|_|\_\                
	#
	#
	### Installation
	#
	### Targets
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run the beggo application in  http://localhost:8080
	bee run

.PHONY: install
install: ## install all project dependency
	go get ./...

.PHONY: vet
vet: ## Vet the code
	go vet ./...

.PHONY: lint
lint: ## Lint the code
	golint -set_exit_status ./...

.PHONY: security-check
security-check: ## Inspect code for security vulnerabilities
	gosec ./...

.PHONY: build
build: ## Inspect code for security vulnerabilities
	go build -o e-work-book main.go

.PHONY: backup-db
backup-db: ## To create a MySQL Dump
	sh backup/create-mysql-dump.sh