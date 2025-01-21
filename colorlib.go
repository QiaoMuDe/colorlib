package colorlib

import (
	"fmt"
	"strings"
)

type ColorLib struct {
	// 空结构体，用于实现接口
	LevelMap map[string]string // LevelMap 是一个映射，用于将日志级别映射到对应的前缀,// 日志级别映射到对应的前缀, 后面留空, 方便后面拼接提示内容
	ColorMap map[string]string // colorMap 是一个映射，用于将颜色名称映射到对应的 ANSI 颜色代码。
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
	Blue(msg ...any)           // 打印蓝色信息到控制台, 无需占位符
	Green(msg ...any)          // 打印绿色信息到控制台, 无需占位符
	Red(msg ...any)            // 打印红色信息到控制台, 无需占位符
	Yellow(msg ...any)         // 打印黄色信息到控制台, 无需占位符
	Purple(msg ...any)         // 打印紫色信息到控制台, 无需占位符
	Sblue(msg ...any) string   // 返回构造后的蓝色字符串, 无需占位符
	Sgreen(msg ...any) string  // 返回构造后的绿色字符串, 无需占位符
	Sred(msg ...any) string    // 返回构造后的红色字符串, 无需占位符
	Syellow(msg ...any) string // 返回构造后的黄色字符串, 无需占位符
	Spurple(msg ...any) string // 返回构造后的紫色字符串, 无需占位符
	PrintSuccess(msg ...any)   // 打印成功信息到控制台, 无需占位符
	PrintError(msg ...any)     // 打印错误信息到控制台, 无需占位符
	PrintWarning(msg ...any)   // 打印警告信息到控制台, 无需占位符
	PrintInfo(msg ...any)      // 打印信息到控制台, 无需占位符
}

// NewColorLib 函数用于创建一个新的 ColorLib 实例。
func NewColorLib() *ColorLib {
	// 创建一个新的 ColorLib 实例
	cl := &ColorLib{
		LevelMap: map[string]string{
			"success": "[Success] ", // 成功信息级别的前缀
			"error":   "[Error] ",   // 错误信息级别的前缀
			"warning": "[Warning] ", // 警告信息级别的前缀
			"info":    "[Info] ",    // 信息信息级别的前缀
		},
		ColorMap: map[string]string{
			"blue":   "34", // 蓝色文本的 ANSI 颜色代码
			"green":  "32", // 绿色文本的 ANSI 颜色代码
			"red":    "31", // 红色文本的 ANSI 颜色代码
			"yellow": "33", // 黄色文本的 ANSI 颜色代码
			"purple": "35", // 紫色文本的 ANSI 颜色代码
		},
	}

	// 返回 ColorLib 实例
	return cl
}

// printWithColor 方法用于将传入的参数以指定颜色文本形式打印到控制台。
func (c *ColorLib) printWithColor(color string, msg ...any) {
	// 获取颜色代码
	code, ok := c.ColorMap[color]
	if !ok {
		fmt.Println("Invalid color:", color)
		return
	}

	// 使用 strings.Builder 拼接字符串
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("\033[1;%sm", code)) // 添加颜色代码

	// 检查 msg 是否为空
	if len(msg) > 0 {
		builder.WriteString(fmt.Sprint(msg...)) // 拼接消息内容
	} else {
		builder.WriteString(" ") // 如果没有消息，添加一个空格，避免完全空白的输出
	}

	builder.WriteString("\033[0m") // 添加颜色重置代码
	fmt.Println(builder.String())  // 使用 fmt.Println 自动处理换行
}

// returnWithColor 方法用于将传入的参数以指定颜色文本形式返回。
func (c *ColorLib) returnWithColor(color string, msg ...any) string {
	// 获取颜色代码
	code, ok := c.ColorMap[color]
	if !ok {
		return fmt.Sprintf("Invalid color: %s", color)
	}

	// 检查 msg 是否为空
	if len(msg) == 0 {
		return fmt.Sprintf("\033[1;%sm\033[0m", code) // 返回空字符串，但带有颜色代码
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return fmt.Sprintf("\033[1;%sm%s\033[0m", code, combinedMsg)
}

// Bluef 方法用于将传入的参数以蓝色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Bluef(format string, a ...any) {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 printWithColor 方法，传入格式化后的字符串
	c.printWithColor("blue", formattedMsg)
}

// Greenf 方法用于将传入的参数以绿色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Greenf(format string, a ...any) {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 printWithColor 方法，传入格式化后的字符串
	c.printWithColor("green", formattedMsg)
}

// Redf 方法用于将传入的参数以红色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Redf(format string, a ...any) {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 printWithColor 方法，传入格式化后的字符串
	c.printWithColor("red", formattedMsg)
}

// Yellowf 方法用于将传入的参数以黄色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Yellowf(format string, a ...any) {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 printWithColor 方法，传入格式化后的字符串
	c.printWithColor("yellow", formattedMsg)
}

// Purplef 方法用于将传入的参数以紫色文本形式打印到控制台（带占位符）。
func (c *ColorLib) Purplef(format string, a ...any) {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	c.printWithColor("purple", formattedMsg)
}

// Sbluef 方法用于将传入的参数以蓝色文本形式返回（带占位符）。
func (c *ColorLib) Sbluef(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("blue", formattedMsg)
}

// Sgreenf 方法用于将传入的参数以绿色文本形式返回（带占位符）。
func (c *ColorLib) Sgreenf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("green", formattedMsg)
}

// Sredf 方法用于将传入的参数以红色文本形式返回（带占位符）。
func (c *ColorLib) Sredf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("red", formattedMsg)
}

// Syellowf 方法用于将传入的参数以黄色文本形式返回（带占位符）。
func (c *ColorLib) Syellowf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("yellow", formattedMsg)
}

// Spurplef 方法用于将传入的参数以紫色文本形式返回（带占位符）。
func (c *ColorLib) Spurplef(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("purple", formattedMsg)
}

// Blue 方法用于将传入的参数以蓝色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Blue(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.printWithColor("blue", combinedMsg)
}

// Green 方法用于将传入的参数以绿色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Green(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.printWithColor("green", combinedMsg)
}

// Red 方法用于将传入的参数以红色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Red(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.printWithColor("red", combinedMsg)
}

// Yellow 方法用于将传入的参数以黄色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Yellow(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.printWithColor("yellow", combinedMsg)
}

// Purple 方法用于将传入的参数以紫色文本形式打印到控制台（不带占位符）。
func (c *ColorLib) Purple(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回
		return
	}
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.printWithColor("purple", combinedMsg)
}

// Sblue 方法用于将传入的参数以蓝色文本形式返回（不带占位符）。
func (c *ColorLib) Sblue(msg ...any) string {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串
		return ""
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("blue", combinedMsg)
}

// Sgreen 方法用于将传入的参数以绿色文本形式返回（不带占位符）。
func (c *ColorLib) Sgreen(msg ...any) string {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串
		return ""
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("green", combinedMsg)
}

// Sred 方法用于将传入的参数以红色文本形式返回（不带占位符）。
func (c *ColorLib) Sred(msg ...any) string {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串
		return ""
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("red", combinedMsg)
}

// Syellow 方法用于将传入的参数以黄色文本形式返回（不带占位符）。
func (c *ColorLib) Syellow(msg ...any) string {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串
		return ""
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("yellow", combinedMsg)
}

// Spurple 方法用于将传入的参数以紫色文本形式返回（不带占位符）。
func (c *ColorLib) Spurple(msg ...any) string {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串
		return ""
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("purple", combinedMsg)
}

// PromptMsg 方法用于打印带有颜色和前缀的消息。
func (c *ColorLib) PromptMsg(level, color, format string, a ...any) {
	// 获取指定级别对应的前缀
	prefix, ok := c.LevelMap[level]
	if !ok {
		fmt.Println("Invalid level:", level)
		return
	}

	// 创建一个 strings.Builder 来构建消息
	var message strings.Builder
	message.WriteString(prefix)

	// 如果没有参数，直接打印前缀
	if len(a) == 0 {
		c.printWithColor(color, message.String())
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprintf(format, a...)
	message.WriteString(combinedMsg)

	// 打印最终消息
	c.printWithColor(color, message.String())
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
func (c *ColorLib) PrintSuccess(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串或默认消息
		c.PromptMsg("success", "green", "%s", "")
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.PromptMsg("success", "green", "%s", combinedMsg)
}

// PrintError 方法用于将传入的参数以红色文本形式打印到控制台，并在文本前添加一个红色的叉号（不带占位符）。
func (c *ColorLib) PrintError(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串或默认消息
		c.PromptMsg("error", "red", "%s", "")
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.PromptMsg("error", "red", "%s", combinedMsg)
}

// PrintWarning 方法用于将传入的参数以黄色文本形式打印到控制台，并在文本前添加一个黄色的感叹号（不带占位符）。
func (c *ColorLib) PrintWarning(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串或默认消息
		c.PromptMsg("warning", "yellow", "%s", "")
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.PromptMsg("warning", "yellow", "%s", combinedMsg)
}

// PrintInfo 方法用于将传入的参数以蓝色文本形式打印到控制台，并在文本前添加一个蓝色的感叹号（不带占位符）。
func (c *ColorLib) PrintInfo(msg ...any) {
	if len(msg) == 0 {
		// 如果没有传入任何参数，直接返回空字符串或默认消息
		c.PromptMsg("info", "blue", "%s", "")
		return
	}

	// 使用 fmt.Sprint 将 msg 中的所有元素拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	c.PromptMsg("info", "blue", "%s", combinedMsg)
}
