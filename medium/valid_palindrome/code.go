package valid_palindrome

import (
	"unicode"
)

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !isAlphanumeric(rune(s[i])) {
			i++
			continue
		}
		if !isAlphanumeric(rune(s[j])) {
			j--
			continue
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphanumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
