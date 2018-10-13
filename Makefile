all: lint test

test: testdeps
	go test -v ./...

testdeps:
	go get -d -v -t ./...

lint: lintdeps
	golint -set_exit_status *.go

lintdeps:
	command -v golint >/dev/null || go get -u golang.org/x/lint/golint

.PHONY: test testdeps lint lintdeps
