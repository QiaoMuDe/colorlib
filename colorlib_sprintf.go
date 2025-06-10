package colorlib

import "fmt"

// SprintColorf 方法根据颜色代码常量打印对应颜色的文本
func (c *ColorLib) SprintColorf(code int, format string, a ...any) string {
	// 根据颜色代码获取颜色名称
	var color string
	switch code {
	case Black:
		color = "black"
	case Red:
		color = "red"
	case Green:
		color = "green"
	case Yellow:
		color = "yellow"
	case Blue:
		color = "blue"
	case Purple:
		color = "purple"
	case Cyan:
		color = "cyan"
	case White:
		color = "white"
	case Gray:
		color = "gray"
	case Lred:
		color = "lred"
	case Lgreen:
		color = "lgreen"
	case Lyellow:
		color = "lyellow"
	case Lblue:
		color = "lblue"
	case Lpurple:
		color = "lpurple"
	case Lcyan:
		color = "lcyan"
	case Lwhite:
		color = "lwhite"
	default:
		return fmt.Sprintf("Invalid color code: %d", code)
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor(color, formattedMsg)
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

// Sblackf 方法用于将传入的参数以黑色文本形式返回（带占位符）。
func (c *ColorLib) Sblackf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("black", formattedMsg)
}

// Scyanf 方法用于将传入的参数以青色文本形式返回（带占位符）。
func (c *ColorLib) Scyanf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("cyan", formattedMsg)
}

// Swhitef 方法用于将传入的参数以白色文本形式返回（带占位符）。
func (c *ColorLib) Swhitef(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("white", formattedMsg)
}

// Sgrayf 方法用于将传入的参数以灰色文本形式返回（带占位符）。
func (c *ColorLib) Sgrayf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("gray", formattedMsg)
}

// Slredf 方法用于将传入的参数以亮红色文本形式返回（带占位符）。
func (c *ColorLib) Slredf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lred", formattedMsg)
}

// Slgreenf 方法用于将传入的参数以亮绿色文本形式返回（带占位符）。
func (c *ColorLib) Slgreenf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lgreen", formattedMsg)
}

// Slyellowf 方法用于将传入的参数以亮黄色文本形式返回（带占位符）。
func (c *ColorLib) Slyellowf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lyellow", formattedMsg)
}

// Slbluef 方法用于将传入的参数以亮蓝色文本形式返回（带占位符）。
func (c *ColorLib) Slbluef(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lblue", formattedMsg)
}

// Slgreenf 方法用于将传入的参数以亮绿色文本形式返回（带占位符）。
func (c *ColorLib) Slpurplef(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lpurple", formattedMsg)
}

// Slcyanf 方法用于将传入的参数以亮青色文本形式返回（带占位符）。
func (c *ColorLib) Slcyanf(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lcyan", formattedMsg)
}

// Slwhitef 方法用于将传入的参数以亮白色文本形式返回（带占位符）。
func (c *ColorLib) Slwhitef(format string, a ...any) string {
	// 使用 fmt.Sprintf 格式化参数
	formattedMsg := fmt.Sprintf(format, a...)

	// 调用 returnWithColor 方法，传入格式化后的字符串
	return c.returnWithColor("lwhite", formattedMsg)
}
