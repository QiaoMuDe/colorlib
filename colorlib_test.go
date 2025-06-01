package colorlib

import (
	"fmt"
	"strings"
	"testing"
)

func TestColorLib(t *testing.T) {
	cl := NewColorLib()

	// 测试颜色
	colors := []string{"black", "red", "green", "yellow", "blue", "purple", "cyan", "white", "gray", "lred", "lgreen", "lyellow", "lblue", "lpurple", "lcyan", "lwhite"}
	for _, color := range colors {
		t.Run(color, func(t *testing.T) {
			// 测试打印方法
			cl.printWithColor(color, "这是一条测试消息\n")
			// 测试返回方法
			msg := cl.returnWithColor(color, "这是一条测试消息")
			if msg == "" {
				t.Errorf("returnWithColor(%s) 返回了一个空字符串", color)
			}
		})
	}

	// 测试带占位符的方法
	formattedColors := []struct {
		color  string
		format string
	}{
		{"black", "这是一个 %s 测试消息"},
		{"red", "这是一个 %s 测试消息"},
		{"green", "这是一个 %s 测试消息"},
		{"yellow", "这是一个 %s 测试消息"},
		{"blue", "这是一个 %s 测试消息"},
		{"purple", "这是一个 %s 测试消息"},
		{"cyan", "这是一个 %s 测试消息"},
		{"white", "这是一个 %s 测试消息"},
		{"gray", "这是一个 %s 测试消息"},
		{"lred", "这是一个 %s 测试消息"},
		{"lgreen", "这是一个 %s 测试消息"},
		{"lyellow", "这是一个 %s 测试消息"},
		{"lblue", "这是一个 %s 测试消息"},
		{"lpurple", "这是一个 %s 测试消息"},
		{"lcyan", "这是一个 %s 测试消息"},
		{"lwhite", "这是一个 %s 测试消息"},
	}
	for _, tc := range formattedColors {
		t.Run(tc.color, func(t *testing.T) {
			cl.printWithColor(tc.color, fmt.Sprintf(tc.format, "格式化"), "\n")
			msg := cl.returnWithColor(tc.color, fmt.Sprintf(tc.format, "格式化"))
			if msg == "" {
				t.Errorf("returnWithColor(%s) 返回了一个空字符串", tc.color)
			}
		})
	}

	// 测试不带占位符的方法
	noFormattedColors := []string{"black", "red", "green", "yellow", "blue", "purple", "cyan", "white", "gray", "lred", "lgreen", "lyellow", "lblue", "lpurple", "lcyan", "lwhite"}
	for _, color := range noFormattedColors {
		t.Run(color, func(t *testing.T) {
			cl.printWithColor(color, "这是一条测试消息\n")
			msg := cl.returnWithColor(color, "这是一条测试消息")
			if msg == "" {
				t.Errorf("returnWithColor(%s) 返回了一个空字符串", color)
			}
		})
	}

	// 测试调试方法
	cl.PrintDebug("这是一条调试消息")
	cl.PrintDebugf("这是一条 %s 调试消息\n", "格式化")
	//测试错误方法
	cl.PrintError("这是一条错误消息")
	cl.PrintErrorf("这是一条 %s 错误消息\n", "格式化")
	// 测试信息方法
	cl.PrintInfo("这是一条信息消息")
	cl.PrintInfof("这是一条 %s 信息消息\n", "格式化")
	// 测试成功方法
	cl.PrintSuccess("这是一条成功消息")
	cl.PrintSuccessf("这是一条 %s 成功消息\n", "格式化")
	// 测试警告方法
	cl.PrintWarning("这是一条警告消息")
	cl.PrintWarningf("这是一条 %s 警告消息\n", "格式化")

	// 新加的简洁方法测试
	cl.PrintDbg("这是一条调试消息")
	cl.PrintDbgf("这是一条 %s 调试消息\n", "格式化")
	cl.PrintErr("这是一条错误消息")
	cl.PrintErrf("这是一条 %s 错误消息\n", "格式化")
	cl.PrintInf("这是一条信息消息")
	cl.PrintInff("这是一条 %s 信息消息\n", "格式化")
	cl.PrintOk("这是一条成功消息")
	cl.PrintOkf("这是一条 %s 成功消息\n", "格式化")
	cl.PrintWarn("这是一条警告消息")
	cl.PrintWarnf("这是一条 %s 警告消息\n", "格式化")
}

func TestNoColor(t *testing.T) {
	// 测试NoColor为true时
	cl := NewColorLib()
	cl.NoColor = true

	// 测试所有颜色方法在NoColor=true时的行为
	colors := []string{"black", "red", "green", "yellow", "blue", "purple", "cyan", "white", "gray", "lred", "lgreen", "lyellow", "lblue", "lpurple", "lcyan", "lwhite"}
	for _, color := range colors {
		t.Run(color, func(t *testing.T) {
			// 测试打印方法
			cl.printWithColor(color, "这是一条禁用颜色的消息\n")
			// 测试返回方法
			msg := cl.returnWithColor(color, "这是一条禁用颜色的消息")
			if msg == "" {
				t.Errorf("returnWithColor(%s) 返回了一个空字符串", color)
			}
		})
	}

	// 测试所有快捷方法
	cl.PrintDebug("调试消息(禁用颜色)")
	cl.PrintError("错误消息(禁用颜色)")
	cl.PrintInfo("信息消息(禁用颜色)")
	cl.PrintSuccess("成功消息(禁用颜色)")
	cl.PrintWarning("警告消息(禁用颜色)")

	// 测试所有简洁方法
	cl.PrintDbg("调试消息(禁用颜色)")
	cl.PrintErr("错误消息(禁用颜色)")
	cl.PrintInf("信息消息(禁用颜色)")
	cl.PrintOk("成功消息(禁用颜色)")
	cl.PrintWarn("警告消息(禁用颜色)")

	// 测试NoColor为false时
	cl.NoColor = false
	// 验证颜色功能恢复
	msg := cl.returnWithColor("red", "红色消息")
	if !strings.Contains(msg, "\033") {
		t.Error("NoColor=false时未正确添加颜色代码")
	}
}

// TestNoBold 测试禁用字体加粗的功能
func TestNoBold(t *testing.T) {
	// 创建 ColorLib 实例
	cl := NewColorLib()
	cl.NoBold = true

	// 测试禁用粗体时的输出
	msg := cl.returnWithColor("red", "测试禁用粗体")
	if strings.Contains(msg, "1;") {
		t.Error("NoBold=true 时仍然包含粗体代码")
	}

	// 测试启用粗体时的输出
	cl.NoBold = false
	msg = cl.returnWithColor("red", "测试启用粗体")
	if !strings.Contains(msg, "1;") {
		t.Error("NoBold=false 时未包含粗体代码")
	}

	// 测试颜色
	colors := []string{"black", "red", "green", "yellow", "blue", "purple", "cyan", "white", "gray", "lred", "lgreen", "lyellow", "lblue", "lpurple", "lcyan", "lwhite"}
	for _, color := range colors {
		t.Run(color, func(t *testing.T) {
			// 测试打印方法
			cl.printWithColor(color, "这是一条禁用字体加粗的消息\n")
			// 测试返回方法
			msg := cl.returnWithColor(color, "这是一条禁用字体加粗的消息")
			if msg == "" {
				t.Errorf("returnWithColor(%s) 返回了一个空字符串", color)
			}
		})
	}
}

func TestNoBold2(t *testing.T) {
	cl := NewColorLib()
	cl.NoBold = true

	cl.PrintOk("这是一条禁用字体加粗的消息")
	cl.NoBold = false
	cl.PrintOk("这是一条启用字体加粗的消息")
}
