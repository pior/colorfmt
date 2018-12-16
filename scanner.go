package colorfmt

const openingChar, closingChar = '{', '}'

func scan(text string, processor tagProcessor) string {
	var in bool
	var tag string
	var output string
	var colorized bool

	for _, char := range text {
		if in {
			if char == '{' && tag == "" { // "{{"
				in = false
				output += string(char)
			} else if char == closingChar {
				colorCode := processor(tag)
				if colorCode != "" {
					output += colorCode
					colorized = true
				}
				in = false
				tag = ""
			} else {
				tag += string(char)
			}
		} else {
			if char == openingChar {
				in = true
			} else {
				output += string(char)
			}
		}
	}

	resetCode := processor("reset")
	if colorized {
		output += resetCode
	}

	return resetCode + output
}
