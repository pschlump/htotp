package htotp

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/pschlump/dbgo"
	"github.com/pschlump/filelib"
)

// Extract URI from QR Code Image
func ExtractURIFromQRCodeImage(fn string) (uri string, err error) {
	if !filelib.Exists(fn) {
		err = fmt.Errorf("Missing file [%s] with QR Code Image in it.", fn)
		return
	}

	// open and decode image file
	file, err := os.Open(fn)
	if err != nil {
		err = fmt.Errorf("Error: %s at: %s\n", err, dbgo.LF())
		return
	}
	img, _, err := image.Decode(file)
	if err != nil {
		err = fmt.Errorf("Error: %s at: %s\n", err, dbgo.LF())
		return
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		err = fmt.Errorf("Error: %s at: %s\n", err, dbgo.LF())
		return
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		err = fmt.Errorf("Error: %s at: %s\n", err, dbgo.LF())
		return
	}

	if dbQr0 {
		fmt.Printf("%sResult: %s%s\n", dbgo.ColorGreen, result, dbgo.ColorReset)
	}
	uri = fmt.Sprintf("%s", result)
	return

	// ------------------------------------------------------------------------------------------------------------------
	/*
		var newCfg ACConfigItem

		// u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
		uu, err := url.Parse(fmt.Sprintf("%s", result))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s at: %s\n", err, dbgo.LF())
			os.Exit(2)
		}

		// u, err := url.Parse("https://example.org/?a=1&a=2&b=&=3&&&&")
		// 		Scheme:   "https",
		//		User:     url.UserPassword("me", "pass"),
		//		Host:     "example.com",
		//		Path:     "foo/bar",
		//		RawQuery: "x=1&y=2",
		//		Fragment: "anchor",
		//
		//	type URL struct {
		//	    Scheme     string
		//	    Opaque     string    // encoded opaque data
		//	    User       *Userinfo // username and password information
		//	    Host       string    // host or host:port
		//	    Path       string    // path (relative paths may omit leading slash)
		//	    RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
		//	    ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
		//	    RawQuery   string    // encoded query values, without '?'
		//	    Fragment   string    // fragment for references, without '#'
		//	}
		//
		// otpauth://totp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&issuer=issuerName
		fmt.Printf("Scheme: ->%s<- User: ->%s<- Host: ->%s<- RawQuery: ->%s<- Fragment: ->%s<-\n", uu.Scheme, uu.User, uu.Host, uu.RawQuery, uu.Fragment)

		if uu.Scheme != "otpauth" {
			fmt.Fprintf(os.Stderr, "Error: Invalid Scheme in URL, url=[%s] at: %s\n", result, dbgo.LF())
			os.Exit(2)
		}
		newCfg.Name = uu.Path
		qq := uu.Query()
		newCfg.Realm = qq.Get("issuer")
		ss := strings.Split(uu.Path, ":")
		newCfg.Username = ss[1]
		newCfg.Secret = qq.Get("secret")

		if pos := InConfig(gCfg.ACConfig.Local, newCfg.Name); pos == -1 {
			gCfg.ACConfig.Local = append(gCfg.ACConfig.Local, newCfg)
			WriteConfig(gCfg)
		} else {
			gCfg.ACConfig.Local[pos] = newCfg
			WriteConfig(gCfg)
		}
	*/
}

// Generate QR Code Image from URI
func GenerateQRCodeFromURI(uri, fn string) {
	ofp, err := filelib.Fopen(fn, "w")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: unable to write: %s error:%s\n", fn, err)
		return
	}

	qrData_QRUrl := uri

	encoder := qrcode.NewQRCodeWriter()
	encodeHints := map[gozxing.EncodeHintType]interface{}{gozxing.EncodeHintType_ERROR_CORRECTION: "L", gozxing.EncodeHintType_CHARACTER_SET: "ASCII"}
	qrImage, err := encoder.Encode(qrData_QRUrl, gozxing.BarcodeFormat_QR_CODE, 256, 256, encodeHints)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: on encode of QR: %s\n", err)
		return
	}
	if err := png.Encode(ofp, qrImage); err != nil {
		fmt.Fprintf(os.Stderr, "Error: on write of QR: %s\n", err)
		return
	}

}

const dbQr0 = false
