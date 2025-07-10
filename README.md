# 鼠标自动移动程序

这个程序可以防止系统因为无操作而锁屏，通过每隔3分钟自动移动鼠标。

## 安装依赖

1. 首先需要安装 Go 语言环境：https://golang.org/dl/

2. 安装 robotgo 依赖库：

在 Windows 上，需要先安装 GCC：
- 安装 MinGW-w64 或 TDM-GCC

然后安装 robotgo 库：
```
go get github.com/go-vgo/robotgo
```

## 编译程序

在程序目录下运行：
```
go build -o mouse_mover.exe mouse_mover.go
```

## 使用方法

直接双击 mouse_mover.exe 运行即可。程序会每隔3分钟移动一下鼠标。
按 Ctrl+C 可以退出程序。

## 设置开机自启

可以将编译好的 exe 文件放入 Windows 启动文件夹：
`C:\Users\用户名\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`
