package assignment

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
