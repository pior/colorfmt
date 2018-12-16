# Colorfmt

Golang ANSI-colors library based on inlined tags.

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

## Tag format

Format: `fgColor+fgAttributes:bgColor+bgAttributes`

Examples:
- `{red}`
- `{yellow:bg}`

Colors:
- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white
- 0...255 (256 colors)

Foreground Attributes:
- B = Blink
- b = bold
- h = high intensity (bright)
- i = inverse
- s = strikethrough
- u = underline

Background Attributes:
- h = high intensity (bright)

Special tags:
- `{reset}`: emit a reset ANSI code to clear all coloring/styling
- `{link}`: simulate the style of a clickable hypertext link


## License

[MIT](https://github.com/pior/colorfmt/blob/master/LICENSE)
