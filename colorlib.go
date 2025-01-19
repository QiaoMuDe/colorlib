package colorlib

import (
	"fmt"
)

type ColorLib struct {
	// 空结构体，用于实现接口
}

// colorMap 是一个映射，用于将颜色名称映射到对应的 ANSI 颜色代码。
var colorMap = map[string]string{
	"blue":   "34",
	"green":  "32",
	"red":    "31",
	"yellow": "33",
	"purple": "35",
}

// LevelMap 是一个映射，用于将日志级别映射到对应的前缀。
var LevelMap = map[string]string{
	// 日志级别映射到对应的前缀, 后面留空, 方便后面拼接提示内容
	"success": "[√] ",
	"error":   "[×] ",
	"warning": "[!] ",
	"info":    "[i] ",
}

// ColorLibInterface 是一个接口，定义了一组方法，用于打印和返回带有颜色的文本。
type ColorLibInterface interface {
	Blue(format string, a ...any)           // 打印蓝色信息到控制台
	Green(format string, a ...any)          // 打印绿色信息到控制台
	Red(format string, a ...any)            // 打印红色信息到控制台
	Yellow(format string, a ...any)         // 打印黄色信息到控制台
	Purple(format string, a ...any)         // 打印紫色信息到控制台
	Sblue(format string, a ...any) string   // 返回构造后的蓝色字符串
	Sgreen(format string, a ...any) string  // 返回构造后的绿色字符串
	Sred(format string, a ...any) string    // 返回构造后的红色字符串
	Syellow(format string, a ...any) string // 返回构造后的黄色字符串
	Spurple(format string, a ...any) string // 返回构造后的紫色字符串
	PrintSuccess(format string, a ...any)   // 打印成功信息到控制台
	PrintError(format string, a ...any)     // 打印错误信息到控制台
	PrintWarning(format string, a ...any)   // 打印警告信息到控制台
	PrintInfo(format string, a ...any)      // 打印信息到控制台
}

// NewColorLib 函数用于创建一个新的 ColorLib 实例。
func NewColorLib() *ColorLib {
	return &ColorLib{}
}

// printWithColor 方法用于将传入的参数以指定颜色文本形式打印到控制台。
func (c *ColorLib) printWithColor(color, format string, a ...any) {
	code, ok := colorMap[color]
	if !ok {
		fmt.Println("Invalid color:", color)
		return
	}
	msg := fmt.Sprintf(format, a...)
	fmt.Printf("\033[1;%sm%s\033[0m\n", code, msg)
}

// returnWithColor 方法用于将传入的参数以指定颜色文本形式返回。
func (c *ColorLib) returnWithColor(color, format string, a ...any) string {
	code, ok := colorMap[color]
	if !ok {
		return fmt.Sprintf("Invalid color: %s", color)
	}
	msg := fmt.Sprintf(format, a...)
	return fmt.Sprintf("\033[1;%sm%s\033[0m", code, msg)
}

// Blue 方法用于将传入的参数以蓝色文本形式打印到控制台。
func (c *ColorLib) Blue(format string, a ...any) {
	c.printWithColor("blue", format, a...)
}

// Green 方法用于将传入的参数以绿色文本形式打印到控制台。
func (c *ColorLib) Green(format string, a ...any) {
	c.printWithColor("green", format, a...)
}

// Red 方法用于将传入的参数以红色文本形式打印到控制台。
func (c *ColorLib) Red(format string, a ...any) {
	c.printWithColor("red", format, a...)
}

// Yellow 方法用于将传入的参数以黄色文本形式打印到控制台。
func (c *ColorLib) Yellow(format string, a ...any) {
	c.printWithColor("yellow", format, a...)
}

// Purple 方法用于将传入的参数以紫色文本形式打印到控制台。
func (c *ColorLib) Purple(format string, a ...any) {
	c.printWithColor("purple", format, a...)
}

// Sblue 方法用于将传入的参数以蓝色文本形式返回。
func (c *ColorLib) Sblue(format string, a ...any) string {
	return c.returnWithColor("blue", format, a...)
}

// Sgreen 方法用于将传入的参数以绿色文本形式返回。
func (c *ColorLib) Sgreen(format string, a ...any) string {
	return c.returnWithColor("green", format, a...)
}

// Sred 方法用于将传入的参数以红色文本形式返回。
func (c *ColorLib) Sred(format string, a ...any) string {
	return c.returnWithColor("red", format, a...)
}

// Syellow 方法用于将传入的参数以黄色文本形式返回。
func (c *ColorLib) Syellow(format string, a ...any) string {
	return c.returnWithColor("yellow", format, a...)
}

// Spurple 方法用于将传入的参数以紫色文本形式返回。
func (c *ColorLib) Spurple(format string, a ...any) string {
	return c.returnWithColor("purple", format, a...)
}

// PromptMsg 方法用于打印带有颜色和前缀的消息。
func (c *ColorLib) PromptMsg(level, color, format string, a ...any) {
	prefix, ok := LevelMap[level]
	if !ok {
		fmt.Println("Invalid level:", level)
		return
	}

	// 将前缀和消息的格式化字符串拼接后传入 printWithColor 方法
	c.printWithColor(color, prefix+format, a...)
}

// PrintSuccess 方法用于将传入的参数以绿色文本形式打印到控制台，并在文本前添加一个绿色的勾号。
func (c *ColorLib) PrintSuccess(format string, a ...any) {
	c.PromptMsg("success", "green", format, a...)
}

// PrintError 方法用于将传入的参数以红色文本形式打印到控制台，并在文本前添加一个红色的叉号。
func (c *ColorLib) PrintError(format string, a ...any) {
	c.PromptMsg("error", "red", format, a...)
}

// PrintWarning 方法用于将传入的参数以黄色文本形式打印到控制台，并在文本前添加一个黄色的感叹号。
func (c *ColorLib) PrintWarning(format string, a ...any) {
	c.PromptMsg("warning", "yellow", format, a...)
}

// PrintInfo 方法用于将传入的参数以蓝色文本形式打印到控制台，并在文本前添加一个蓝色的感叹号。
func (c *ColorLib) PrintInfo(format string, a ...any) {
	c.PromptMsg("info", "blue", format, a...)
}
