package colorlib

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

// TestGlobalInstance 测试全局实例CL的功能
func TestConcurrencySafety(t *testing.T) {
	fmt.Println("=== 并发安全测试 ===")
	cl := NewColorLib()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cl.PrintInfo("并发测试")
			cl.Sgreen("并发字符串")
		}()
	}

	wg.Wait()
}

// TestColorLib 测试ColorLib的功能
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
	cl.NoColor.Store(true)

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
	cl.NoColor.Store(false)
	// 验证颜色功能恢复
	msg := cl.returnWithColor("red", "红色消息")
	if !strings.Contains(msg, "\033") {
		t.Error("NoColor=false时未正确添加颜色代码")
	}

	// 新增测试：验证禁用颜色时不会输出ANSI转义字符
	cl.NoColor.Store(true)
	msg = cl.returnWithColor("red", "测试禁用颜色输出")
	if strings.Contains(msg, "\033") {
		t.Error("NoColor=true时仍然包含ANSI转义字符")
	}
}

// TestNoBold 测试禁用字体加粗的功能
func TestNoBold(t *testing.T) {
	// 创建 ColorLib 实例
	cl := NewColorLib()
	cl.NoBold.Store(true)

	// 测试禁用粗体时的输出
	msg := cl.returnWithColor("red", "测试禁用粗体")
	if strings.Contains(msg, "1;") {
		t.Error("NoBold=true 时仍然包含粗体代码")
	}

	// 测试启用粗体时的输出
	cl.NoBold.Store(false)
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
	cl.NoBold.Store(true)

	cl.PrintOk("这是一条禁用字体加粗的消息")
	cl.NoBold.Store(false)
	cl.PrintOk("这是一条启用字体加粗的消息")
}

// TestGlobalInstance 测试全局实例的功能
func TestGlobalInstance(t *testing.T) {
	// 测试打印方法
	CL.PrintDebug("这是一条来自全局实例的调试消息")
	CL.PrintError("这是一条来自全局实例的错误消息")
	CL.PrintInfo("这是一条来自全局实例的信息消息")

	// 测试颜色方法
	CL.Blue("这是一条来自全局实例的蓝色消息")
	CL.Red("这是一条来自全局实例的红色消息")

	// 测试返回方法
	msg := CL.Sgreen("这是一条来自全局实例的绿色字符串")
	if msg == "" {
		t.Error("Sgreen方法返回了一个空字符串")
	}
}

// TestUnderlineAndBlink 测试下划线和闪烁属性
func TestUnderlineAndBlink(t *testing.T) {
	cl := NewColorLib()

	// 测试下划线
	cl.Underline.Store(true)
	fmt.Println("=== 下划线效果测试 ===")
	cl.printWithColor("red", "这是一条带下划线的红色消息\n")
	msg := cl.returnWithColor("red", "带下划线的消息")
	if !strings.Contains(msg, "4;") {
		t.Error("未正确添加下划线代码")
	}

	// 测试闪烁
	cl.Blink.Store(true)
	fmt.Println("\n=== 闪烁效果测试 ===")
	cl.printWithColor("blue", "这是一条带闪烁的蓝色消息\n")
	msg = cl.returnWithColor("blue", "带闪烁的消息")
	if !strings.Contains(msg, "5;") {
		t.Error("未正确添加闪烁代码")
	}

	// 测试同时启用下划线和闪烁
	fmt.Println("\n=== 下划线+闪烁效果测试 ===")
	cl.printWithColor("green", "这是一条同时带下划线和闪烁的绿色消息\n")
	msg = cl.returnWithColor("green", "带下划线和闪烁的消息")
	if !strings.Contains(msg, "4;") || !strings.Contains(msg, "5;") {
		t.Error("未正确添加下划线和闪烁代码")
	}

	// 测试禁用下划线和闪烁
	cl.Underline.Store(false)
	cl.Blink.Store(false)
	fmt.Println("\n=== 普通效果测试 ===")
	cl.printWithColor("yellow", "这是一条普通黄色消息\n")
	msg = cl.returnWithColor("yellow", "普通消息")
	if strings.Contains(msg, "4;") || strings.Contains(msg, "5;") {
		t.Error("错误地添加了下划线或闪烁代码")
	}
}

// 测试新增的通用颜色方法
func TestColorLib_NewColorMethods(t *testing.T) {
	fmt.Println("\n=== 测试新增的通用颜色方法 ===")

	cl := NewColorLib()

	cl.PrintColorln(Blue, "这是一条蓝色消息")

	cl.PrintColorf(Blue, "这是一条 %s 蓝色消息\n", "格式化")

	cl.PrintColorln(Red, "这是一条红色消息")

	cl.PrintColorf(Red, "这是一条 %s 红色消息\n", "格式化")

	fmt.Println(cl.SColor(Blue, "这是一条返回的蓝色字符串"))

	fmt.Println(cl.SColorf(Blue, "这是一条 %s 返回的蓝色字符串\n", "格式化"))
}
