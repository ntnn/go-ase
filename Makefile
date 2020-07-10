BINS ?= $(patsubst cmd/%,%,$(wildcard cmd/*))

build: $(BINS)
$(BINS):
	go build -o $@ ./cmd/$@/

generate:
ifeq (x$(TARGET),x)
	grep -r '^// Code generated by ".*"\; DO NOT EDIT.$\' ./ | awk -F: '{ print $$1 }' | xargs rm
	go generate ./...
else
	grep '^// Code generated by ".*"\; DO NOT EDIT.$\' ./$(TARGET)/* | awk -F: '{ print $$1 }' | xargs rm
	go generate ./$(TARGET)
endif

test:
	go test -vet all -cover ./cgo/... ./go/... ./cmd/... ./libase/...

integration: integration-cgo
integration-cgo:
	go test ./tests/cgotest
	go test ./examples/cgo/...

GO_EXAMPLES := $(wildcard examples/go/*)
CGO_EXAMPLES := $(wildcard examples/cgo/*)
EXAMPLES := $(GO_EXAMPLES) $(CGO_EXAMPLES)

examples: $(EXAMPLES)

.PHONY: $(EXAMPLES)
$(EXAMPLES):
	@echo Running example: $@
	go run ./$@/main.go
