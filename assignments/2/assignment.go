package assignment

import (
	"sort"
	"unicode"
)

// Reverse reverses the order of a slice of strings.
func Reverse(s []string) []string {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Palindrome returns true if the slice of strings is a palindrome.
func Palindrome(s []string) bool {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// Anagram returns true if the two slices of strings are anagrams.
func Anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Runes are like chars in Go but unicode
	r1 := []rune(s1)
	r2 := []rune(s2)

	sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
	sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

	// Compare sorted slices
	for i := range r1 {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}

// RemoveDigits removes all digit characters from a string
func RemoveDigits(s string) string {
	var result []rune
	for _, r := range s {
		if !unicode.IsDigit(r) {
			result = append(result, r)
		}
	}
	return string(result)
}

// ReplaceDigits replaces all digit characters in a string with the given replacement character.
func ReplaceDigits(s string, r string) string {
	if len(r) != 1 {
		return s
	}

	replacementRune := []rune(r)[0] // Convert the replacement string to a rune
	var result []rune

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			result = append(result, replacementRune) // Replace digit
		} else {
			result = append(result, ch) // Keep non-digit
		}
	}

	return string(result)
}
