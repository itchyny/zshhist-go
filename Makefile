.PHONY: all
all: lint test

.PHONY: test
test: testdeps
	go test -v ./...

.PHONY: testdeps
testdeps:
	go get -d -v -t ./...

.PHONY: lint
lint: lintdeps
	golint -set_exit_status *.go

.PHONY: lintdeps
lintdeps:
	command -v golint >/dev/null || go get -u golang.org/x/lint/golint
