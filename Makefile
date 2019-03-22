export GO111MODULE=on

.PHONY: all
all: test lint

.PHONY: deps
deps:
	go get -d -v ./...

.PHONY: test
test: deps
	go test -v ./...

.PHONY: lint
lint: lintdeps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: lintdeps
lintdeps:
	GO111MODULE=off go get -u golang.org/x/lint/golint
