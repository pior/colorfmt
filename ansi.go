/*
Code extracted from https://github.com/mgutz/ansi

The MIT License (MIT)
Copyright (c) 2013 Mario L. Gutierrez
*/
package colorfmt

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
	defaultt = 9

	normalIntensityFG = 30
	highIntensityFG   = 90
	normalIntensityBG = 40
	highIntensityBG   = 100

	start         = "\033[0;"
	bold          = "1;"
	blink         = "5;"
	underline     = "4;"
	inverse       = "7;"
	strikethrough = "9;"

	reset = "\033[0m"
)

var colors = map[string]int{
	"black":   black,
	"red":     red,
	"green":   green,
	"yellow":  yellow,
	"blue":    blue,
	"magenta": magenta,
	"cyan":    cyan,
	"white":   white,
	"default": defaultt,
}

func init() {
	for i := 0; i < 256; i++ {
		colors[strconv.Itoa(i)] = i
	}
}

func colorCode(style string) *bytes.Buffer {
	buf := bytes.NewBufferString("")
	if style == "" {
		return buf
	}
	if style == "reset" {
		buf.WriteString(reset)
		return buf
	}

	foregroundBackground := strings.Split(style, ":")
	foreground := strings.Split(foregroundBackground[0], "+")
	fgKey := foreground[0]
	fg := colors[fgKey]
	fgStyle := ""
	if len(foreground) > 1 {
		fgStyle = foreground[1]
	}

	bg, bgStyle := "", ""

	if len(foregroundBackground) > 1 {
		background := strings.Split(foregroundBackground[1], "+")
		bg = background[0]
		if len(background) > 1 {
			bgStyle = background[1]
		}
	}

	buf.WriteString(start)
	base := normalIntensityFG
	if len(fgStyle) > 0 {
		if strings.Contains(fgStyle, "b") {
			buf.WriteString(bold)
		}
		if strings.Contains(fgStyle, "B") {
			buf.WriteString(blink)
		}
		if strings.Contains(fgStyle, "u") {
			buf.WriteString(underline)
		}
		if strings.Contains(fgStyle, "i") {
			buf.WriteString(inverse)
		}
		if strings.Contains(fgStyle, "s") {
			buf.WriteString(strikethrough)
		}
		if strings.Contains(fgStyle, "h") {
			base = highIntensityFG
		}
	}

	// if 256-color
	n, err := strconv.Atoi(fgKey)
	if err == nil {
		fmt.Fprintf(buf, "38;5;%d;", n)
	} else {
		fmt.Fprintf(buf, "%d;", base+fg)
	}

	base = normalIntensityBG
	if len(bg) > 0 {
		if strings.Contains(bgStyle, "h") {
			base = highIntensityBG
		}
		// if 256-color
		n, err := strconv.Atoi(bg)
		if err == nil {
			fmt.Fprintf(buf, "48;5;%d;", n)
		} else {
			fmt.Fprintf(buf, "%d;", base+colors[bg])
		}
	}

	// remove last ";"
	buf.Truncate(buf.Len() - 1)
	buf.WriteRune('m')
	return buf
}
