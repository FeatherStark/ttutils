package tcrypt

import (
	"strconv"
	"strings"
)

// UnicodeEncodeString 将字符串编码为 Unicode 转义序列
// Args: input string 要编码的字符串。
// Returns: string 编码后的字符串。
func UnicodeEncodeString(input string) string {
	var result string
	for _, r := range input {
		// 将每个字符转换为 Unicode 转义形式 (\uXXXX 或 \UXXXXXXXX)
		if r <= 0xFFFF {
			// 使用 \uXXXX 格式表示
			result += "\\u" + strconv.FormatInt(int64(r), 16)
		} else {
			// 使用 \UXXXXXXXX 格式表示
			result += "\\U" + strconv.FormatInt(int64(r), 16)
		}
	}
	return result
}

// UnicodeDecodeString 将字符串中的Unicode编码解码为对应的字符。
// Args: str string unicode 编码。
// Returns: string unicode 解码后的字符串。
func UnicodeDecodeString(str string) string {
	var result strings.Builder
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' && i+1 < len(str) && str[i+1] == 'u' {
			// 读取下一个4个十六进制数字
			if i+6 <= len(str) {
				hex := str[i+2 : i+6]
				codePoint, err := strconv.ParseInt(hex, 16, 32)
				if err != nil {
					return str
				}
				result.WriteRune(rune(codePoint))
				i += 5 // 跳过 '\uXXXX'
			} else {
				return str
			}
		} else {
			result.WriteByte(str[i])
		}
	}
	return result.String()
}
