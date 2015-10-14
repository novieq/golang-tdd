package main
import (
	"testing"
	"bytes"
	"image/png"
	"errors"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "5656334")


	if buffer.Len() == 0 {
		t.Errorf("Generated QR Code has no data")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "5656334")

	//png.decode doesnt work on byte slices but on types satisfying the io.Reader interface
	_,err := png.Decode(buffer)
	if err != nil {
		t.Errorf("Generated QRCode is not a png : %s",err)
	}
}

type ErrorWriter struct {}

func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected Error")
}

func TestGenerateQRCodePropagatesError(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "5557685")
	if err == nil ||  err.Error() != "Expected Error" {
		t.Errorf("Error not propagated correctly")
	}
}