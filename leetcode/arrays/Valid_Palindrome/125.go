package validpalindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for i, j := 0, len(s)-1; i < j; {

		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--

	}
	return true
}

func isAlphanumeric(c byte) bool {
	return ('a' <= c && c <= 'z') || ('0' <= c && c <= '9')
}
