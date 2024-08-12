ABSOLUTE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))

# config
SAAS_TEST_API_PROTO=$(shell cd $(PROJECT_PATH) && find api/test-service -name "*.proto")
#SAAS_TEST_INTERNAL_PROTO=$(shell cd $(PROJECT_PATH) && find app/config/internal/conf -name "*.proto")
SAAS_TEST_INTERNAL_PROTO=
SAAS_TEST_PROTO_FILES=""
ifneq ($(SAAS_TEST_INTERNAL_PROTO), "")
	SAAS_TEST_PROTO_FILES=$(SAAS_TEST_API_PROTO) $(SAAS_TEST_INTERNAL_PROTO)
else
	SAAS_TEST_PROTO_FILES=$(SAAS_TEST_API_PROTO)
endif
.PHONY: protoc-test-protobuf
# protoc :-->: generate test service protobuf
protoc-test-protobuf:
	@echo "# generate test service protobuf"
	echo "ABSOLUTE_PATH $(ABSOLUTE_PATH)"
	$(call protoc_protobuf,$(SAAS_TEST_PROTO_FILES))
