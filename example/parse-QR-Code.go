package main

// Extract the URI from a QR-Code (actually extract whatever string is in the QR code).

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"git.q8s.co/pschlump/htotp"
)

var Fn = flag.String("fn", "", "The image to parse")
var Out = flag.String("outURI", "", "The file to put the results in, empty is stdout")

func main() {
	flag.Parse()
	fns := flag.Args()

	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Additional arguments not allowed.\n")
		os.Exit(1)
	}

	if *Fn == "" {
		fmt.Fprintf(os.Stderr, "Missing --fn 'file.png'\n")
		os.Exit(1)
	}

	uri, err := htotp.ExtractURIFromQRCodeImage(*Fn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open/process qr image. filename: %s error:%s\n", *Fn, err)
		os.Exit(1)
	}

	if *Out == "" {
		fmt.Printf("%s\n", uri)
	} else {
		ioutil.WriteFile(*Out, []byte(uri+"\n"), 0644)
	}

}
