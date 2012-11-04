package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type BITMAPFILEHEADER struct {
	Type uint16
	Size uint16
}

func main() {
	header := BITMAPFILEHEADER{}
	header.Type = 0x4D42

	bitmap := new(bytes.Buffer)
	err := binary.Write(bitmap, binary.LittleEndian, header)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	} else {
		fmt.Printf("%s", bitmap)
	}
}
