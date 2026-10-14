package core

import (
	"io"
	"os"
)

func openFile(path string) (*os.File, error) {
	return os.Open(path)
}

func containsAny(haystack []byte, needles []string) bool {
	if len(haystack) == 0 {
		return false
	}
	for _, n := range needles {
		if indexBytes(haystack, []byte(n)) >= 0 {
			return true
		}
	}
	return false
}

func indexBytes(haystack, needle []byte) int {
	n := len(needle)
	if n == 0 || len(haystack) < n {
		return -1
	}
	first := needle[0]
	for i := 0; i+n <= len(haystack); i++ {
		if haystack[i] != first {
			continue
		}
		if string(haystack[i:i+n]) == string(needle) {
			return i
		}
	}
	return -1
}

// replaceFirst replaces the first occurrence of old with new, returning the
// result and how many replacements were made.
func replaceFirst(s, old, new string) (string, int) {
	if old == "" {
		return s, 0
	}
	i := indexBytes([]byte(s), []byte(old))
	if i < 0 {
		return s, 0
	}
	return s[:i] + new + s[i+len(old):], 1
}

var _ io.Reader = (*os.File)(nil)
