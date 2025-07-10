package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32                = syscall.NewLazyDLL("user32.dll")
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	procGetCursorPos      = user32.NewProc("GetCursorPos")
	procSetCursorPos      = user32.NewProc("SetCursorPos")
	procKeybd_event       = user32.NewProc("keybd_event")
	procMouse_event       = user32.NewProc("mouse_event")
	procSetThreadExecutionState = kernel32.NewProc("SetThreadExecutionState")
)

// 常量定义
const (
	INPUT_MOUSE    = 0
	MOUSEEVENTF_MOVE        = 0x0001
	MOUSEEVENTF_LEFTDOWN    = 0x0002
	MOUSEEVENTF_LEFTUP      = 0x0004
	MOUSEEVENTF_RIGHTDOWN   = 0x0008
	MOUSEEVENTF_RIGHTUP     = 0x0010
	MOUSEEVENTF_MIDDLEDOWN  = 0x0020
	MOUSEEVENTF_MIDDLEUP    = 0x0040
	MOUSEEVENTF_WHEEL       = 0x0800
	KEYEVENTF_KEYUP         = 0x0002
	VK_F15                  = 0x7E
	
	// SetThreadExecutionState flags
	ES_CONTINUOUS        = 0x80000000
	ES_SYSTEM_REQUIRED   = 0x00000001
	ES_DISPLAY_REQUIRED  = 0x00000002
	ES_AWAYMODE_REQUIRED = 0x00000040
)

type POINT struct {
	X, Y int32
}

func getCursorPos() (x, y int32) {
	pt := POINT{}
	procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return pt.X, pt.Y
}

func setCursorPos(x, y int32) {
	procSetCursorPos.Call(uintptr(x), uintptr(y))
}

func mouseEvent(dwFlags, dx, dy, dwData uintptr) {
	procMouse_event.Call(dwFlags, dx, dy, dwData, 0)
}

func keyboardEvent(bVk byte) {
	procKeybd_event.Call(uintptr(bVk), 0, 0, 0)
	time.Sleep(50 * time.Millisecond)
	procKeybd_event.Call(uintptr(bVk), 0, KEYEVENTF_KEYUP, 0)
}

// 防止系统休眠的函数
func preventSleep() {
	// ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED
	procSetThreadExecutionState.Call(ES_CONTINUOUS | ES_SYSTEM_REQUIRED | ES_DISPLAY_REQUIRED)
}

func randomMove() {
	// 获取当前鼠标位置
	x, y := getCursorPos()
	
	// 随机移动方向
	dx := rand.Int31n(21) - 10 // -10 到 10 之间的随机数
	dy := rand.Int31n(21) - 10
	
	// 移动鼠标
	setCursorPos(x+dx, y+dy)
	time.Sleep(100 * time.Millisecond)
	
	// 移回原位
	setCursorPos(x, y)
}

func simulateActivity() {
	fmt.Printf("当前时间: %s, 模拟用户活动\n", time.Now().Format("15:04:05"))
	
	// 方法1: 随机鼠标移动
	randomMove()
	
	// 方法2: 模拟按下无害的按键 (F15)
	keyboardEvent(VK_F15)
	
	// 方法3: 使用 SetThreadExecutionState API 防止系统休眠
	preventSleep()
}

func main() {
	fmt.Println("增强版鼠标自动移动程序已启动")
	fmt.Println("每60秒将模拟一次用户活动以防止系统锁屏")
	fmt.Println("按 Ctrl+C 退出程序")
	
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 设置活动间隔为60秒 (更频繁地模拟活动)
	interval := 60 * time.Second
	
	// 初始调用一次防止系统休眠
	preventSleep()
	
	for {
		// 模拟用户活动
		simulateActivity()
		
		// 等待下一次活动
		fmt.Printf("等待 %v 后再次模拟用户活动...\n", interval)
		time.Sleep(interval)
	}
}
