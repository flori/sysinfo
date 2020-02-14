GLOBAL_GOPATH := $(HOME)/go/bin
GOPATH := $(shell pwd)/gospace

.EXPORT_ALL_VARIABLES:

.PHONY: build build-info

all: fetch sysinfo

sysinfo: cmd/sysinfo/main.go *.go
	go build -o $@ $<

fetch: fake-package

fake-package:
	rm -rf $(GOPATH)/src/github.com/flori/sysinfo
	mkdir -p $(GOPATH)/src/github.com/flori
	ln -s $(shell pwd) $(GOPATH)/src/github.com/flori/sysinfo

test:
	@go test

coverage:
	@go test -coverprofile=coverage.out

coverage-display: coverage
	@go tool cover -html=coverage.out

clean:
	@rm -f sysinfo coverage.out tags

clobber: clean
	@rm -rf $(GOPATH)/*

tags: clean
	@gotags -tag-relative=false -silent=true -R=true -f $@ . $(GOPATH)

install: all
	@cp -f sysinfo $(GLOBAL_GOPATH)/
