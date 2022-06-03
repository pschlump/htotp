package htotp

import (
	"fmt"
	"os"

	"github.com/pschlump/dbgo"
)

// time-based OTP, TOTP - uses the counter as a time window.
type TOTP struct {
	OTP
	timeWindow int
	skew       uint // before, current, after match - probelematic if more than 1
	pskew      uint // before, current match - problematic if more than 1
}

func NewTOTP(secret string, nDigits, timeWindow int, hasher *Hasher) *TOTP {
	return &TOTP{OTP: NewOTP(secret, nDigits, hasher), timeWindow: timeWindow}
}

// xyzzy -test-
func (tt *TOTP) SetSkew(skew, pskew uint) {
	if skew > 1 || pskew > 1 {
		fmt.Fprintf(os.Stderr, "Warning: unusual skew value, %s.  Normally 0 or 1.  Have skew=%d and pskew=%d\n", dbgo.LF(-2), skew, pskew)
	}
	tt.skew = skew
	tt.pskew = pskew
}

// NewDefaultTOTP returns a TOTP with default setup.  6 digits and a time
// window of 30 seconds.  The hashis the default hashing algorythm.
func NewDefaultTOTP(secret string) *TOTP {
	return NewTOTP(secret, 6, 30, nil)
}

// Generate time OTP of given timestamp
func (tt *TOTP) At(timestamp int) string {
	tc := uint(timestamp / tt.timeWindow)    // generate the time window.
	tr := 60 - uint(timestamp%tt.timeWindow) // generate the time window.
	if false {
		fmt.Printf("%s time window tc = %d - mod value = %d %s\n", dbgo.ColorRed, tc, tr, dbgo.ColorReset)
		fmt.Printf("  Path: %s\n", dbgo.LF(-4))
	}
	genotp, err := tt.GenerateOTP(tc) // TOTP for this window
	if err != nil {
		return "000000"
	}
	return genotp
}

// Generate time OTP of given timestamp
func (tt *TOTP) AtTL(timestamp int) (pin2fa string, tr uint) {
	tc := uint(timestamp / tt.timeWindow)                  // generate the time window.
	tr = uint(tt.timeWindow - (timestamp % tt.timeWindow)) // calculate the time remaining
	if false {
		fmt.Printf("%s time window tc = %d - mod value = %d %s\n", dbgo.ColorRed, tc, tr, dbgo.ColorReset)
		fmt.Printf("  Path: %s\n", dbgo.LF(-4))
	}
	pin2fa, tr, err := tt.GenerateOTPTL(tc, tr) // TOTP for this window
	if err != nil {
		return "000000", 0
	}
	return
}

// Generate the current time OTP
func (tt *TOTP) Now() string {
	return tt.At(CurrentTimestamp())
}

// GenerateTOTPAndSecondsLeft
// Generate the current time OTP and number of seconds left in window.
// This is useful for a count-down display of how long the OTP is valid.
func (tt *TOTP) GenerateTOTPAndSecondsLeft() (otp string, timeLeft int) {
	timeCode := CurrentTimestamp() / tt.timeWindow
	timeLeft = (timeCode + 1) * tt.timeWindow
	var err error
	otp, err = tt.GenerateOTP(uint(timeCode))
	if err != nil {
		return "000000", 0
	}
	return
}

// Verify checks the OTP at the current timestamp
func (tt *TOTP) Verify(otp string) bool {
	timestamp := CurrentTimestamp()
	st, mx := 0, 0
	if tt.pskew > 0 {
		st, mx = -int(tt.pskew), 0
	} else {
		st, mx = -int(tt.skew), int(tt.skew)
	}
	for ii := st; ii <= mx; ii++ {
		if otp == tt.At(int(timestamp)+ii) {
			return true
		}
	}
	return false
}

// VerifyAtTimestamp checks the OTP at a given timestamp
func (tt *TOTP) VerifyAtTimestamp(otp string, timestamp int) bool {
	return otp == tt.At(timestamp)
}

// ProvisioningUri returns the provisioning URI for the OTP.  This can then be
// encoded in a QR Code and used to provision an OTP app like Google Authenticator.
//
// See also: https://github.com/google/google-authenticator/wiki/Key-Uri-Format
//
// Parameters:
//	un 			Username - used on the server to lookup the secret.
//  realm		Name of the issuer of the OTP
func (tt *TOTP) ProvisioningUri(un, realm string) string {
	return BuildURI(OtpTypeTotp, tt.secret, un, realm, tt.hasher.HashName, 0, tt.nDigits, tt.timeWindow)
}

func (tt TOTP) GetTimeWindow() int {
	return tt.timeWindow
}

// xyzzy - get skew
// CurrentTimeWindow -- xyzzy
