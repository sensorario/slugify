package slugify

import (
	"strings"
)

func Slugify(sentence string) string {
    s := map[string]string{
        " ": "-",
        "'": "-",
        "\n": "",
    }

    for k, t := range s {
        sentence = strings.ReplaceAll(sentence, k, t)
    }

	return sentence
}
