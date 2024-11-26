# /bin/sh

# service image
export BUILD_FROM_IMAGE=go-micro-saas/golang-base-image:latest
export RUN_SERVICE_IMAGE=go-micro-saas/golang-release-image:latest
echo "==> build service image BUILD_FROM_IMAGE : ${BUILD_FROM_IMAGE}"
echo "==> build service image RUN_SERVICE_IMAGE : ${RUN_SERVICE_IMAGE}"
docker build \
    --build-arg BUILD_FROM_IMAGE=${BUILD_FROM_IMAGE} \
    --build-arg RUN_SERVICE_IMAGE=${RUN_SERVICE_IMAGE} \
    --build-arg APP_DIR=app \
    --build-arg SERVICE_NAME=nodeid-service \
    --build-arg VERSION=latest \
    -t nodeid-service:latest \
    -f ./devops/docker-build/service-image/Dockerfile_service_image .