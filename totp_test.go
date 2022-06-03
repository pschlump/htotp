package htotp_test

import (
	"fmt"
	"testing"

	"github.com/pschlump/htotp"
)

// var secret = "HLYP32J745ZB7LNW" // in other file.

func TestTOTP_At(t *testing.T) {
	totp := htotp.NewDefaultTOTP(secret)
	if totp.Now() != totp.At(htotp.CurrentTimestamp()) {
		t.Errorf("TOTP time stamp did not match current time.\n")
	}
}

func TestTOTP_NowWithExpiration(t *testing.T) {
	totp := htotp.NewDefaultTOTP(secret)
	otp, exp := totp.GenerateTOTPAndSecondsLeft()
	if db2 {
		fmt.Printf("otp [%s] exp %v\n", otp, exp)
	}
	cts := htotp.CurrentTimestamp()
	if otp != totp.Now() {
		t.Errorf("TOTP generate OTP current key failed.\n") // xyzzy
	}
	if a, b := totp.At(cts+30), totp.At(int(exp)); a != b {
		t.Errorf("TOTP failed to verify expiration based keys. a = [%s] b = [%s]\n", a, b)
	}
}

func TestTOTP_ProvisioningUri(t *testing.T) {
	totp := htotp.NewDefaultTOTP(secret)
	expect := "otpauth://totp/www.2c-why.com:example%40www.2c-why.com?issuer=www.2c-why.com&secret=" + secret
	got := totp.ProvisioningUri("example@www.2c-why.com", "www.2c-why.com")
	if db1 {
		fmt.Printf("URI: %s\n", got)
	}
	if expect != got {
		t.Errorf("ProvisioningURI invalid,\n    expected [%s]\n    got      [%s]\n", expect, got)
	}
}

func TestTOTP_VerifyOTP_01(t *testing.T) {
	// TODO - Verify with skew - inside/outside window
	// TODO - Verify with pskew - inside/outside window
	totp := htotp.NewDefaultTOTP(secret)
	otp := totp.At(112233)
	expect := "362727"
	if otp != expect {
		t.Errorf("VerifyOTP_01 error, expected [%s] got [%s]", expect, otp)
	}
}

const db1 = false
const db2 = false
