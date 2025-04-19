package colorlib

import (
	"fmt"
	"testing"
)

func TestColorLib(t *testing.T) {
	cl := NewColorLib()

	// 测试颜色
	colors := []string{"black", "red", "green", "yellow", "blue", "purple", "cyan", "white", "gray", "lred", "lgreen", "lyellow", "lblue", "lpurple", "lcyan", "lwhite"}
	for _, color := range colors {
		t.Run(color, func(t *testing.T) {
			// 测试打印方法
			cl.printWithColor(color, "这是一条测试消息")
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
			cl.printWithColor(tc.color, fmt.Sprintf(tc.format, "格式化"))
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
			cl.printWithColor(color, "这是一条测试消息")
			msg := cl.returnWithColor(color, "这是一条测试消息")
			if msg == "" {
				t.Errorf("returnWithColor(%s) 返回了一个空字符串", color)
			}
		})
	}

	// 测试调试方法
	cl.PrintDebug("这是一条调试消息")
	cl.PrintDebugf("这是一条 %s 调试消息", "格式化")
	// 测试错误方法
	cl.PrintError("这是一条错误消息")
	cl.PrintErrorf("这是一条 %s 错误消息", "格式化")
	// 测试信息方法
	cl.PrintInfo("这是一条信息消息")
	cl.PrintInfof("这是一条 %s 信息消息", "格式化")
	// 测试成功方法
	cl.PrintSuccess("这是一条成功消息")
	cl.PrintSuccessf("这是一条 %s 成功消息", "格式化")
	// 测试警告方法
	cl.PrintWarning("这是一条警告消息")
	cl.PrintWarningf("这是一条 %s 警告消息", "格式化")

	// 新加的简洁方法测试
	println("简洁方法测试：")
	cl.PrintDbg("这是一条调试消息")
	cl.PrintDbgf("这是一条 %s 调试消息", "格式化")
	cl.PrintErr("这是一条错误消息")
	cl.PrintErrf("这是一条 %s 错误消息", "格式化")
	cl.PrintInf("这是一条信息消息")
	cl.PrintInff("这是一条 %s 信息消息", "格式化")
	cl.PrintOk("这是一条成功消息")
	cl.PrintOkf("这是一条 %s 成功消息", "格式化")
	cl.PrintWarn("这是一条警告消息")
	cl.PrintWarnf("这是一条 %s 警告消息", "格式化")

	// 测试仓库镜像同步
}
