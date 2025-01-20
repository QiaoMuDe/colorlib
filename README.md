# ColorLib - Go 语言的彩色终端输出库

## 用途和特点

`ColorLib` 是一个用于在 Go 语言中实现彩色终端输出的库。它提供了丰富的接口，用于打印和返回带有颜色的文本，并支持自定义颜色和日志级别。该库的主要特点包括：

- 支持多种颜色输出，包括蓝色、绿色、红色、黄色和紫色。
- 提供带占位符和不带占位符的打印方法。
- 支持日志级别前缀，方便打印带有提示信息的消息。
- 简洁易用的接口，方便开发者快速集成。

## 定义的颜色和数字

以下是库中定义的颜色及其对应的 ANSI 颜色代码：

| 颜色名称 | ANSI 颜色代码 |
| -------- | ------------- |
| blue     | 34            |
| green    | 32            |
| red      | 31            |
| yellow   | 33            |
| purple   | 35            |

## 提示信息级别和名称

以下是库中定义的日志级别及其对应的前缀：

| 日志级别 | 前缀名称 |
| -------- | -------- |
| success  | [√]      |
| error    | [×]      |
| warning  | [!]      |
| info     | [i]      |

## 结构体方法和函数

### 结构体方法

`ColorLib` 结构体实现了以下方法：

| 方法名称                    | 描述                                 |
| --------------------------- | ------------------------------------ |
| Bluef(format, a...)         | 打印蓝色信息到控制台（带占位符）     |
| Greenf(format, a...)        | 打印绿色信息到控制台（带占位符）     |
| Redf(format, a...)          | 打印红色信息到控制台（带占位符）     |
| Yellowf(format, a...)       | 打印黄色信息到控制台（带占位符）     |
| Purplef(format, a...)       | 打印紫色信息到控制台（带占位符）     |
| Sbluef(format, a...)        | 返回构造后的蓝色字符串（带占位符）   |
| Sgreenf(format, a...)       | 返回构造后的绿色字符串（带占位符）   |
| Sredf(format, a...)         | 返回构造后的红色字符串（带占位符）   |
| Syellowf(format, a...)      | 返回构造后的黄色字符串（带占位符）   |
| Spurplef(format, a...)      | 返回构造后的紫色字符串（带占位符）   |
| Blue(msg)                   | 打印蓝色信息到控制台（不带占位符）   |
| Green(msg)                  | 打印绿色信息到控制台（不带占位符）   |
| Red(msg)                    | 打印红色信息到控制台（不带占位符）   |
| Yellow(msg)                 | 打印黄色信息到控制台（不带占位符）   |
| Purple(msg)                 | 打印紫色信息到控制台（不带占位符）   |
| Sblue(msg)                  | 返回构造后的蓝色字符串（不带占位符） |
| Sgreen(msg)                 | 返回构造后的绿色字符串（不带占位符） |
| Sred(msg)                   | 返回构造后的红色字符串（不带占位符） |
| Syellow(msg)                | 返回构造后的黄色字符串（不带占位符） |
| Spurple(msg)                | 返回构造后的紫色字符串（不带占位符） |
| PrintSuccessf(format, a...) | 打印成功信息到控制台（带占位符）     |
| PrintErrorf(format, a...)   | 打印错误信息到控制台（带占位符）     |
| PrintWarningf(format, a...) | 打印警告信息到控制台（带占位符）     |
| PrintInfof(format, a...)    | 打印信息到控制台（带占位符）         |
| PrintSuccess(msg)           | 打印成功信息到控制台（不带占位符）   |
| PrintError(msg)             | 打印错误信息到控制台（不带占位符）   |
| PrintWarning(msg)           | 打印警告信息到控制台（不带占位符）   |
| PrintInfo(msg)              | 打印信息到控制台（不带占位符）       |

### 包含的函数

- `NewColorLib()`：创建一个新的 `ColorLib` 实例。

## 下载和使用

### 下载

通过 Go 模块管理工具下载 `ColorLib`：

```bash
go get gitee.com/MM-Q/colorlib
```

### 引入和使用

在您的 Go 代码中引入 `ColorLib`：

```go
package main

import (
	"gitee.com/MM-Q/colorlib"
)

func main() {
	cl := colorlib.NewColorLib()

	// 打印带有颜色的文本
	cl.Blue("这是一条蓝色的消息")
	cl.Greenf("这是一条绿色的消息：%s", "Hello, ColorLib!")

	// 返回带有颜色的字符串
	coloredString := cl.Sred("这是一条红色的字符串")
	fmt.Println(coloredString)

	// 打印带有日志级别的消息
	cl.PrintSuccess("操作成功！")
	cl.PrintError("发生了一个错误")
	cl.PrintWarning("请注意：这是一个警告")
	cl.PrintInfo("这是一条普通信息")
}
```

## 常用用法

以下是 `ColorLib` 的一些常用用法示例：

### 打印彩色文本

```go
cl := colorlib.NewColorLib()
cl.Blue("蓝色文本")
cl.Greenf("绿色文本：%s", "带占位符")
```

### 返回彩色字符串

```go
coloredString := cl.Spurple("紫色字符串")
fmt.Println(coloredString)
```

### 打印日志级别消息

```go
cl.PrintSuccess("操作成功！")
cl.PrintError("发生错误：参数无效")
cl.PrintWarning("警告：磁盘空间不足")
cl.PrintInfo("正在处理数据...")
```
