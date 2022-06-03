package main

// Generate a suitable secret for HOTP and TOTP

import (
	"flag"
	"fmt"
	"os"

	"github.com/pschlump/htotp"
)

var Length = flag.Int("lenght", 16, "Default length 6 for the secret")

func main() {
	flag.Parse()
	fns := flag.Args()

	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Additional arguments not allowed.\n")
		os.Exit(1)
	}

	if *Length <= 0 {
		fmt.Fprintf(os.Stderr, "Ivalid secret size, normally 16 long\n")
		os.Exit(1)
	}

	// func RandomSecret(length int) string {
	secret := htotp.RandomSecret(*Length)

	fmt.Printf("Secret: \"%s\"\n", secret)
}
