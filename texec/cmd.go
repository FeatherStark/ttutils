package texec

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os/exec"
	"runtime"
	"strings"
)

// CommandExecResult 命令执行结果
type CommandExecResult struct {
	Success       bool   // Success 命令是否执行成功
	CommandString string // CommandString 命令字符串
	Output        string // Output 命令执行输出
	Error         error  // Error 命令执行错误
}

// CommandExecute 执行系统命令并打印结果
// Args: command string 命令字符串
// Returns: CommandExecResult 命令执行结果
func CommandExecute(command string) CommandExecResult {
	var result CommandExecResult
	result.CommandString = command
	var cmd *exec.Cmd
	args := strings.Split(command, ",")
	switch os := runtime.GOOS; os {
	case "windows":
		cmd = exec.Command("cmd", append([]string{"/c"}, args...)...)
	case "linux", "darwin":
		cmd = exec.Command("/bin/sh", append([]string{"-c"}, args...)...)
	default:
		result.Success = false
		result.Error = fmt.Errorf("不支持的操作系统: %s", os)
		return result
	}
	output, err := cmd.CombinedOutput()
	result.Output = decodeCommandOutput(output)
	if err != nil {
		result.Error = err
		return result
	}
	result.Success = true
	return result
}

// decodeCommandOutput 根据操作系统解码输出内容
func decodeCommandOutput(output []byte) string {
	switch os := runtime.GOOS; os {
	case "windows":
		// Windows 默认使用 GBK 编码，尝试将其转换为 UTF-8
		decoder := simplifiedchinese.GBK.NewDecoder()
		reader := transform.NewReader(bytes.NewReader(output), decoder)
		decoded, err := ioutil.ReadAll(reader)
		if err != nil {
			return string(output)
		}
		return string(decoded)
	case "linux", "darwin":
		// Linux 和 macOS 默认使用 UTF-8 编码，直接返回字符串
		return string(output)
	default:
		return string(output)
	}
}
