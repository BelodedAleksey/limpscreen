package main

import (
	"syscall"
	"unsafe"
)

var (
	modGdi32 = syscall.NewLazyDLL("Gdi32.dll")

	procCreateCompatibleDC     = modGdi32.NewProc("CreateCompatibleDC")
	procCreateCompatibleBitmap = modGdi32.NewProc("CreateCompatibleBitmap")
	procDeleteDC               = modGdi32.NewProc("DeleteDC")
	procSelectObject           = modGdi32.NewProc("SelectObject")
	procDeleteObject           = modGdi32.NewProc("DeleteObject")
	procGetObject              = modGdi32.NewProc("GetObjectW")
	procGetBitmapBits          = modGdi32.NewProc("GetBitmapBits")
	procBitBlt                 = modGdi32.NewProc("BitBlt")
	procGetDIBits              = modGdi32.NewProc("GetDIBits")
	procSetStretchBltMode      = modGdi32.NewProc("SetStretchBltMode")
	procStretchBlt             = modGdi32.NewProc("StretchBlt")
)

func GetDIBits(hdc HDC, hbmp HBITMAP, startScan, scanLines uint32, data unsafe.Pointer, bitmapInfo unsafe.Pointer, useage uint32) bool {
	ret, _, _ := procGetDIBits.Call(uintptr(hdc), uintptr(hbmp), uintptr(startScan), uintptr(scanLines), uintptr(data), uintptr(bitmapInfo), uintptr(useage))
	return ret != 0
}

func CreateCompatibleDC(hdc HDC) HDC {
	ret, _, _ := procCreateCompatibleDC.Call(uintptr(hdc))
	return HDC(ret)
}

func CreateCompatibleBitmap(hdc HDC, width, height int) HBITMAP {
	ret, _, _ := procCreateCompatibleBitmap.Call(uintptr(hdc), uintptr(width), uintptr(height))
	return HBITMAP(ret)
}

func DeleteDC(hdc HDC) bool {
	ret, _, _ := procDeleteDC.Call(
		uintptr(hdc))

	return ret != 0
}

func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	ret, _, _ := procSelectObject.Call(
		uintptr(hdc),
		uintptr(hgdiobj))

	if ret == 0 {
		panic("SelectObject failed")
	}

	return HGDIOBJ(ret)
}

func DeleteObject(hObject HGDIOBJ) bool {
	ret, _, _ := procDeleteObject.Call(
		uintptr(hObject))

	return ret != 0
}

func GetBitmapBits(hbmp HBITMAP, length int32, data unsafe.Pointer) bool {
	ret, _, _ := procGetBitmapBits.Call(uintptr(hbmp), uintptr(length), uintptr(data))
	return ret != 0
}

func GetObject(hgdiobj HGDIOBJ, length uintptr, object *BITMAP) bool {
	ret, _, _ := procGetObject.Call(uintptr(hgdiobj), length, uintptr(unsafe.Pointer(object)))
	return ret != 0
}

func BitBlt(hdcDest HDC, nXDest, nYDest, nWidth, nHeight int32, hdcSrc HDC, nXSrc, nYSrc int32, dwRop DWORD) bool {
	ret, _, _ := procBitBlt.Call(uintptr(hdcDest), uintptr(nXDest), uintptr(nYDest), uintptr(nWidth), uintptr(nHeight), uintptr(hdcSrc), uintptr(nXSrc), uintptr(nYSrc), uintptr(dwRop))
	return ret != 0
}

func SetStretchBltMode(hdc HDC, stretchMode int32) int32 {
	ret, _, _ := procSetStretchBltMode.Call(uintptr(hdc), uintptr(stretchMode))
	return int32(ret)
}

func StretchBlt(hdcDest HDC, XOriginDest, YOriginDest, WidthDest, HeightDest int, hdcSrc HDC, XOriginSrc, YOriginSrc, WidthSrc, HeightSrc int, dwRop DWORD) bool {
	ret, _, _ := procStretchBlt.Call(uintptr(hdcDest), uintptr(XOriginDest), uintptr(YOriginDest), uintptr(WidthDest), uintptr(HeightDest), uintptr(hdcSrc), uintptr(XOriginSrc), uintptr(YOriginSrc), uintptr(WidthSrc), uintptr(HeightSrc))
	return ret != 0
}
