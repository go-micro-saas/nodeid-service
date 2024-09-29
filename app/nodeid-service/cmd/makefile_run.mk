.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/nodeid-service/cmd/store-configuration/... -conf=./app/nodeid-service/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/nodeid-service/cmd/database-migration/... -conf=./app/nodeid-service/configs

.PHONY: run-nodeid-service
# run service :-->: run nodeid-service
run-nodeid-service:
	go run ./app/nodeid-service/cmd/nodeid-service/... -conf=./app/nodeid-service/configs

.PHONY: run-service
# run service :-->: run nodeid-service
run-service:
	#@$(MAKE) run-nodeid-service
	go run ./app/nodeid-service/cmd/nodeid-service/... -conf=./app/nodeid-service/configs

.PHONY: testing-service
# testing service :-->: testing ping-service
testing-service:
	go run testdata/get-node-id/main.go
