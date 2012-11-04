package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	modUser32 = syscall.NewLazyDLL("user32.dll")
	modWinmm  = syscall.NewLazyDLL("Winmm.dll")

	procFindWindow          = modUser32.NewProc("FindWindowW")
	procGetWindowRect       = modUser32.NewProc("GetWindowRect")
	procGetDC               = modUser32.NewProc("GetDC")
	procReleaseDC           = modUser32.NewProc("ReleaseDC")
	procPrintWindow         = modUser32.NewProc("PrintWindow")
	procOpenClipboard       = modUser32.NewProc("OpenClipboard")
	procCloseClipboard      = modUser32.NewProc("CloseClipboard")
	procEmptyClipboard      = modUser32.NewProc("EmptyClipboard")
	procSetClipboardData    = modUser32.NewProc("SetClipboardData")
	procGetForegroundWindow = modUser32.NewProc("GetForegroundWindow")
	procGetAsyncKeyState    = modUser32.NewProc("GetAsyncKeyState")
	procGetSystemMetrics    = modUser32.NewProc("GetSystemMetrics")

	procPlaySound = modWinmm.NewProc("PlaySound")
)

func GetSystemMetrics(index int) int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(index))
	return int(ret)
}

func PlaySound(soundName string) {
	procPlaySound.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(soundName))), uintptr(0), uintptr(SND_ALIAS|SND_ASYNC))
}

func FindWindow(className string) HWND {
	ret, _, _ := procFindWindow.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(className))),
		uintptr(0))
	return HWND(ret)
}

func GetForegroundWindow() HWND {
	ret, _, _ := procGetForegroundWindow.Call()
	return HWND(ret)
}

func GetWindowRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetWindowRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		panic(fmt.Sprintf("GetWindowRect(%d) failed", hwnd))
	}

	return &rect
}

func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(
		uintptr(hwnd))

	return HDC(ret)
}

func GetScreenDC() HDC {
	ret, _, _ := procGetDC.Call(uintptr(0))

	return HDC(ret)
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}

func ReleaseScreenDC(hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(0),
		uintptr(hDC))

	return ret != 0
}

func PrintWindow(hwnd HWND, hdc HDC, flags uint) bool {
	ret, _, _ := procPrintWindow.Call(
		uintptr(hwnd),
		uintptr(hdc),
		uintptr(flags))

	return ret != 0
}

func OpenClipboard(hWndNewOwner HWND) bool {
	ret, _, _ := procOpenClipboard.Call(
		uintptr(hWndNewOwner))
	return ret != 0
}

func OpenTaskClipboard() bool {
	ret, _, _ := procOpenClipboard.Call(
		uintptr(0))
	return ret != 0
}

func CloseClipboard() bool {
	ret, _, _ := procCloseClipboard.Call()
	return ret != 0
}

func EmptyClipboard() bool {
	ret, _, _ := procEmptyClipboard.Call()
	return ret != 0
}

func SetClipboardData(uFormat uint, hMem HANDLE) HANDLE {
	ret, _, _ := procSetClipboardData.Call(
		uintptr(uFormat),
		uintptr(hMem))
	return HANDLE(ret)
}

func GetAsyncKeyState(vKey int) int {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return int(ret)
}
