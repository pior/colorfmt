package colorfmt

type tagProcessor func(string) string

type colorizeTagProcessor struct {
	specialColors map[string]string
}

func newColorizeTagProcessor() *colorizeTagProcessor {
	return &colorizeTagProcessor{
		specialColors: map[string]string{
			"link": "green+ubh",
		},
	}
}

func (c *colorizeTagProcessor) process(text string) string {
	if translated, ok := c.specialColors[text]; ok {
		text = translated
	}
	return colorCode(text).String()
}

func ignoreTagProcessor(_ string) string {
	return ""
}
