package main

// Generate a QR code from a URI

import (
	"flag"
	"fmt"
	"os"

	"github.com/pschlump/htotp"
)

var URI = flag.String("URI", "", "The URI to encode in the image")
var OutFn = flag.String("out-img", "", "File name to put the QR Image into.")

func main() {
	flag.Parse()
	fns := flag.Args()

	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Additional arguments not allowed.\n")
		os.Exit(1)
	}

	if *URI == "" {
		fmt.Fprintf(os.Stderr, "Missing --URI 'string'\n")
		os.Exit(1)
	}
	if *OutFn == "" {
		fmt.Fprintf(os.Stderr, "Missing --OutFn 'file-name.png'\n")
		os.Exit(1)
	}

	htotp.GenerateQRCodeFromURI(*URI, *OutFn)
}
