.PHONY: run-uuid-service
# run service :-->: run uuid-service
run-uuid-service:
	go run ./app/uuid-service/cmd/uuid-service/... -conf=./app/uuid-service/configs
