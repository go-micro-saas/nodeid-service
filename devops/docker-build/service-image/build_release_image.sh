# /bin/sh

# release image
export FROM_IMAGE_NAME=debian:latest
export BASE_IMAGE_NAME=go-micro-saas/golang-release-image:latest
export IS_EXIST_BASE_IMAGE=0
docker images --format "{{.Repository}}:{{.Tag}}" | grep -q "^${BASE_IMAGE_NAME}$" && export IS_EXIST_BASE_IMAGE=1 || echo "CANNOT FOUND ${BASE_IMAGE_NAME}"
if [ "${IS_EXIST_BASE_IMAGE}" -eq 1 ]; then
  export FROM_IMAGE_NAME=${BASE_IMAGE_NAME}
else
  echo "==> docker pull ${FROM_IMAGE_NAME}"
  docker pull ${FROM_IMAGE_NAME}
fi
echo "==> build release image FROM_IMAGE_NAME : ${FROM_IMAGE_NAME}"

echo "==> build release image : ${BASE_IMAGE_NAME}"
docker build \
		--build-arg BUILD_FROM_IMAGE=${FROM_IMAGE_NAME} \
		-t ${BASE_IMAGE_NAME} \
		-f ./devops/docker-build/service-image/Dockerfile_release_image .