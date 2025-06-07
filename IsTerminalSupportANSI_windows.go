package colorlib

import (
	"os"
	"syscall"
)

// 检测Windows终端是否支持ANSI转义字符
func isTerminalSupportANSI() bool {
	// 获取标准输出的句柄
	stdout := syscall.Handle(os.Stdout.Fd())

	// 定义变量用于存储控制台模式
	var mode uint32

	// 调用GetConsoleMode获取控制台模式
	err := syscall.GetConsoleMode(stdout, &mode)
	if err != nil {
		// 如果获取失败，通常意味着不是控制台环境
		return false
	}

	// 检查是否启用了虚拟终端处理（ANSI转义字符支持）
	return mode&0x0004 != 0 // ENABLE_VIRTUAL_TERMINAL_PROCESSING
}
