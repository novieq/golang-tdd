package main
import "testing"

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	qrcode := GenerateQRCode("5656334")

	if qrcode == nil {
		t.Errorf("Generated QR Code is nil")
	}

	if len(qrcode) == 0 {
		t.Errorf("Generated QR Code has no data")
	}
}
