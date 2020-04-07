package me_win32

import (
	win "github.com/lxn/win"
	"golang.org/x/sys/windows"
	"syscall"
	"unsafe"
)
var(
	// Library
	libuser32 *windows.LazyDLL

	// Functions
	getWindowText 				*windows.LazyProc
	enumWindows 				*windows.LazyProc
	setWindowLongPtr            *windows.LazyProc
	isWindow					*windows.LazyProc
	isWindowEnabled 			*windows.LazyProc
	isWindowVisible             *windows.LazyProc
)

func init() {
	is64bit := unsafe.Sizeof(uintptr(0)) == 8

	// Library
	libuser32 = windows.NewLazySystemDLL("user32.dll")

	// Functions
	getWindowText = libuser32.NewProc("GetWindowTextW")
	enumWindows = libuser32.NewProc("EnumWindows")
	isWindow = libuser32.NewProc("IsWindow")
	isWindowEnabled = libuser32.NewProc("IsWindowEnabled")
	isWindowVisible = libuser32.NewProc("IsWindowVisible")
	// On 32 bit SetWindowLongPtrW is not available
	if is64bit {
		setWindowLongPtr = libuser32.NewProc("SetWindowLongPtrW")
	} else {
		setWindowLongPtr = libuser32.NewProc("SetWindowLongW")
	}
}

func GetWindowText(hWnd win.HWND, lpString *uint16,maxCount int) uintptr {
	ret, _, _ := syscall.Syscall(getWindowText.Addr(), 3,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(maxCount))
	return ret
}

func EnumWindows(lpEnumFunc, lParam uintptr) bool {
	ret, _, _ := syscall.Syscall(enumWindows.Addr(), 2,
		lpEnumFunc,
		lParam,
		0)
	return ret != 0
}

func IsWindow(hWnd win.HWND)bool {
	ret, _, _ := syscall.Syscall(isWindow.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}
func IsWindowEnabled(hWnd win.HWND) bool {
	ret, _, _ := syscall.Syscall(isWindowEnabled.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

func IsWindowVisible(hWnd win.HWND) bool {
	ret, _, _ := syscall.Syscall(isWindowVisible.Addr(), 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}

