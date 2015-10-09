package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code")
	qrcode := GenerateQRCode("555-8866")
	ioutil.WriteFile("qrcode.png",qrcode,0644)

}
