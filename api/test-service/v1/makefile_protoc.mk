override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

SAAS_TEST_V1_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_TEST_V1_INTERNAL_PROTO := "app/test-service/internal/conf/config.conf.proto"
SAAS_TEST_V1_PROTO_FILES := ""
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
