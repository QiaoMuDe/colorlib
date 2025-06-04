package colorlib

import "fmt"

// NewColorLib 函数用于创建一个新的 ColorLib 实例
func NewColorLib() *ColorLib {
	// 创建一个新的 ColorLib 实例
	cl := &ColorLib{
		levelMap: make(map[string]string),
		colorMap: make(map[string]int),
	}

	// 初始化颜色映射
	for k, v := range colorMap {
		cl.colorMap[k] = v
	}

	// 初始化日志级别映射
	for k, v := range levelMap {
		cl.levelMap[k] = v
	}

	return cl
}

// printWithColor 方法用于将传入的参数以指定颜色文本形式打印到控制台。
func (c *ColorLib) printWithColor(color string, msg ...any) {
	// 检查是否禁用颜色输出
	if c.NoColor {
		fmt.Print(msg...)
		return
	}

	// 获取颜色代码
	code, ok := c.colorMap[color]
	if !ok {
		fmt.Println("Invalid color:", color)
		return
	}

	// 清理缓冲区
	c.formatBuffer.Reset()

	// 检查是否禁用粗体输出
	if c.NoBold {
		// 写入前缀
		c.formatBuffer.WriteString(fmt.Sprintf("\033[%dm", code))
	} else {
		// 写入前缀
		c.formatBuffer.WriteString(fmt.Sprintf("\033[1;%dm", code))
	}

	// 写入消息
	if len(msg) > 0 {
		c.formatBuffer.WriteString(fmt.Sprint(msg...)) // 拼接消息内容
	} else {
		c.formatBuffer.WriteString(" ") // 如果没有消息，添加一个空格，避免完全空白的输出
	}

	// 写入颜色重置代码
	c.formatBuffer.WriteString(fmt.Sprintf("\033[%dm", reset))

	// 使用 fmt.Print 根据外部调用选择性添加换行符
	fmt.Print(c.formatBuffer.String())

	// 重置缓冲区
	c.formatBuffer.Reset()
}

// returnWithColor 方法用于将传入的参数以指定颜色文本形式返回。
func (c *ColorLib) returnWithColor(color string, msg ...any) string {
	// 检查是否禁用颜色输出
	if c.NoColor {
		return fmt.Sprint(msg...)
	}

	// 获取颜色代码
	code, ok := c.colorMap[color]
	if !ok {
		return fmt.Sprintf("Invalid color: %s", color)
	}

	// 检查 msg 是否为空
	if len(msg) == 0 {
		if c.NoBold {
			return fmt.Sprintf("\033[%dm\033[%dm", code, reset) // 返回空字符串，但带有颜色代码
		} else {
			return fmt.Sprintf("\033[1;%dm\033[%dm", code, reset) // 返回空字符串，但带有颜色代码
		}
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprint(msg...)

	// 清理缓冲区
	c.formatBuffer.Reset()

	// 写入前缀
	if c.NoBold {
		c.formatBuffer.WriteString(fmt.Sprintf("\033[%dm", code))
	} else {
		c.formatBuffer.WriteString(fmt.Sprintf("\033[1;%dm", code)) // 添加颜色代码，并加粗
	}

	// 写入消息
	c.formatBuffer.WriteString(combinedMsg) // 拼接消息内容

	// 写入颜色重置代码
	c.formatBuffer.WriteString(fmt.Sprintf("\033[%dm", reset))

	// 获取最终字符串
	result := c.formatBuffer.String()

	// 重置缓冲区
	c.formatBuffer.Reset()

	return result
}

// PromptMsg 方法用于打印带有颜色和前缀的消息。
func (c *ColorLib) PromptMsg(level, color, format string, a ...any) {
	// 获取指定级别对应的前缀
	prefix, ok := c.levelMap[level]
	if !ok {
		fmt.Println("Invalid level:", level)
		return
	}

	// 清理缓冲区
	c.formatBuffer.Reset()

	// 写入前缀
	c.formatBuffer.WriteString(prefix)

	// 如果没有参数，直接打印前缀
	if len(a) == 0 {
		if c.NoColor {
			fmt.Print(c.formatBuffer.String())
			c.formatBuffer.Reset()
		} else {
			c.printWithColor(color, c.formatBuffer.String())
			c.formatBuffer.Reset()
		}
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprintf(format, a...)

	// 写入消息
	c.formatBuffer.WriteString(combinedMsg)

	// 打印最终消息
	if c.NoColor {
		fmt.Print(c.formatBuffer.String())
	} else {
		c.printWithColor(color, c.formatBuffer.String())
	}

	// 重置缓冲区
	c.formatBuffer.Reset()
}

// PMsg 方法用于打印带有颜色和前缀的消息。
func (c *ColorLib) PMsg(level, color, format string, a ...any) {
	// 获取指定级别对应的前缀
	prefix, ok := c.levelMap[level]
	if !ok {
		fmt.Println("Invalid level:", level)
		return
	}

	// 清理缓冲区
	c.formatBuffer.Reset()

	// 写入前缀
	c.formatBuffer.WriteString(prefix)

	// 如果没有参数，直接打印前缀
	if len(a) == 0 {
		if c.NoColor {
			fmt.Print(c.formatBuffer.String())
			c.formatBuffer.Reset()
		} else {
			c.printWithColor(color, c.formatBuffer.String())
			c.formatBuffer.Reset()
		}
		return
	}

	// 使用 fmt.Sprint 将所有参数拼接成一个字符串
	combinedMsg := fmt.Sprintf(format, a...)

	// 写入消息
	c.formatBuffer.WriteString(combinedMsg)

	// 打印最终消息
	if c.NoColor {
		fmt.Print(c.formatBuffer.String())
	} else {
		c.printWithColor(color, c.formatBuffer.String())
	}

	// 重置缓冲区
	c.formatBuffer.Reset()
}
