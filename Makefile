GO_DIRS := $(wildcard ./go/day*)
GO_SRC := $(wildcard go/day*/main.go)
GO_BIN := $(patsubst ./go/%, ./build/go/%, $(GO_DIRS))


all: go
.PHONY: all

build/go:
	mkdir -p build/go

build/go/%: build/go go/%/main.go
	go build -o $@ ./go/$*

go-run: $(GO_BIN)
	for d in $(GO_DIRS); do go run ./$$d ; done
.PHONY: go-run

go-check: $(GO_SRC)
	go test $(GO_DIRS)
.PHONY: go-check

go: go-check go-run
.PHONY: go

clean:
	rm -rf build
.PHONY: clean

check: go-check
.PHONY: check
