GO_DIRS := $(wildcard ./go/day*)
GO_SRC := $(wildcard go/day*/main.go)


all: go
.PHONY: all

build/go/%: go/%/main.go
	@mkdir -p build/go
	go build -o $@ ./go/$*

go-check: $(GO_SRC)
	go test $(GO_DIRS)
.PHONY: go-check

go: $(patsubst ./go/%, ./build/go/%, $(GO_DIRS))
.PHONY: go

clean:
	rm -rf build
.PHONY: clean

check: go-check
.PHONY: check
