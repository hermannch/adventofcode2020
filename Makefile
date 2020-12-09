GO_DIRS := $(wildcard ./go/day*)
GO_SRC := $(wildcard go/day*/main.go)


all: go
.PHONY: all

build/go/%: go/%
	@mkdir -p build/go
	go build -o $@ ./go/$*

go: $(patsubst ./go/%, ./build/go/%, $(GO_DIRS))
.PHONY: go

clean:
	rm -rf build
.PHONY: clean

check:
	go test $(GO_DIRS)
.PHONY: check
