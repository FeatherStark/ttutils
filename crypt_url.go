package ttutils

import (
	"net/url"
	"strconv"
	"strings"
	"unicode"
)

// UrlEncodeNonAlphanumeric 对字符串中除了数字和字母以外的其他字符进行URL编码
// Args: input - 要进行URL编码的字符串
// Returns: string 编码后的字符串
// e.g.: "Hello, World! 123" -> "Hello%2C+World%21+123"
func UrlEncodeNonAlphanumeric(input string) string {
	var encodedString string
	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			// 如果是字母或数字，则直接保留
			encodedString += string(char)
		} else {
			// 否则对字符进行URL编码
			encodedString += url.QueryEscape(string(char))
		}
	}
	return encodedString
}

// UrlEncodeAllCharacters 对字符串中的所有字符（包括字母和数字）都进行URL编码
// Args: input - 要进行URL编码的字符串
// Returns: string 编码后的字符串
// e.g.: "Hello, World! 123" -> "%48%65%6c%6c%6f%2c%20%57%6f%72%6c%64%21%20%31%32%33"
func UrlEncodeAllCharacters(input string) string {
	var builder strings.Builder
	for _, char := range input {
		// 将每个字符转换为对应的 ASCII 十六进制编码
		builder.WriteString("%")
		builder.WriteString(strconv.FormatInt(int64(char), 16))
	}

	return builder.String()
}

// UrlDecodeString 对字符串进行URL解码
// Args: str - 要进行URL解码的字符串
// Returns: string 解码后的字符串
// e.g.: "%48%65%6c%6c%6f%2c%20%57%6f%72%6c%64%21%20%31%32%33" -> "Hello, World! 123"
func UrlDecodeString(str string) string {
	decoded, err := url.QueryUnescape(str)
	if err != nil {
		return str
	}
	return decoded
}
