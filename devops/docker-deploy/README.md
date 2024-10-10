# deploy

```shell
# general config
go run ./app/nodeid-service/cmd/store-configuration/... \
  -conf=./app/nodeid-service/configs \
  -source_dir ./devops/docker-deploy/general-configs \
  -store_dir go-micro-saas/general-configs/testing

# service config
go run ./app/nodeid-service/cmd/store-configuration/... \
  -conf=./app/nodeid-service/configs \
  -source_dir ./devops/docker-deploy/service-configs \
  -store_dir go-micro-saas/nodeid-service/testing/latest
  
# run
docker run -it --rm \
  --entrypoint /bin/bash \
  -v $PWD/devops/docker-deploy/configs:/myworkspace/golang/src/workspace/configs \
  nodeid-service:latest
```