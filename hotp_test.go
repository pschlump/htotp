package htotp_test

import (
	"testing"

	"git.q8s.co/pschlump/htotp"
)

const secret = "HLYP32J745ZB7LNW"
const hOtp = "333575"

func TestHOTP_At(t *testing.T) {
	hotp_data := htotp.NewDefaultHOTP(secret)
	got, err := hotp_data.At(12345)
	if err != nil {
		t.Errorf("Error: unexpected error: %s\n", err)
	}
	expect := hOtp
	if expect != got {
		t.Errorf("Error: expected %v got %v, HTOP not generated correctly.", expect, got)
	}
}

func TestHOTP_Verify(t *testing.T) {
	hotp_data := htotp.NewDefaultHOTP(secret)
	if !hotp_data.Verify(hOtp, 12345) {
		t.Errorf("Error: verify failed.  Expected true, got false, did not verify.")
	}
}
