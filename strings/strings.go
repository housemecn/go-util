package strings

import "unicode"

// Capitalize make first letter to upper
// Source https://www.cnblogs.com/Detector/p/9686443.html
func Capitalize(str string) string {
	var upperStr string
	vv := []rune(str)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				upperStr += string(vv[i])
			} else {
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

// IsCapitalize whether the first letter is upper
func IsCapitalize(s string) bool {
	if len(s) < 1 {
		return false
	}
	return unicode.IsUpper([]rune(s)[0])
}

// SplitToChunks Split large string in n-size chunks
func SplitToChunks(s string, chunkSize int) []string {
	if chunkSize < 1 {
		return []string{}
	}

	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks
}
