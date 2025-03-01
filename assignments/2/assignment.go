package assignment

import "sort"

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

	// Runes are like chars in Go
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
