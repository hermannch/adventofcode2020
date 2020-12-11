GO_DIRS := $(wildcard ./go/day*)
GO_SRC := $(wildcard go/day*/main.go)
GO_BIN := $(patsubst ./go/%, ./build/go/%, $(GO_DIRS))


all: go
.PHONY: all

build/go/%: go/%/main.go
	@mkdir -p build/go
	ln -s ../../input build/go
	go build -o $@ ./go/$*

go-run: $(GO_BIN)
	for d in $(GO_DIRS); do go run ./$$d ; done
.PHONY: go-run

go-check: $(GO_SRC)
	go test $(GO_DIRS)
.PHONY: go-check

go: $(GO_BIN)
.PHONY: go

clean:
	rm -rf build
.PHONY: clean

check: go-check
.PHONY: check
