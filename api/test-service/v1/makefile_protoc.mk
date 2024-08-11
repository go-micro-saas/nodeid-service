# config
SAAS_TEST_V1_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/test-service/v1 -name "*.proto")
#SAAS_TEST_V1_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
SAAS_TEST_V1_INTERNAL_PROTO=
SAAS_TEST_V1_PROTO_FILES=""
ifneq ($(SAAS_TEST_V1_INTERNAL_PROTO), "")
	SAAS_TEST_V1_PROTO_FILES=$(SAAS_TEST_V1_API_PROTO) $(SAAS_TEST_V1_INTERNAL_PROTO)
else
	SAAS_TEST_V1_PROTO_FILES=$(SAAS_TEST_V1_API_PROTO)
endif
.PHONY: protoc-test-v1-protobuf
# protoc :-->: generate test service v1 protobuf
protoc-test-v1-protobuf:
	@echo "# generate test service v1 protobuf"
	$(call protoc_protobuf,$(SAAS_TEST_V1_PROTO_FILES))
