# deploy

```shell
# general config
go run testdata/configuration/main.go \
  -consul_config consul \
  -source_dir ./devops/docker-deploy/general-configs \
  -store_dir go-micro-saas/general-config/testing

# service config
go run testdata/configuration/main.go \
  -consul_config consul \
  -source_dir ./devops/docker-deploy/service-configs \
  -store_dir go-micro-saas/nodeid-service/testing/v1.0.0
```