# test service
TEST_V1_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/test-service/v1 -name "*.proto")
#TEST_V1_INTERNAL_CONFIG_PROTO=$(shell cd $(PROJECT_PATH) && find app/test-service/internal/conf -name "*.proto")
TEST_V1_INTERNAL_CONFIG_PROTO=""
TEST_V1_PROTO_FILES=""
ifneq ($(TEST_V1_INTERNAL_CONFIG_PROTO), "")
	TEST_V1_PROTO_FILES=$(TEST_V1_API_PROTO) $(TEST_V1_INTERNAL_CONFIG_PROTO)
else
	TEST_V1_PROTO_FILES=$(TEST_V1_API_PROTO)
endif
.PHONY: protoc-test-v1
# protoc :-->: generate test v1 service protobuf
protoc-test-v1:
	@echo "# generate test-service protobuf"
	$(call protoc_protobuf,$(TEST_V1_PROTO_FILES))
