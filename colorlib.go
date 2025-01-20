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
	Bluef(format string, a ...any)           // 打印蓝色信息到控制台（带占位符）
	Greenf(format string, a ...any)          // 打印绿色信息到控制台（带占位符）
	Redf(format string, a ...any)            // 打印红色信息到控制台（带占位符）
	Yellowf(format string, a ...any)         // 打印黄色信息到控制台（带占位符）
	Purplef(format string, a ...any)         // 打印紫色信息到控制台（带占位符）
	Sbluef(format string, a ...any) string   // 返回构造后的蓝色字符串（带占位符）
	Sgreenf(format string, a ...any) string  // 返回构造后的绿色字符串（带占位符）
	Sredf(format string, a ...any) string    // 返回构造后的红色字符串（带占位符）
	Syellowf(format string, a ...any) string // 返回构造后的黄色字符串（带占位符）
	Spurplef(format string, a ...any) string // 返回构造后的紫色字符串（带占位符）
	PrintSuccessf(format string, a ...any)   // 打印成功信息到控制台（带占位符）
	PrintErrorf(format string, a ...any)     // 打印错误信息到控制台（带占位符）
	PrintWarningf(format string, a ...any)   // 打印警告信息到控制台（带占位符）
	PrintInfof(format string, a ...any)      // 打印信息到控制台（带占位符）

	// 新增不带占位符的方法
	Blue(msg string)
	Green(msg string)
	Red(msg string)
	Yellow(msg string)
	Purple(msg string)
	Sblue(msg string) string
	Sgreen(msg string) string
	Sred(msg string) string
	Syellow(msg string) string
	Spurple(msg string) string
	PrintSuccess(msg string)
	PrintError(msg string)
	PrintWarning(msg string)
	PrintInfo(msg string)
}

// NewColorLib 函数用于创建一个新的 ColorLib 实例。
func NewColorLib() *ColorLib {
	return &ColorLib{}
}

// printWithColor 方法用于将传入的参数以指定颜色文本形式打印到控制台。
func (c *ColorLib) printWithColor(color, msg string) {
	code, ok := colorMap[color]
	if !ok {
		fmt.Println("Invalid color:", color)
		return
	}
	fmt.Printf("\033[1;%sm%s\033[0m\n", code, msg)
}

// returnWithColor 方法用于将传入的参数以指定颜色文本形式返回。
func (c *ColorLib) returnWithColor(color, msg string) string {
	code, ok := colorMap[color]
	if !ok {
		return fmt.Sprintf("Invalid color: %s", color)
	}
	return fmt.Sprintf("\033[1;%sm%s\033[0m", code, msg)
}

// Bluef 方法用于将传入的参数以蓝色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Bluef(format string, a ...any) {
	c.printWithColor("blue", fmt.Sprintf(format, a...))
}

// Greenf 方法用于将传入的参数以绿色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Greenf(format string, a ...any) {
	c.printWithColor("green", fmt.Sprintf(format, a...))
}

// Redf 方法用于将传入的参数以红色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Redf(format string, a ...any) {
	c.printWithColor("red", fmt.Sprintf(format, a...))
}

// Yellowf 方法用于将传入的参数以黄色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Yellowf(format string, a ...any) {
	c.printWithColor("yellow", fmt.Sprintf(format, a...))
}

// Purplef 方法用于将传入的参数以紫色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Purplef(format string, a ...any) {
	c.printWithColor("purple", fmt.Sprintf(format, a...))
}

// Sbluef 方法用于将传入的参数以蓝色文本形式返回（带占位符）。
func (c *ColorLib) Sbluef(format string, a ...any) string {
	return c.returnWithColor("blue", fmt.Sprintf(format, a...))
}

// Sgreenf 方法用于将传入的参数以绿色文本形式返回（带占位符）。
func (c *ColorLib) Sgreenf(format string, a ...any) string {
	return c.returnWithColor("green", fmt.Sprintf(format, a...))
}

// Sredf 方法用于将传入的参数以红色文本形式返回（带占位符）。
func (c *ColorLib) Sredf(format string, a ...any) string {
	return c.returnWithColor("red", fmt.Sprintf(format, a...))
}

// Syellowf 方法用于将传入的参数以黄色文本形式返回（带占位符）。
func (c *ColorLib) Syellowf(format string, a ...any) string {
	return c.returnWithColor("yellow", fmt.Sprintf(format, a...))
}

// Spurplef 方法用于将传入的参数以紫色文本形式返回（带占位符）。
func (c *ColorLib) Spurplef(format string, a ...any) string {
	return c.returnWithColor("purple", fmt.Sprintf(format, a...))
}

// Blue 方法用于将传入的参数以蓝色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Blue(msg string) {
	c.printWithColor("blue", msg)
}

// Green 方法用于将传入的参数以绿色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Green(msg string) {
	c.printWithColor("green", msg)
}

// Red 方法用于将传入的参数以红色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Red(msg string) {
	c.printWithColor("red", msg)
}

// Yellow 方法用于将传入的参数以黄色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Yellow(msg string) {
	c.printWithColor("yellow", msg)
}

// Purple 方法用于将传入的参数以紫色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Purple(msg string) {
	c.printWithColor("purple", msg)
}

// Sblue 方法用于将传入的参数以蓝色文本形式返回（不带占位符）。
func (c *ColorLib) Sblue(msg string) string {
	return c.returnWithColor("blue", msg)
}

// Sgreen 方法用于将传入的参数以绿色文本形式返回（不带占位符）。
func (c *ColorLib) Sgreen(msg string) string {
	return c.returnWithColor("green", msg)
}

// Sred 方法用于将传入的参数以红色文本形式返回（不带占位符）。
func (c *ColorLib) Sred(msg string) string {
	return c.returnWithColor("red", msg)
}

// Syellow 方法用于将传入的参数以黄色文本形式返回（不带占位符）。
func (c *ColorLib) Syellow(msg string) string {
	return c.returnWithColor("yellow", msg)
}

// Spurple 方法用于将传入的参数以紫色文本形式返回（不带占位符）。
func (c *ColorLib) Spurple(msg string) string {
	return c.returnWithColor("purple", msg)
}

// PromptMsg 方法用于打印带有颜色和前缀的消息。
func (c *ColorLib) PromptMsg(level, color, format string, a ...any) {
	prefix, ok := LevelMap[level]
	if !ok {
		fmt.Println("Invalid level:", level)
		return
	}

	// 如果 format 是 "%s" 且 a 只有一个参数，直接拼接字符串
	if format == "%s" && len(a) == 1 {
		if str, ok := a[0].(string); ok {
			c.printWithColor(color, prefix+str)
			return
		}
	}

	// 对于其他情况，使用 fmt.Sprintf 进行格式化
	c.printWithColor(color, prefix+fmt.Sprintf(format, a...))
}

// PrintSuccessf 方法用于将传入的参数以绿色文本形式打印到控制台，并在文本前添加一个绿色的勾号（带占位符）。
func (c *ColorLib) PrintSuccessf(format string, a ...any) {
	c.PromptMsg("success", "green", format, a...)
}

// PrintErrorf 方法用于将传入的参数以红色文本形式打印到控制台，并在文本前添加一个红色的叉号（带占位符）。
func (c *ColorLib) PrintErrorf(format string, a ...any) {
	c.PromptMsg("error", "red", format, a...)
}

// PrintWarningf 方法用于将传入的参数以黄色文本形式打印到控制台，并在文本前添加一个黄色的感叹号（带占位符）。
func (c *ColorLib) PrintWarningf(format string, a ...any) {
	c.PromptMsg("warning", "yellow", format, a...)
}

// PrintInfof 方法用于将传入的参数以蓝色文本形式打印到控制台，并在文本前添加一个蓝色的感叹号（带占位符）。
func (c *ColorLib) PrintInfof(format string, a ...any) {
	c.PromptMsg("info", "blue", format, a...)
}

// PrintSuccess 方法用于将传入的参数以绿色文本形式打印到控制台，并在文本前添加一个绿色的勾号（不带占位符）。
func (c *ColorLib) PrintSuccess(msg string) {
	c.PromptMsg("success", "green", "%s", msg)
}

// PrintError 方法用于将传入的参数以红色文本形式打印到控制台，并在文本前添加一个红色的叉号（不带占位符）。
func (c *ColorLib) PrintError(msg string) {
	c.PromptMsg("error", "red", "%s", msg)
}

// PrintWarning 方法用于将传入的参数以黄色文本形式打印到控制台，并在文本前添加一个黄色的感叹号（不带占位符）。
func (c *ColorLib) PrintWarning(msg string) {
	c.PromptMsg("warning", "yellow", "%s", msg)
}

// PrintInfo 方法用于将传入的参数以蓝色文本形式打印到控制台，并在文本前添加一个蓝色的感叹号（不带占位符）。
func (c *ColorLib) PrintInfo(msg string) {
	c.PromptMsg("info", "blue", "%s", msg)
}
