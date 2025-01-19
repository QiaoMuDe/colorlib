# ColorLib - Go 语言颜色输出库

`ColorLib` 是一个用于在终端中输出彩色文本的 Go 语言库。它提供了多种颜色和日志级别的支持，使得在控制台中输出带有颜色的文本变得更加简单和直观。通过 `ColorLib`，你可以轻松地为你的命令行工具添加颜色提示，提升用户体验。

## 特点

- **支持多种颜色**：包括蓝色、绿色、红色、黄色和紫色。
- **日志级别支持**：提供了成功、错误、警告和信息四种日志级别，每种级别都有对应的前缀符号。
- **简单易用**：通过简单的 API 调用，即可在控制台中输出带有颜色的文本。
- **灵活性强**：支持直接打印彩色文本，也支持返回彩色字符串，方便进一步处理。

## 颜色映射表

| 颜色名称 | ANSI 颜色代码 |
| -------- | ------------- |
| blue     | 34            |
| green    | 32            |
| red      | 31            |
| yellow   | 33            |
| purple   | 35            |

## 日志级别映射表

| 日志级别 | 前缀符号 |
| -------- | -------- |
| success  | [√]      |
| error    | [×]      |
| warning  | [!]      |
| info     | [i]      |

## 结构体方法

`ColorLib` 结构体实现了 `ColorLibInterface` 接口，提供了以下方法：

- **颜色打印方法**：
  - `Blue(format string, a ...any)`：打印蓝色文本。
  - `Green(format string, a ...any)`：打印绿色文本。
  - `Red(format string, a ...any)`：打印红色文本。
  - `Yellow(format string, a ...any)`：打印黄色文本。
  - `Purple(format string, a ...any)`：打印紫色文本。

- **颜色字符串返回方法**：
  - `Sblue(format string, a ...any) string`：返回蓝色字符串。
  - `Sgreen(format string, a ...any) string`：返回绿色字符串。
  - `Sred(format string, a ...any) string`：返回红色字符串。
  - `Syellow(format string, a ...any) string`：返回黄色字符串。
  - `Spurple(format string, a ...any) string`：返回紫色字符串。

- **日志级别打印方法**：
  - `PrintSuccess(format string, a ...any)`：打印成功信息（绿色）。
  - `PrintError(format string, a ...any)`：打印错误信息（红色）。
  - `PrintWarning(format string, a ...any)`：打印警告信息（黄色）。
  - `PrintInfo(format string, a ...any)`：打印信息（蓝色）。

## 安装与使用

### 下载库

你可以使用 `go get` 命令来下载 `ColorLib` 库：

```bash
go get gitee.com/MM-Q/colorlib
```

### 引入库

在你的 Go 代码中引入 `ColorLib`：

```go
import "gitee.com/MM-Q/colorlib"
```

### 使用示例

以下是一些常见的使用示例：

```go
package main

import (
	"gitee.com/MM-Q/colorlib"
)

func main() {
	cl := colorlib.NewColorLib()

	// 打印彩色文本
	cl.Blue("This is a blue message")
	cl.Green("This is a green message")
	cl.Red("This is a red message")
	cl.Yellow("This is a yellow message")
	cl.Purple("This is a purple message")

	// 返回彩色字符串
	blueMsg := cl.Sblue("This is a blue message")
	greenMsg := cl.Sgreen("This is a green message")
	fmt.Println(blueMsg)
	fmt.Println(greenMsg)

	// 打印日志级别信息
	cl.PrintSuccess("Operation completed successfully")
	cl.PrintError("An error occurred")
	cl.PrintWarning("This is a warning")
	cl.PrintInfo("This is an info message")
}
```

### 输出示例

运行上述代码后，你将在终端中看到如下输出：

```powershell
[蓝色文本] This is a blue message
[绿色文本] This is a green message
[红色文本] This is a red message
[黄色文本] This is a yellow message
[紫色文本] This is a purple message
[√] Operation completed successfully
[×] An error occurred
[!] This is a warning
[i] This is an info message
```
