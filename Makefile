GLOBAL_GOPATH := $(HOME)/go/bin
GOPATH := $(shell pwd)/gospace

.EXPORT_ALL_VARIABLES:

.PHONY: build build-info

check-%:
	@if [ "${${*}}" = "" ]; then \
		echo >&2 "Environment variable $* not set"; \
		exit 1; \
	fi

all: sysinfo

sysinfo: cmd/sysinfo/main.go *.go
	go build -o $@ $<

setup: fake-package
	go mod download

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

validate-tag:
	@if ! echo "${TAG}" | grep -qE '^v[0-9]+\.[0-9]+\.[0-9]+$$'; then \
		echo >&2 "Error: TAG must be in the format 'v1.2.3'"; \
		exit 1; \
	fi # '

release: check-TAG validate-tag
	git push origin master
	git tag "$(TAG)"
	git push origin "$(TAG)"

tags: clean
	@gotags -tag-relative=false -silent=true -R=true -f $@ . $(GOPATH)

install: all
	@cp -f sysinfo $(GLOBAL_GOPATH)/
