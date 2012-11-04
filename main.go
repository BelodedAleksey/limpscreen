package main

import (
	"bytes"
	_ "encoding/binary"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	_ "net/url"
	_ "os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	FreeConsole()
	for {
		if isScreenshotKeyPressed() {
			takeScreenshot()
			time.Sleep(2 * time.Second)
		}

		// Yield hack
		time.Sleep(100 * time.Millisecond)
	}
}

func isScreenshotKeyPressed() bool {
	return (GetAsyncKeyState(VK_SCROLL) & 0x8000) != 0
}

func takeScreenshot() {

	hwnd := GetForegroundWindow()

	rc := GetWindowRect(hwnd)

	hdcScreen := GetScreenDC()
	hdcTarget := CreateCompatibleDC(hdcScreen)
	hbmp := CreateCompatibleBitmap(hdcScreen, rc.Right-rc.Left, rc.Bottom-rc.Top)

	SelectObject(hdcTarget, HGDIOBJ(hbmp))

	if !BitBlt(hdcTarget, 0, 0, int32(rc.Right-rc.Left), int32(rc.Bottom-rc.Top), hdcScreen, int32(rc.Left), int32(rc.Top), SRCCPY|CAPTUREBLT) {
		fmt.Printf("Error: BitBlt: %d\n", GetLastError())
	}

	var data = make([]byte, (rc.Right-rc.Left)*4*(rc.Bottom-rc.Top))

	// Get color data
	GetBitmapBits(hbmp, int32((rc.Right-rc.Left)*4*(rc.Bottom-rc.Top)), unsafe.Pointer(&data[0]))

	bmp := BITMAP{}
	GetObject(HGDIOBJ(hbmp), unsafe.Sizeof(BITMAP{}), &bmp)

	DeleteDC(hdcTarget)
	DeleteObject(HGDIOBJ(hbmp))
	ReleaseScreenDC(hdcScreen)

	bitmap := buildBitmap(bmp, data)
	uploadToImgur(bitmap)
}

func uploadToImgur(dataBuffer []byte) {
	fmt.Printf("Uploading...")
	img := bytes.NewBuffer(dataBuffer)

	bodyBuffer := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(bodyBuffer)

	// Send the API Key
	key, _ := multipartWriter.CreateFormField("key")
	key.Write([]byte("0ad2497b4d9ebe9870f71fe6ea18cb3f"))

	// Send the file
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", `form-data; name="image"; filename="blob"`)
	h.Set("Content-Type", "image/png")
	partWriter, err := multipartWriter.CreatePart(h)
	if err != nil {
		fmt.Printf("Create Part: %s\n", err)
	}

	_, err = io.Copy(partWriter, img)
	if err != nil {
		fmt.Printf("Copy: %s\n", err)
	}

	multipartWriter.Close()

	req, err := http.NewRequest("POST", "http://api.imgur.com/2/upload.json", bodyBuffer)
	if err != nil {
		fmt.Printf("Create Req: %s\n", err)
	}

	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	jsonResp := ImgurUploadResponse{}
	dec := json.NewDecoder(bytes.NewBuffer(body))
	err = dec.Decode(&jsonResp)
	if err != nil {
		fmt.Printf("%s", err)
	}

	// Save the refs
	storeToFuriousMist(&jsonResp)

	PlaySound("whatever")

	hMem := GlobalAlloc(GMEM_FIXED|GMEM_ZEROINIT, 1024)
	ptr := GlobalLock(hMem)
	CopyMemory(ptr, unsafe.Pointer(syscall.StringToUTF16Ptr(jsonResp.Upload.Links.Original)), 1023)
	GlobalUnlock(hMem)

	// copy the url to clipboard
	OpenTaskClipboard()
	EmptyClipboard()
	SetClipboardData(CF_UNICODETEXT, HANDLE(hMem))
	CloseClipboard()
}

func storeToFuriousMist(data *ImgurUploadResponse) {
	origParts := strings.Split(data.Upload.Links.Original, "/")
	thumbParts := strings.Split(data.Upload.Links.Small_Square, "/")

	http.PostForm("http://furious-mist-4119.herokuapp.com/uploaded", url.Values{
		"hash":       {data.Upload.Image.Hash},
		"deletehash": {data.Upload.Image.DeleteHash},
		"orig":       {origParts[cap(origParts)-1]},
		"thumb":      {thumbParts[cap(thumbParts)-1]},
	})
}

func buildBitmap(bmp BITMAP, dataBuffer []byte) []byte {
	pngBuffer := image.NewNRGBA(image.Rect(0, 0, int(bmp.Width), int(bmp.Height)))
	var ptr int = 0
	for y := 0; y < int(bmp.Height); y++ {
		for x := 0; x < int(bmp.Width); x++ {
			pngBuffer.Set(x, y, color.NRGBA{uint8(dataBuffer[ptr+2]), uint8(dataBuffer[ptr+1]), uint8(dataBuffer[ptr+0]), 255})
			ptr += 4
		}
	}

	image := new(bytes.Buffer)
	png.Encode(image, pngBuffer)

	return image.Bytes()
}
