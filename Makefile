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
	cd && go get golang.org/x/lint/golint
