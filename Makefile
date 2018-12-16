HTTP_PORT=8000
DOCKER_IMAGE=ukko/cowsay
DOCKER_TAG=latest
ENTRY=./bin/entry

help:
	@echo "Developers commands:"
	@echo "make run             Run app"
	@echo "make build           Build app sources"
	@echo "make vendor          Update vendor dependencies"
	@echo "make browser         Open link in browser"
	@echo ""
	@echo "Docker commands:"
	@echo "make docker.build    Build docker container"
	@echo "make docker.push     Push docker container"
	@echo "make docker.run      Run docker container"

build:
	go build -o ${ENTRY} src/main.go

run: build
	HTTP_PORT=${HTTP_PORT} ${ENTRY}

vendor:
	dep ensure

browser:
	xdg-open http://localhost:${HTTP_PORT}/

docker.build: build
	docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .

docker.push: docker.build
	docker push ${DOCKER_IMAGE}:${DOCKER_TAG}

docker.run:
	docker run --publish=${HTTP_PORT}:${HTTP_PORT} ${DOCKER_IMAGE}:${DOCKER_TAG}

docker.exec:
	docker exec -it `docker ps --filter=ancestor=${DOCKER_IMAGE}:${DOCKER_TAG} -q` bash