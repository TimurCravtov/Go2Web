package html

import "regexp"

func colorizeURLs(text string) string {

	const colorBlue = "\033[94m"
	const colorReset = "\033[0m"

	reURL := regexp.MustCompile(`https?://[^\s)]+`)

	coloredText := reURL.ReplaceAllStringFunc(text, func(match string) string {
		return colorBlue + match + colorReset
	})
	return coloredText
}