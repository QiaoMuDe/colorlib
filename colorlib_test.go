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
}
