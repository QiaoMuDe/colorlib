package colorlib

import (
	"testing"
)

func TestColorLib(t *testing.T) {
	// 创建 ColorLib 实例
	cl := NewColorLib()

	// 测试各种颜色的打印和返回方法（带占位符）
	cl.Bluef("这是一条带有占位符的蓝色消息: %d", 123)
	cl.Greenf("这是一条带有占位符的绿色消息: %s", "测试")
	cl.Redf("这是一条带有占位符的红色消息: %f", 3.14)
	cl.Yellowf("这是一条带有占位符的黄色消息: %v", []int{1, 2, 3})
	cl.Purplef("这是一条带有占位符的紫色消息: %t", true)

	blueStr := cl.Sbluef("这是一条带有占位符的蓝色字符串: %d", 456)
	greenStr := cl.Sgreenf("这是一条带有占位符的绿色字符串: %s", "测试2")
	redStr := cl.Sredf("这是一条带有占位符的红色字符串: %f", 6.28)
	yellowStr := cl.Syellowf("这是一条带有占位符的黄色字符串: %v", []int{4, 5, 6})
	purpleStr := cl.Spurplef("这是一条带有占位符的紫色字符串: %t", false)

	t.Logf("蓝色字符串: %s", blueStr)
	t.Logf("绿色字符串: %s", greenStr)
	t.Logf("红色字符串: %s", redStr)
	t.Logf("黄色字符串: %s", yellowStr)
	t.Logf("紫色字符串: %s", purpleStr)

	// 测试各种颜色的打印和返回方法（不带占位符）
	cl.Blue("这是一条没有占位符的蓝色消息")
	cl.Green("这是一条没有占位符的绿色消息")
	cl.Red("这是一条没有占位符的红色消息")
	cl.Yellow("这是一条没有占位符的黄色消息")
	cl.Purple("这是一条没有占位符的紫色消息")

	blueStr2 := cl.Sblue("这是一条没有占位符的蓝色字符串")
	greenStr2 := cl.Sgreen("这是一条没有占位符的绿色字符串")
	redStr2 := cl.Sred("这是一条没有占位符的红色字符串")
	yellowStr2 := cl.Syellow("这是一条没有占位符的黄色字符串")
	purpleStr2 := cl.Spurple("这是一条没有占位符的紫色字符串")

	t.Logf("蓝色字符串2: %s", blueStr2)
	t.Logf("绿色字符串2: %s", greenStr2)
	t.Logf("红色字符串2: %s", redStr2)
	t.Logf("黄色字符串2: %s", yellowStr2)
	t.Logf("紫色字符串2: %s", purpleStr2)

	// 测试 PromptMsg 方法
	cl.PromptMsg("成功", "绿色", "这是一条成功消息")
	cl.PromptMsg("错误", "红色", "这是一条错误消息")
	cl.PromptMsg("警告", "黄色", "这是一条警告消息")
	cl.PromptMsg("信息", "蓝色", "这是一条信息消息")

	// 测试 PrintSuccessf、PrintErrorf、PrintWarningf、PrintInfof 方法（带占位符）
	cl.PrintSuccessf("这是一条带有占位符的成功消息: %d", 789)
	cl.PrintErrorf("这是一条带有占位符的错误消息: %s", "测试3")
	cl.PrintWarningf("这是一条带有占位符的警告消息: %f", 9.87)
	cl.PrintInfof("这是一条带有占位符的信息消息: %v", []int{7, 8, 9})

	// 测试 PrintSuccess、PrintError、PrintWarning、PrintInfo 方法（不带占位符）
	cl.PrintSuccess("这是一条没有占位符的成功消息")
	cl.PrintError("这是一条没有占位符的错误消息")
	cl.PrintWarning("这是一条没有占位符的警告消息")
	cl.PrintInfo("这是一条没有占位符的信息消息")
}
