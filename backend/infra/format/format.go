package format

import "strings"

func FirstLetterUppercase(str *string) {
	*str = strings.Title(strings.ToLower(*str))
}
