# build-image
.PHONY: build
# build :-->: service image
build:
	#docker build -t nodeid-service:v1.0.0 -f ./devops/Dockerfile .
	docker build \
		--build-arg SERVICE_NAME=nodeid-service \
		--build-arg VERSION=v1.0.0 \
		-t nodeid-service:v1.0.0 \
		-f ./devops/Dockerfile .

