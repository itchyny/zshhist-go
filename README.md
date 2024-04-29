# zshhist-go
[![CI Status](https://github.com/itchyny/zshhist-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/itchyny/zshhist-go/actions?query=branch:main)
[![Go Report Card](https://goreportcard.com/badge/github.com/itchyny/zshhist-go)](https://goreportcard.com/report/github.com/itchyny/zshhist-go)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/itchyny/zshhist-go/blob/main/LICENSE)
[![pkg.go.dev](https://pkg.go.dev/badge/github.com/itchyny/zshhist-go)](https://pkg.go.dev/github.com/itchyny/zshhist-go)

zsh histfile utility for Go

## Usage
```go
package main

import (
	"fmt"
	"os"

	"github.com/itchyny/zshhist-go"
	"github.com/mitchellh/go-homedir"
)

func main() {
	path, err := homedir.Expand("~/.histfile")
	if err != nil {
		panic(err)
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := zshhist.NewReader(f)
	for r.Scan() {
		fmt.Printf("%+v\n", r.History())
	}
	if r.Err() != nil {
		panic(r.Err())
	}
}
```

## Bug Tracker
Report bug at [Issues・itchyny/zshhist-go - GitHub](https://github.com/itchyny/zshhist-go/issues).

## Author
itchyny (<https://github.com/itchyny>)

## License
This software is released under the MIT License, see LICENSE.
