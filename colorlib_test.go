package colorlib

import (
	"testing"
)

func TestColorLib(t *testing.T) {
	// 创建 ColorLib 实例
	colorLib := NewColorLib()

	// 测试 Blue 方法
	colorLib.Blue("This is a blue message")

	// 测试 Green 方法
	colorLib.Green("This is a green message")

	// 测试 Red 方法
	colorLib.Red("This is a red message")

	// 测试 Yellow 方法
	colorLib.Yellow("This is a yellow message")

	// 测试 Purple 方法
	colorLib.Purple("This is a purple message")

	// 测试 Sblue 方法
	blueStr := colorLib.Sblue("This is a blue string")
	if blueStr != "\033[1;34mThis is a blue string\033[0m" {
		t.Errorf("Sblue method returned incorrect result")
	}

	// 测试 Sgreen 方法
	greenStr := colorLib.Sgreen("This is a green string")
	if greenStr != "\033[1;32mThis is a green string\033[0m" {
		t.Errorf("Sgreen method returned incorrect result")
	}

	// 测试 Sred 方法
	redStr := colorLib.Sred("This is a red string")
	if redStr != "\033[1;31mThis is a red string\033[0m" {
		t.Errorf("Sred method returned incorrect result")
	}

	// 测试 Syellow 方法
	yellowStr := colorLib.Syellow("This is a yellow string")
	if yellowStr != "\033[1;33mThis is a yellow string\033[0m" {
		t.Errorf("Syellow method returned incorrect result")
	}

	// 测试 Spurple 方法
	purpleStr := colorLib.Spurple("This is a purple string")
	if purpleStr != "\033[1;35mThis is a purple string\033[0m" {
		t.Errorf("Spurple method returned incorrect result")
	}

	// 测试 PrintSuccess 方法
	colorLib.PrintSuccess("This is a success message")

	// 测试 PrintError 方法
	colorLib.PrintError("This is an error message")

	// 测试 PrintWarning 方法
	colorLib.PrintWarning("This is a warning message")

	// 测试 PrintInfo 方法
	colorLib.PrintInfo("This is an info message")

	// 测试无效颜色
	colorLib.printWithColor("invalidColor", "This is an invalid color message")
	colorLib.returnWithColor("invalidColor", "This is an invalid color string")
}
