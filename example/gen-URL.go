package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pschlump/htotp"
)

var Un = flag.String("un", "example@www.2c-why.com", "Username for this user.  Default 'example@www.2c-why.com'.")
var Realm = flag.String("auth-realm", "www.2c-why.com", "Authenticaiton realm, defaults to www.2c-why.com")
var Secret = flag.String("secret", "", "Secret for this user, if '' (empty) then a secret will be generated and printed out")
var Raw = flag.Bool("raw", false, "Just print out the OTP, nothing else.")

func main() {
	flag.Parse()
	fns := flag.Args()

	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Additional arguments not allowed.\n")
		os.Exit(1)
	}

	if *Secret == "" {
		secret := htotp.RandomSecret(16)
		Secret = &secret
		if !*Raw {
			fmt.Printf("Generated Secret: %s\n", secret)
		}
	}

	u := GenerateURL(*Un, *Realm, *Secret)

	if *Raw {
		fmt.Printf("%s\n", u)
	} else {
		fmt.Printf("URL: %s\n", u)
	}

}

func GenerateURL(un, realm, secret string) (otpURL string) {
	totp := htotp.NewDefaultTOTP(secret)
	otpURL = totp.ProvisioningUri(un, realm) // otpauth://totp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&issuer=issuerName
	return
}
