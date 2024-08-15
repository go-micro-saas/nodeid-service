# ping-service
.PHONY: run-database-migration
# run service :-->: run database-migration
run-database-migration:
	go run ./app/nodeid-service/cmd/database-migration/... -conf=./app/nodeid-service/configs
