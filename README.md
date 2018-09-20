# zshhist-go
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
	ys, err := zshhist.ReadAll(f)
	if err != nil {
		panic(err)
	}
	for _, y := range ys {
		fmt.Printf("%+v\n", y)
	}
}
```

## Bug Tracker
Report bug at [Issues・itchyny/zshhist-go - GitHub](https://github.com/itchyny/zshhist-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
