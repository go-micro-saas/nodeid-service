override ABSOLUTE_MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
override ABSOLUTE_PATH := $(patsubst %/,%,$(dir $(ABSOLUTE_MAKEFILE)))
override REL_PROJECT_PATH := $(subst $(PROJECT_ABS_PATH)/,,$(ABSOLUTE_PATH))

# config
SAAS_NODEID_API_PROTO := $(shell find ./$(REL_PROJECT_PATH) -name "*.proto")
SAAS_NODEID_INTERNAL_PROTO := ""
SAAS_NODEID_PROTO_FILES := ""
ifneq ($(SAAS_NODEID_INTERNAL_PROTO), "")
	SAAS_NODEID_PROTO_FILES=$(SAAS_NODEID_API_PROTO) $(SAAS_NODEID_INTERNAL_PROTO)
else
	SAAS_NODEID_PROTO_FILES=$(SAAS_NODEID_API_PROTO)
endif
.PHONY: protoc-test-protobuf
# protoc :-->: generate test service protobuf
protoc-test-protobuf:
	@echo "# generate test service protobuf"
	$(call protoc_protobuf,$(SAAS_NODEID_PROTO_FILES))