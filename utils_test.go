package htotp

import (
	"testing"
)

const secret = "HLYP32J745ZB7LNW"

func TestBuildURI(t *testing.T) {
	//                                                                                   InitialCount
	//                                                                                   |  Digits
	//                                                                                   |  |  Period
	//              Type of OTP			 Addount-Name          Issuer-Name       Hash-Alg.  |  |
	got := BuildURI(OtpTypeTotp, secret, "example@2c-why.com", "www.2c-why.com", "sha1", 0, 6, 0)

	expect := "otpauth://totp/www.2c-why.com:example%402c-why.com?issuer=www.2c-why.com&secret=HLYP32J745ZB7LNW"
	if got != expect {
		t.Errorf("BuildURI failed\n    expected [%s]\n    got      [%s]\n", expect, got)
	}
}

func TestPadRight(t *testing.T) {
	got := PadRight("abc", "=", 8)
	expect := "abc====="
	if got != expect {
		t.Errorf("Failed to right pad, got [%s] expected [%s]\n", got, expect)
	}
}
