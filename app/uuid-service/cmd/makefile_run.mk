.PHONY: run-uuid-service
# run service :-->: run uuid-service
run-uuid-service:
	go run ./app/uuid-service/cmd/uuid-service/... -conf=./app/uuid-service/configs

.PHONY: testing-uuid-service
# testing service :-->: testing uuid-service
testing-uuid-service:
	curl http://127.0.0.1:10201/api/v1/uuid/next-id
