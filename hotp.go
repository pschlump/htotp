package htotp

// OTP Based on HMAC-based and a counter.

type HOTP struct {
	OTP
}

func NewHOTP(secret string, nDigits int, hasher *Hasher) *HOTP {
	return &HOTP{OTP: NewOTP(secret, nDigits, hasher)}
}

func NewDefaultHOTP(secret string) *HOTP {
	return NewHOTP(secret, 6, nil)
}

// Generates the OTP for the given count.
func (hh *HOTP) At(count int) (rv string, err error) {
	return hh.GenerateOTP(uint(count))
}

func (hh *HOTP) AtTL(count int) (rv string, tl uint, err error) {
	return hh.GenerateOTPTL(uint(count), 1)
}

// Verify will check if the `otp` that is passed is valid
// given the `count` and the Secret.
// 	otp		The one time key that has been sent from the client
// 	countr	The counter that is changing over time.
func (hh *HOTP) Verify(otp string, count int) bool {
	genotp, err := hh.At(count)
	if err != nil {
		return false
	}
	return otp == genotp
}

// ProvisioningUri returns the provisioning URI for the OTP.  This can then be
// encoded in a QR Code and used to provision an OTP app like Google Authenticator.
//
// See also: https://github.com/google/google-authenticator/wiki/Key-Uri-Format
//
// Parameters:
//	un 			Username - used on the server to lookup the secret.
//  realm		Name of the issuer of the OTP
//	cnt			Initial Count to start with
func (hh *HOTP) ProvisioningUri(un, realm string, cnt int) string {
	return BuildURI(OtpTypeHotp, hh.secret, un, realm, hh.hasher.HashName, cnt, hh.nDigits, 0)
}
