package main

import (
	"syscall"
	"unsafe"
)

var (
	modKernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGlobalAlloc  = modKernel32.NewProc("GlobalAlloc")
	procGlobalLock   = modKernel32.NewProc("GlobalLock")
	procGlobalUnlock = modKernel32.NewProc("GlobalUnlock")
	procCopyMemory   = modKernel32.NewProc("RtlCopyMemory")
	procFreeConsole  = modKernel32.NewProc("FreeConsole")
	procGetLastError = modKernel32.NewProc("GetLastError")
)

func GlobalAlloc(flags uint32, size int32) HGLOBAL {
	ret, _, _ := procGlobalAlloc.Call(uintptr(flags), uintptr(size))
	return HGLOBAL(ret)
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
	ret, _, _ := procGlobalLock.Call(uintptr(hMem))
	return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) {
	procGlobalUnlock.Call(uintptr(hMem))
}

func CopyMemory(dest unsafe.Pointer, src unsafe.Pointer, size int32) {
	procCopyMemory.Call(uintptr(dest), uintptr(src), uintptr(size))
}

func FreeConsole() {
	procFreeConsole.Call()
}

func GetLastError() uintptr {
	ret, _, _ := procGetLastError.Call()
	return ret
}
