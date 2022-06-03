package main

// Check that an OTP is correct(valid)

// Note: username is not needed for this because we are supplying the secret.
// Normally in a real system the username and authentication realm are used
// with a database to look up the secret.  The "authenticator" client
// saves the "secret" locally.

import (
	"flag"
	"fmt"
	"os"

	"github.com/pschlump/htotp"
)

var OTP = flag.String("otp", "", "The one time password OTP")
var Secret = flag.String("secret", "", "Secret for this username and authentication realm.")

func main() {
	flag.Parse()
	fns := flag.Args()

	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Additional arguments not allowed.\n")
		os.Exit(1)
	}

	if *OTP == "" {
		fmt.Fprintf(os.Stderr, "Missing --otp 123456\n")
		os.Exit(1)
	}
	if *Secret == "" {
		fmt.Fprintf(os.Stderr, "Missing --secret 'xxxxxxxxxxx'\n")
		os.Exit(1)
	}

	if ok := CheckRfc6238TOTPKey("--not-needed--", *OTP, *Secret); ok {
		fmt.Printf("OTK is valid\n")
	} else {
		fmt.Printf("you may not pass....  Invalid OTK\n")
	}
}

// CheckRfc6238TOTPKey uses the `otp` and `secret` to check that the `otp` is valid.
func CheckRfc6238TOTPKey(un, otp, secret string) bool {
	totp := htotp.NewDefaultTOTP(secret)
	totp.Now()
	return totp.Verify(otp)
}
