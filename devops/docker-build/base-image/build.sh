# /bin/sh

export FROM_IMAGE_NAME=golang:1.22.8
export BASE_IMAGE_NAME=go-micro-saas/golang-base-image:latest
export IS_EXIST_BASE_IMAGE=0
docker images --format "{{.Repository}}:{{.Tag}}" | grep -q "^${BASE_IMAGE_NAME}$" && export IS_EXIST_BASE_IMAGE=1 || echo "CANNOT FOUND ${BASE_IMAGE_NAME}"
if [ "${IS_EXIST_BASE_IMAGE}" -eq 1 ]; then
  export FROM_IMAGE_NAME=${BASE_IMAGE_NAME}
fi
echo "==> FROM_IMAGE_NAME : $FROM_IMAGE_NAME"

