package htotp

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/pschlump/filelib"
)

// func GenerateQRCodeFromURI(uri, fn string) {
func Test_GenerateQRCodeFromURI(t *testing.T) {
	uri := "otpauth://totp/www.2c-why.com:example%402c-why.com?issuer=www.2c-why.com&secret=HLYP32J745ZB7LNW"
	os.MkdirAll("./out", 0755)
	GenerateQRCodeFromURI(uri, "./out/Test_GenerateQRCodeFromURI.png")

	if !CompareFiles("./out/Test_GenerateQRCodeFromURI.png", "./ref/Test_GenerateQRCodeFromURI.png") {
		t.Errorf("Error: QR Code did not match")
	}
}

// func ExtractURIFromQRCodeImage(fn string) (uri string, err error) {
func Test_ExtractURIFromQRCodeImage(t *testing.T) {
	uri, err := ExtractURIFromQRCodeImage("./testdata/Test_GenerateQRCodeFromURI.png")
	if err != nil {
		t.Errorf("Error: unable to extract URI from QR Image")
	}
	expect := "otpauth://totp/www.2c-why.com:example%402c-why.com?issuer=www.2c-why.com&secret=HLYP32J745ZB7LNW"
	if uri != expect {
		t.Errorf("Error: unable to extract URI from QR Image")
	}
}

func CompareFiles(fna, fnb string) bool {
	if !filelib.Exists(fna) {
		return false
	}
	if !filelib.Exists(fnb) {
		return false
	}
	aa, err := ioutil.ReadFile(fna)
	if err != nil {
		return false
	}
	bb, err := ioutil.ReadFile(fnb)
	if err != nil {
		return false
	}

	if bytes.Compare(aa, bb) == 0 {
		return true
	}
	return false
}
