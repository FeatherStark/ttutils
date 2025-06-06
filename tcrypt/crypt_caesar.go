package tcrypt

import "unicode"

// CaesarEncrypt 凯撒编码。
// Args: plainText string 明文, shift int 加密偏移量。
// Returns: string 密文。
func CaesarEncrypt(plainText string, shift int) string {
	var result string
	for _, char := range plainText {
		if unicode.IsLetter(char) { // 如果是字母
			offset := rune('a')
			if unicode.IsUpper(char) { // 大写字母
				offset = rune('A')
			}
			// 计算加密后的字母
			encryptedChar := (char-offset+rune(shift))%26 + offset
			result += string(encryptedChar)
		} else {
			// 非字母字符直接添加
			result += string(char)
		}
	}
	return result
}

// CaesarDecrypt 凯撒解码。
// Args: encryptedText string 密文, shift int 解密偏移量。
// Returns: string 明文。
func CaesarDecrypt(encryptedText string, shift int) string {
	var result string
	for _, char := range encryptedText {
		if unicode.IsLetter(char) { // 如果是字母
			offset := rune('a')
			if unicode.IsUpper(char) { // 大写字母
				offset = rune('A')
			}
			// 计算解密后的字母
			decryptedChar := (char-offset-rune(shift)+26)%26 + offset
			result += string(decryptedChar)
		} else {
			// 非字母字符直接添加
			result += string(char)
		}
	}
	return result
}
