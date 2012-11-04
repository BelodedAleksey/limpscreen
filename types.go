package main

import "unsafe"

type (
	HANDLE  uintptr
	HWND    HANDLE
	HDC     HANDLE
	HBITMAP HANDLE
	HGDIOBJ HANDLE
	HGLOBAL HANDLE
	DWORD   uint32
	LONG    int32
	WORD    uint16
)

type RECT struct {
	Left, Top, Right, Bottom int
}

type BITMAP struct {
	Type       int32
	Width      int32
	Height     int32
	WidthBytes int32
	Planes     uint16
	BitsPixel  uint16
	Bits       unsafe.Pointer
}

type BITMAPFILEHEADER struct {
	Type      uint16
	Size      uint32
	Reserved1 uint16
	Reserved2 uint16
	OffBits   uint32
}

type BITMAPINFOHEADER struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

type RGBQUAD struct {
	Blue     byte
	Green    byte
	Red      byte
	Reserved byte
}

type BITMAPINFO struct {
	Header BITMAPINFOHEADER
	Colors [1]RGBQUAD
}

type ImgurImage struct {
	Hash       string `hash`
	DeleteHash string `deletehash`
}
type ImgurLinks struct {
	Original     string `original`
	Small_Square string `small_square`
}
type ImgurUpload struct {
	Image ImgurImage `image`
	Links ImgurLinks `links`
}
type ImgurUploadResponse struct {
	Upload ImgurUpload `upload`
}

const SRCCPY = 0xCC0020
const CAPTUREBLT = 0x40000000

const PW_CLIENTONLY = 0x00000001

const CF_TEXT = 1
const CF_BITMAP = 2
const CF_UNICODETEXT = 13

const VK_SCROLL = 0x91

const SND_ALIAS = int32(0x00010000)
const SND_ASYNC = int32(0x0001)

const GMEM_FIXED = 0x0000
const GMEM_ZEROINIT = 0x0040

const DIB_RGB_COLORS = 0

const HALFTONE = 4

const SM_CXSCREEN = 0
const SM_CYSCREEN = 1
