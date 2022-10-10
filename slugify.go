package slugify

import (
	"strings"
)

func Slugify(sentence string) string {
	s := map[string]string{
		" ":  "-",
		"'":  "-",
		"\n": "",
	}

	sentence = strings.Trim(sentence, " ")

	for k, t := range s {
		sentence = strings.ReplaceAll(sentence, k, t)
	}

	return strings.ToLower(sentence)
}
