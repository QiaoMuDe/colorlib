package colorlib

import "fmt"

// SColor 方法根据颜色代码常量打印对应颜色的文本
func (c *ColorLib) SColor(code int, msg ...any) string {
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
	combinedMsg := fmt.Sprint(msg...)

	// 调用 returnWithColor 方法，传入拼接后的字符串
	return c.returnWithColor(color, combinedMsg)
}

// Sblue 方法用于将传入的参数以蓝色文本形式返回（不带占位符）。
func (c *ColorLib) Sblue(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("blue", combinedMsg)
}

// Sgreen 方法用于将传入的参数以绿色文本形式返回（不带占位符）。
func (c *ColorLib) Sgreen(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("green", combinedMsg)
}

// Sred 方法用于将传入的参数以红色文本形式返回（不带占位符）。
func (c *ColorLib) Sred(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("red", combinedMsg)
}

// Syellow 方法用于将传入的参数以黄色文本形式返回（不带占位符）。
func (c *ColorLib) Syellow(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("yellow", combinedMsg)
}

// Spurple 方法用于将传入的参数以紫色文本形式返回（不带占位符）。
func (c *ColorLib) Spurple(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("purple", combinedMsg)
}

// Sblack 方法用于将传入的参数以黑色文本形式返回（不带占位符）。
func (c *ColorLib) Sblack(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("black", combinedMsg)
}

// Scyan 方法用于将传入的参数以青色文本形式返回（不带占位符）。
func (c *ColorLib) Scyan(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("cyan", combinedMsg)
}

// Swhite 方法用于将传入的参数以白色文本形式返回（不带占位符）。
func (c *ColorLib) Swhite(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("white", combinedMsg)
}

// Sgray 方法用于将传入的参数以灰色文本形式返回（不带占位符）。
func (c *ColorLib) Sgray(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("gray", combinedMsg)
}

// Slred 方法用于将传入的参数以亮红色文本形式返回（不带占位符）。
func (c *ColorLib) Slred(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lred", combinedMsg)
}

// Slgreen 方法用于将传入的参数以亮绿色文本形式返回（不带占位符）。
func (c *ColorLib) Slgreen(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lgreen", combinedMsg)
}

// Slyellow 方法用于将传入的参数以亮黄色文本形式返回（不带占位符）。
func (c *ColorLib) Slyellow(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lyellow", combinedMsg)
}

// Slblue 方法用于将传入的参数以亮蓝色文本形式返回（不带占位符）。
func (c *ColorLib) Slblue(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lblue", combinedMsg)
}

// Slgreen 方法用于将传入的参数以亮绿色文本形式返回（不带占位符）。
func (c *ColorLib) Slpurple(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lpurple", combinedMsg)
}

// Slcyan 方法用于将传入的参数以亮青色文本形式返回（不带占位符）。
func (c *ColorLib) Slcyan(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lcyan", combinedMsg)
}

// Slwhite 方法用于将传入的参数以亮白色文本形式返回（不带占位符）。
func (c *ColorLib) Slwhite(msg ...any) string {
	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)
	return c.returnWithColor("lwhite", combinedMsg)
}
