package ttutils

import (
	"fmt"
	"os"
	"path/filepath"
)

// FileOverWrite 如果文件已经存在，内容会被覆盖。
// Args: filePath string 文件路径, data string 文件内容
// Returns: error 错误信息。
func FileOverWrite(filePath string, data string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(data)
	return err
}

// FileAppendData 将字符串追加到文件尾，如果文件不存在则创建该文件。
// Args: filePath string 文件路径, data string 要追加的数据,。
// Return: error 错误信息。
func FileAppendData(filePath string, data string) error {
	// 打开文件，如果文件不存在则创建文件，如果文件存在则追加内容
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// 写入字符串到文件
	_, err = file.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

// FileRead 读取指定路径文件内容
// Args: filePath string 文件路径。
// Returns: string 文件内容, error 错误信息。
func FileRead(filePath string) (string, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// 读取文件内容
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), err
}

// FileDelete 删除文件
// Args: filePath string 文件路径。
// Return: error 错误信息。
func FileDelete(filePath string) error {
	return os.Remove(filePath)
}

// FileEnsureDirAndFileExists 确保目录和文件存在，如果文件不存在则创建目录和文件，同时写入文件内容。
// Args: filePath string 文件路径, data string 文件内容。
// Return: error 错误信息。
func FileEnsureDirAndFileExists(filePath string, data string) error {
	// 获取文件所在的目录路径
	dir := filepath.Dir(filePath)

	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm) // 递归创建目录
		if err != nil {
			return fmt.Errorf("无法创建目录: %w", err)
		}
	}

	// 检查文件是否存在，如果不存在则创建并写入数据
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath) // 创建文件
		if err != nil {
			return fmt.Errorf("无法创建文件: %w", err)
		}
		defer file.Close() // 确保函数结束时关闭文件

		// 写入数据到文件
		_, err = file.WriteString(data)
		if err != nil {
			return fmt.Errorf("无法写入数据到文件: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("检查文件时出错: %w", err)
	}
	return nil
}
