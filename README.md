# HTOTP - Golang One-Time Password, Two Factor Authentication

HTOTP is based on [PyOTP][PyOTP].   The underlying standards that this uses are [RFC 4226][RFC 4226]
and [RFC 6238][RFC 6238].

It uses `github.com/makiuchi-d/gozxing/qrcode` for generation of QR Codes and extraction of URLs from QR codes.  There
are examples and tools that can do this wit the QR code as an image.

This library (or the library and tools) can be used for multi-factor authentication (MFA) or two-factor authentication (2FA).
The tools include a command line authenticator (Similar to Google Authenticator) but it works at the command line.

![MIT License][license-badge]

## Library/Module Installation

```
$ go get gitlab.com/pschlump/htotp
$ go test 
```

If you want to use the tools (the command line authenticator for example) look in the ./tools directory.
Each of them has its own README.md file and follow the build instructions for the tools.  There is
a Makefile for the top level that will compile and test all of the examples.

```
$ cd ./examples
$ make 
$ cd ../tools/ac
$ go build
```













## Examples

### Example 1 - Generate OTK

From `./example/gen-OTK.go` an example of generating a one time passkey based on time:

```
package main

// Generate a time based One Time Passkey (OTP) given a secret.

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/pschlump/htotp"
)

// take the secret from the command line: not very secure - but this is ....... an example!
var Secret = flag.String("secret", "", "Use this as the SECRET")

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

	fmt.Printf("Time based OTP is: %s\n", htotp.NewDefaultTOTP(*Secret).Now())
}
```

### Example 2 - Generate

From `./example/gen-secret.go`.
Generate a "secret" that needs to be stored both on the server and in the authenticator.

```
package main

// Generate a suitable secret for HOTP and TOTP

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/pschlump/htotp"
)

var Length = flag.Int("lenght", 16, "Default length 6 for the secret")

func main() {
	if *Length <= 0 {
		fmt.Fprintf(os.Stderr, "Ivalid secret size, normally 16 long\n")
		os.Exit(1)
	}

	// func RandomSecret(length int) string {
	secret := htotp.RandomSecret(*Length)

	fmt.Printf("Secret: \"%s\"\n", secret)
}
```

### Example 3 - Validate

Using the secret and the one time passkey (OTK) validate it.

Note: username is not needed for this because we are supplying the secret.
Normally in a real system the username and authentication realm are used
with a database to look up the secret.  The "authenticator" client
saves the "secret" locally.

```
package main

// Check that an OTP is correct(valid)


import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/pschlump/htotp"
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
	return totp.Verify(otp, htotp.CurrentTimestamp())
}
```


### Example 4 - Generate URL for use in a QR Code

From `./example/gen-URL.go`:

```
package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/pschlump/htotp"
)

var Un = flag.String("un", "example@www.2c-why.com", "Username for this user.  Default 'example@www.2c-why.com'.")
var Realm = flag.String("auth-realm", "www.2c-why.com", "Authenticaiton realm, defaults to www.2c-why.com")
var Secret = flag.String("secret", "", "Secret for this user, if '' (empty) then a secret will be generated and printed out")

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

	fmt.Printf("URL: %s\n", u)
}

func GenerateURL(un, realm, secret string) (otpURL string) {
	totp := htotp.NewDefaultTOTP(secret)
	otpURL = totp.ProvisioningUri(un, realm) // otpauth://totp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&issuer=issuerName
	return
}
```









## MIT License

HTOTP is licensed under the [MIT License][License]

[license-badge]:   https://img.shields.io/badge/license-MIT-000000.svg
[RFC 4226]: https://tools.ietf.org/html/rfc4226 "RFC 4226"
[RFC 6238]: https://tools.ietf.org/html/rfc6238 "RFC 6238"
[PyOTP]: https://github.com/pyotp/pyotp
[License]: https://gitlab.com/pschlump/htotp/blob/master/LICENSE

