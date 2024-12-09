# Variables (optional)
IMAGE_NAME=jenkins/jenkins
CONTAINER_NAME=jenkins
FLAGS=-p 8080:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home

# Default target
run:
	docker run --name $(CONTAINER_NAME) $(IMAGE_NAME) $(FLAGS)
