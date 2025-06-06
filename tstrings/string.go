package tstrings

import (
	"unicode"
)

// IsAlphaNumeric 检测字符串只包含字母和数字，如果是则返回true。
// Args: str string 需要检测的字符串。
// Returns: bool true表示只包含字母和数字，false表示不包含字母和数字。
func IsAlphaNumeric(str string) bool {
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
