package core

import "strings"

// jsCaptureGroups counts capturing groups in a JavaScript regex body.
//
// JavaScript's own trick is to append an empty alternative and read the match
// length, which Go's RE2 cannot reproduce. So the body is scanned directly: a
// '(' opens a group unless it starts a non-capturing construct, and a '(' inside
// a character class or preceded by a backslash never opens one.
func jsCaptureGroups(pattern string) int {
	count := 0
	inClass := false

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]

		if c == '\\' {
			i++ // skip the escaped byte
			continue
		}

		if inClass {
			if c == ']' {
				inClass = false
			}
			continue
		}

		switch c {
		case '[':
			inClass = true
			// A ']' in the first position is a literal, not the class close.
			if i+1 < len(pattern) && pattern[i+1] == '^' {
				i++
			}
			if i+1 < len(pattern) && pattern[i+1] == ']' {
				i++
			}
		case '(':
			if isCapturingGroup(pattern[i:]) {
				count++
			}
		}
	}
	return count
}

// isCapturingGroup reports whether the '(' at the start of s opens a capturing
// group. Named groups in the (?<name>...) form capture; the (?<= and (?<!
// lookbehinds do not, which is why the character following '?' decides.
func isCapturingGroup(s string) bool {
	if !strings.HasPrefix(s, "(") {
		return false
	}
	rest := s[1:]
	if rest == "" || rest[0] != '?' {
		// A plain group captures.
		return true
	}
	// Every construct that begins with (? is non-capturing except a named group.
	if len(rest) >= 2 && (rest[1] == '=' || rest[1] == '!') {
		return false
	}
	if strings.HasPrefix(rest, "<=") || strings.HasPrefix(rest, "<!") {
		return false
	}
	// (?<name>...) is a named capturing group in JavaScript.
	return true
}
