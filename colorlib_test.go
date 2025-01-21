package colorlib

import (
	"testing"
)

func TestColorLib(t *testing.T) {
	// 创建 ColorLib 实例
	cl := NewColorLib()

	// 测试 Bluef 方法
	cl.Bluef("这是一条带有占位符的蓝色消息: %d", 123)

	// 测试 Greenf 方法
	cl.Greenf("这是一条带有占位符的绿色消息: %s", "测试")

	// 测试 Redf 方法
	cl.Redf("这是一条带有占位符的红色消息")

	// 测试 Yellowf 方法
	cl.Yellowf("这是一条带有占位符的黄色消息")

	// 测试 Purplef 方法
	cl.Purplef("这是一条带有占位符的紫色消息")

	// 测试 Sbluef 方法
	blueStr := cl.Sbluef("这是一条带有占位符的蓝色字符串: %d", 456)
	if blueStr != "\033[1;34m这是一条带有占位符的蓝色字符串: 456\033[0m" {
		t.Errorf("Sbluef 方法返回的字符串不正确")
	}

	// 测试 Sgreenf 方法
	greenStr := cl.Sgreenf("这是一条带有占位符的绿色字符串: %s", "测试")
	if greenStr != "\033[1;32m这是一条带有占位符的绿色字符串: 测试\033[0m" {
		t.Errorf("Sgreenf 方法返回的字符串不正确")
	}

	// 测试 Sredf 方法
	redStr := cl.Sredf("这是一条带有占位符的红色字符串")
	if redStr != "\033[1;31m这是一条带有占位符的红色字符串\033[0m" {
		t.Errorf("Sredf 方法返回的字符串不正确")
	}

	// 测试 Syellowf 方法
	yellowStr := cl.Syellowf("这是一条带有占位符的黄色字符串")
	if yellowStr != "\033[1;33m这是一条带有占位符的黄色字符串\033[0m" {
		t.Errorf("Syellowf 方法返回的字符串不正确")
	}

	// 测试 Spurplef 方法
	purpleStr := cl.Spurplef("这是一条带有占位符的紫色字符串")
	if purpleStr != "\033[1;35m这是一条带有占位符的紫色字符串\033[0m" {
		t.Errorf("Spurplef 方法返回的字符串不正确")
	}

	// 测试 Blue 方法
	cl.Blue("这是一条没有占位符的蓝色消息")

	// 测试 Green 方法
	cl.Green("这是一条没有占位符的绿色消息")

	// 测试 Red 方法
	cl.Red("这是一条没有占位符的红色消息")

	// 测试 Yellow 方法
	cl.Yellow("这是一条没有占位符的黄色消息")

	// 测试 Purple 方法
	cl.Purple("这是一条没有占位符的紫色消息")

	// 测试 Sblue 方法
	blueMsg := cl.Sblue("这是一条没有占位符的蓝色消息")
	if blueMsg != "\033[1;34m这是一条没有占位符的蓝色消息\033[0m" {
		t.Errorf("Sblue 方法返回的字符串不正确")
	}

	// 测试 Sgreen 方法
	greenMsg := cl.Sgreen("这是一条没有占位符的绿色消息")
	if greenMsg != "\033[1;32m这是一条没有占位符的绿色消息\033[0m" {
		t.Errorf("Sgreen 方法返回的字符串不正确")
	}

	// 测试 Sred 方法
	redMsg := cl.Sred("这是一条没有占位符的红色消息")
	if redMsg != "\033[1;31m这是一条没有占位符的红色消息\033[0m" {
		t.Errorf("Sred 方法返回的字符串不正确")
	}

	// 测试 Syellow 方法
	yellowMsg := cl.Syellow("这是一条没有占位符的黄色消息")
	if yellowMsg != "\033[1;33m这是一条没有占位符的黄色消息\033[0m" {
		t.Errorf("Syellow 方法返回的字符串不正确")
	}

	// 测试 Spurple 方法
	purpleMsg := cl.Spurple("这是一条没有占位符的紫色消息")
	if purpleMsg != "\033[1;35m这是一条没有占位符的紫色消息\033[0m" {
		t.Errorf("Spurple 方法返回的字符串不正确")
	}

	// 测试 PrintSuccessf 方法
	cl.PrintSuccessf("这是一条带有占位符的成功消息: %d", 789)

	// 测试 PrintErrorf 方法
	cl.PrintErrorf("这是一条带有占位符的错误消息")

	// 测试 PrintWarningf 方法
	cl.PrintWarningf("这是一条带有占位符的警告消息")

	// 测试 PrintInfof 方法
	cl.PrintInfof("这是一条带有占位符的信息消息")

	// 测试 PrintSuccess 方法
	cl.PrintSuccess("这是一条没有占位符的成功消息")

	// 测试 PrintError 方法
	cl.PrintError("这是一条没有占位符的错误消息")

	// 测试 PrintWarning 方法
	cl.PrintWarning("这是一条没有占位符的警告消息")

	// 测试 PrintInfo 方法
	cl.PrintInfo("这是一条没有占位符的信息消息")

	// 增加更多测试情况
	cl.Bluef("这是一条带有多个占位符的蓝色消息: %d, %s", 123, "测试")
	cl.Greenf("这是一条带有多个占位符的绿色消息: %s, %d", "测试", 456)
	cl.Redf("这是一条带有多个占位符的红色消息: %s, %d", "测试", 789)
	cl.Yellowf("这是一条带有多个占位符的黄色消息: %d, %s", 123, "测试")
	cl.Purplef("这是一条带有多个占位符的紫色消息: %s, %d", "测试", 456)

	cl.Blue("这是一条没有占位符的蓝色消息")
	cl.Green("这是一条没有占位符的绿色消息")
	cl.Red("这是一条没有占位符的红色消息")
	cl.Yellow("这是一条没有占位符的黄色消息")
	cl.Purple("这是一条没有占位符的紫色消息")

	cl.Sbluef("这是一条带有多个占位符的蓝色字符串: %d, %s", 123, "测试")
	cl.Sgreenf("这是一条带有多个占位符的绿色字符串: %s, %d", "测试", 456)
	cl.Sredf("这是一条带有多个占位符的红色字符串: %s, %d", "测试", 789)
	cl.Syellowf("这是一条带有多个占位符的黄色字符串: %d, %s", 123, "测试")
	cl.Spurplef("这是一条带有多个占位符的紫色字符串: %s, %d", "测试", 456)

	cl.Sblue("这是一条没有占位符的蓝色字符串")
	cl.Sgreen("这是一条没有占位符的绿色字符串")
	cl.Sred("这是一条没有占位符的红色字符串")
	cl.Syellow("这是一条没有占位符的黄色字符串")
	cl.Spurple("这是一条没有占位符的紫色字符串")

	cl.PrintSuccessf("这是一条带有多个占位符的成功消息: %d, %s", 123, "测试")
	cl.PrintErrorf("这是一条带有多个占位符的错误消息: %s, %d", "测试", 456)
	cl.PrintWarningf("这是一条带有多个占位符的警告消息: %s, %d", "测试", 789)
	cl.PrintInfof("这是一条带有多个占位符的信息消息: %d, %s", 123, "测试")
}
