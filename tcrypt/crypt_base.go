package tcrypt

import "encoding/base64"

// Base64EncodeString 将输入的字符串编码为 Base64 字符串
// Args: str string 明文字符串。
// Returns: string 编码后的 Base64 字符串。
func Base64EncodeString(str string) string {
	// 将字符串转换为字节数组
	data := []byte(str)
	return base64.StdEncoding.EncodeToString(data)
}

// Base64DecodeString 将输入的 Base64 字符串解码为明文字符串
// Args: input string Base64 字符串。
// Returns: string 解码后的明文字符串, error 错误信息。
func Base64DecodeString(input string) (string, error) {
	// 使用 base64.StdEncoding.DecodeString 进行解码
	data, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		// 如果解码失败，返回错误信息
		return "", err
	}
	// 将解码后的字节数组转换为字符串
	return string(data), nil
}
