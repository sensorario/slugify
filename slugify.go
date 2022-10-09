package slugify

import (
	"strings"
)

func Slugify(sentence string) string {
	sentence = strings.ReplaceAll(sentence, " ", "-")
	sentence = strings.ReplaceAll(sentence, "'", "-")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	return sentence
}
