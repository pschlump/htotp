package htotp

import (
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

type OTPType int

const (
	OtpTypeTotp OTPType = 1
	OtpTypeHotp OTPType = 2
)

func (xx OTPType) String() string {
	switch xx {
	case OtpTypeTotp:
		return "totp"
	case OtpTypeHotp:
		return "hotp"
	}
	return "--unreachable case--"
}

// BuildURI returns the provisioning URI for the OTP.  This works for both TOTP or HOTP.
// The URI is suitable for encoding in a QR-Code that can then be read by Google Authenticator
// or the command line ./tools/ac with the --import flag.   The ./tools/ac can use the
// --importURI to import the string URI instead of parsing a QR code image.
//
// See also:
//     https://github.com/google/google-authenticator/wiki/Key-Uri-Format
//
// Parameters:
//     otpType：     OTP type, must be one of { "totp", "hotp" }
//     secret:       The secret used to generate the URI
//     accountName:  Username - used on server to retretive the secret.
//     issuerName:   Name of the OTP issuer. This is the title of the OTP entry in Google Authenticator.
//     algorithm:    The algorithm used in the OTP generation, defauilt is "sha1" but other hashes can be used.
//     startCount:   For HOTP the starting counter value.  Ignored in TOTP.
//     digits:       The number of digits in the generated OTP.  Default 6. Pass 0 for the default.
//     period:       The number of seconds before the OTP expires.  Default 30.  Pass 0 for the default.
//
// Returns: the provisioning URI
func BuildURI(otpType OTPType, secret, accountName, issuerName, algorithm string, startCount, digits, period int) string {
	q := url.Values{}
	q.Add("secret", secret)
	if otpType == OtpTypeHotp {
		q.Add("counter", fmt.Sprintf("%d", startCount))
	}
	an := url.QueryEscape(accountName)
	if issuerName != "" {
		an = fmt.Sprintf("%s:%s", url.QueryEscape(issuerName), an)
		q.Add("issuer", issuerName)
	}
	if algorithm != "" && algorithm != "sha1" {
		q.Add("algorithm", strings.ToUpper(algorithm))
	}
	if digits != 0 && digits != 6 {
		q.Add("digits", fmt.Sprintf("%d", digits))
	}
	if period != 0 && period != 30 {
		q.Add("period", fmt.Sprintf("%d", period))
	}
	return fmt.Sprintf("otpauth://%s/%s?%s", otpType, an, q.Encode())
}

// RandomSecret will generate a random secret of given length in a suitable format for use in
// either HOTP or TOTP.
func RandomSecret(length int) string {
	bytes := make([]rune, length)

	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
	for i := range bytes {
		// TODO - replace with cryptograpic random generation!
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(bytes)
}

// CurrentTimestamp will return the current Unix timestamp.
func CurrentTimestamp() int {
	return int(time.Now().Unix())
}

// PadRight adds `pad` to `str` until reaching the `length`
func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

// CheckRfc6238TOTPKey uses the `otp` and `secret` to check that the `otp` is valid.
func CheckRfc6238TOTPKey(un, otp, secret string) bool {
	totp := NewDefaultTOTP(secret)
	totp.Now()
	return totp.Verify(otp)
}

func CheckRfc6238TOTPKeyWithSkew(un, otp, secret string, skew, pskew uint) bool {
	totp := NewDefaultTOTP(secret)
	totp.SetSkew(skew, pskew)
	totp.Now()
	return totp.Verify(otp)
}

func GenerateRfc6238TOTPKey(un, secret string) string {
	totp := NewDefaultTOTP(secret)
	totp.Now()
	s := totp.At(CurrentTimestamp())
	return s
}

func GenerateRfc6238TOTPKeyTL(un, secret string) (pin2fa string, tl uint) {
	totp := NewDefaultTOTP(secret)
	totp.Now()
	pin2fa, tl = totp.AtTL(CurrentTimestamp())
	return
}

func CheckRfc4226TOTPKey(n int, otp, secret string) bool {
	hotp := NewDefaultHOTP(secret)
	return hotp.Verify(otp, n)
}

func GenerateRfc4226HOTPKey(n int, secret string) (string, error) {
	hotp := NewDefaultHOTP(secret)
	s, err := hotp.At(n)
	return s, err
}
