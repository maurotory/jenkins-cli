JENKINS_CLI_DIR = $(shell pwd)
SHELL := /bin/bash

# Variables (optional)
IMAGE_NAME=jenkins/jenkins
CONTAINER_NAME=jenkins
FLAGS=-d --rm -p 8080:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home

.PHONY: all
all: jenkins-cli

.PHONY:jenkins-cli 
jenkins-cli: check-go check-os
	@echo Build jctl
	go install -mod=readonly ./cmd/jctl
	go build -mod=readonly -a -o \
	bin/jctl cmd/jctl/main.go

.PHONY: check-go
check-go:
	@which go > /dev/null 2>&1 && echo "Go is installed: $$(go version)" || (echo "Go is not installed" && exit 1)

.PHONY: check-docker
check-docker:
	@which docker > /dev/null 2>&1 && echo "Docker is installed" || (echo "Docker is not installed" && exit 1)

.PHONY: check-os
check-os:
	@echo "Checking Operating System..."
	@uname -s | tr '[:upper:]' '[:lower:]' | grep -q "linux" && echo "Operating System: Linux" || \
	(uname -s | tr '[:upper:]' '[:lower:]' | grep -q "darwin" && echo "Operating System: macOS" || \
	 (echo "Operating System: Unknown" && exit 1))


.PHONY: run-jenkins
run-jenkins: check-docker
	docker run $(FLAGS) --name $(CONTAINER_NAME) $(IMAGE_NAME) 

