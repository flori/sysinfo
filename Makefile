GOPATH := $(shell pwd)/gospace

.EXPORT_ALL_VARIABLES:

.PHONY: build build-info

all: cpuload

cpuload: cmd/cpuload/main.go *.go
	go build -o $@ $<

fetch: fake-package

fake-package:
	rm -rf $(GOPATH)/src/github.com/flori/cpuload
	mkdir -p $(GOPATH)/src/github.com/flori
	ln -s $(shell pwd) $(GOPATH)/src/github.com/flori/cpuload

test:
	@go test

coverage:
	@go test -coverprofile=coverage.out

coverage-display: coverage
	@go tool cover -html=coverage.out

clean:
	@rm -f cpuload coverage.out tags

clobber: clean
	@rm -rf $(GOPATH)/*

tags: clean
	@gotags -tag-relative=false -silent=true -R=true -f $@ . $(GOPATH)
