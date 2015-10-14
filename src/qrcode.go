package main

import (
	"fmt"
	"image"
	"image/png"

	"os"
	"io"
)

func main() {
	fmt.Println("Hello QR Code")
	file,_ := os.Create("qrcode.png")
	defer file.Close()
	GenerateQRCode(file, "555-8866")

}

func GenerateQRCode(w io.Writer,code string) error{
	img := image.NewNRGBA(image.Rect(0,0,21,21))
	return png.Encode(w,img)
}
