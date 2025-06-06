package tcrypt

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5HashEncrypt md5加密字符串。
// Args: text string 需要加密的文本。
// Returns: string 加密后的文本。
func Md5HashEncrypt(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Md5HashEncryptFile 计算文件的 md5 值。
// Args: filepath string 文件路径。
// Returns: string 文件的 md5 值, error 错误信息。
func Md5HashEncryptFile(filepath string) (string, error) {
	// 打开文件
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 创建 MD5 哈希器
	hash := md5.New()
	// 将文件内容复制到哈希器中
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// 计算 MD5 值并转换为十六进制字符串
	md5Sum := hash.Sum(nil)
	md5Hex := hex.EncodeToString(md5Sum)
	return md5Hex, nil
}
