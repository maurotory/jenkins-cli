GLOBAL_TESTS_DIR = $(shell pwd)
SHELL := /bin/bash

# Variables (optional)
IMAGE_NAME=jenkins/jenkins
CONTAINER_NAME=jenkins
FLAGS=-d --rm -p 8080:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home

# Default target
.PHONY: run-jenkins
run-jenkins: 
	docker run $(FLAGS) --name $(CONTAINER_NAME) $(IMAGE_NAME) 
