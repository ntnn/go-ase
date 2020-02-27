BINS ?= $(patsubst cmd/%,%,$(wildcard cmd/*))

default: testrun
testrun:
	docker exec -ti -u sybtst $(shell id -u -n) sh +x /sybase/dlv run go run ./cmd/goconntest

build: $(BINS)
$(BINS):
	go build -o $@ ./cmd/$@/

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
