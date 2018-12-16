# Colorfmt

Golang ANSI-colors library based on tags

## Install

```
go get -u github.com/dolab/colorize
```

## Usage

```go
package main

import (
	"os"

	"github.com/pior/colorfmt"
)

func main() {
	colorfmt.Printf("{red+bh}-> {yellow+h}Go to {link}%s", "https://wikipedia.com")
	colorfmt.Printf(" {white+bh}NOW !!{reset}\n")

	text := "{Red+bh}Warning: {yellow}this some color\n"

	colorfmt.New(os.Stderr, true).Printf(text)  // To stdout
	colorfmt.New(os.Stderr, false).Printf(text) // Without color
}
```
