package iteration

import "strings"

const repeatedCount = 5

func Repeat(char string) string {
	var repeated strings.Builder

	for i := 0; i < repeatedCount; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}