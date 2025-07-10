# Active-Mouse 防锁屏工具

## 项目简介

这是一个高级防锁屏工具，可以无害地通过模拟用户活动来绕过目前所有主流行为监控管理系统，强制电脑保持亮屏状态。该工具使用 Go 语言编写，通过多种方法综合防止 Windows 系统锁屏。无需联网!！！不用担心安全问题。

## 主要功能

- 随机鼠标移动（移动方向和距离随机）
- 模拟无害按键
- 直接调用 Windows API 阻止系统休眠
- 可在后台静默运行
- 每 60 秒自动模拟一次用户活动

## 使用方法

### 直接运行

双击 `active_mouse.exe` 即可运行程序。程序会在命令行窗口中显示运行状态，每 60 秒自动模拟一次用户活动。

按 `Ctrl+C` 可以退出程序。

### 后台运行

创建以下 VBS 脚本，可以实现后台静默运行：

```vbs
Set WshShell = CreateObject("WScript.Shell")
WshShell.Run """C:\path\to\active_mouse.exe""", 0, False
```

将上面的代码保存为 `start_active_mouse.vbs` 文件（将路径替换为实际的 exe 文件路径），然后双击这个 vbs 文件即可在后台启动程序。

### 开机自启

将程序或其快捷方式放入 Windows 启动文件夹：
`C:\Users\用户名\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`

## 工作原理

本工具使用三种不同的方法来防止系统锁屏：

1. **随机鼠标移动**：每次活动时随机移动鼠标，然后移回原位置
2. **模拟按键**：模拟按下一个无害的按键
3. **系统 API 调用**：直接调用 Windows 的API 来阻止系统进入睡眠状态

## 注意事项

- 本工具仅供学习研究使用
- 请遵守所在组织的安全策略

