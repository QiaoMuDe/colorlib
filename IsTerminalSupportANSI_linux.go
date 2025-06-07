package colorlib

import (
	"os"

	"golang.org/x/sys/unix"
)

// isTerminalSupportANSI 检测Linux环境下终端是否支持ANSI转义字符
func isTerminalSupportANSI() bool {
	// 检查标准输出是否连接到终端
	if !isatty(int(os.Stdout.Fd())) {
		return false
	}

	// 检查TERM环境变量
	term := os.Getenv("TERM")
	return term != "" && term != "dumb"
}

// isatty 检查文件描述符是否为终端
func isatty(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	return err == nil
}
