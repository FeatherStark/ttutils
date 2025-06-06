package tlogger

import (
	"fmt"
	"github.com/FeatherStark/ttutils/ttime"
	"github.com/logrusorgru/aurora"
)

// LogInfoPrintln 输出 Info 信息
func LogInfoPrintln(msg string) {
	logPrint(msg, "INFO")
}

// LogWarningPrintln 输出 Warning 信息
func LogWarningPrintln(msg string) {
	logPrint(msg, "WARN")
}

// LogErrorPrintln 输出 Error 信息
func LogErrorPrintln(msg string) {
	logPrint(msg, "ERROR")
}

// LogDebugPrintln 输出 Debug 信息
func LogDebugPrintln(msg string) {
	logPrint(msg, "DEBUG")
}

func logPrint(msg, level string) {
	if level == "INFO" {
		fmt.Println(aurora.Green("[+] " + ttime.GetNowTime() + " " + level + " " + msg))
	} else if level == "WARN" {
		fmt.Println(aurora.Yellow("[#] " + ttime.GetNowTime() + " " + level + " " + msg))
	} else if level == "ERROR" {
		fmt.Println(aurora.Red("[-] " + ttime.GetNowTime() + " " + level + " " + msg))
	} else if level == "DEBUG" {
		fmt.Println(aurora.Blue("[~] " + ttime.GetNowTime() + " " + level + " " + msg))
	}
}
