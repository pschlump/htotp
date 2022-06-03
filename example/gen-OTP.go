package main

// Generate a time based One Time Passkey (OTP) given a secret.

import (
	"flag"
	"fmt"
	"os"

	"git.q8s.co/pschlump/htotp"
)

// take the secret from the command line: not very secure - but this is ....... an example!
var Secret = flag.String("secret", "", "Use this as the SECRET")
var Raw = flag.Bool("raw", false, "Just print out the OTP, nothing else.")

func main() {
	flag.Parse()
	fns := flag.Args()

	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "No additional arguments.\n")
		os.Exit(1)
	}

	if *Secret == "" {
		fmt.Fprintf(os.Stderr, "Must supply --secret 'secret' to generate a time based OTP\n")
		os.Exit(1)
	}

	if *Raw {
		fmt.Printf("%s\n", htotp.NewDefaultTOTP(*Secret).Now())
	} else {
		fmt.Printf("Time based OTP is: %s\n", htotp.NewDefaultTOTP(*Secret).Now())
	}
}
