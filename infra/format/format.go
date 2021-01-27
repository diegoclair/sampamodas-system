package format

import "strings"

//FirstLetterUpperCase - returns upper case of all first letter of the words
func FirstLetterUpperCase(str *string) {
	*str = strings.Title(strings.ToLower(*str))
}

//ToUpperCase - transform all string to upper case
func ToUpperCase(str *string) {
	*str = strings.ToUpper(*str)
}
