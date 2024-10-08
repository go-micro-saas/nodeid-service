# build-image
.PHONY: build
# build :-->: service image
build:
	#docker build -t nodeid-service:v1.0.0 -f ./devops/Dockerfile .
	#docker pull golang:1.22.8
	#docker pull debian:stable-20240926-slim
	docker build \
		--build-arg BUILD_FROM_IMAGE=golang:1.22.8 \
		--build-arg RUN_SERVICE_IMAGE=debian:stable-20240926-slim \
		--build-arg APP_DIR=app \
		--build-arg SERVICE_NAME=nodeid-service \
		--build-arg VERSION=latest \
		-t nodeid-service:latest \
		-f ./devops/Dockerfile .

