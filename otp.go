package htotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"hash"
	"math"
)

// From Copyright (c) 2018 xlzd, https://github.com/xlzd/gotp, MIT Licensed.
type Hasher struct {
	HashName string
	Digest   func() hash.Hash
}

// From Copyright (c) 2018 xlzd, https://github.com/xlzd/gotp, MIT Licensed.
type OTP struct {
	secret  string  // secret in base32 format
	nDigits int     // number of integers in the OTP. Some apps expect this to be 6 nDigits, others support more.
	hasher  *Hasher // digest function to use in the HMAC (expected to be sha1)
}

// From Copyright (c) 2018 xlzd, https://github.com/xlzd/gotp, MIT Licensed.
func NewOTP(secret string, nDigits int, hasher *Hasher) OTP {
	if hasher == nil {
		hasher = &Hasher{
			HashName: "sha1",
			Digest:   sha1.New,
		}
	}
	return OTP{
		secret:  secret,
		nDigits: nDigits,
		hasher:  hasher,
	}
}

func (oo *OTP) GenerateOTP(input uint) (string, error) {
	byteSec, err := oo.ConvertSecretToBytes()
	if err != nil {
		return "", err
	}
	hasher := hmac.New(oo.hasher.Digest, byteSec)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(input)) // Potential Protability Problem
	hasher.Write(bs)
	hmacHash := hasher.Sum(nil)

	whereInHash := int(hmacHash[len(hmacHash)-1] & 0xf)
	code := uint(binary.BigEndian.Uint32(hmacHash[whereInHash:whereInHash+4]) & 0x7fffffff)
	code %= uint(math.Pow10(oo.nDigits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", oo.nDigits), code), nil
}

func (oo *OTP) GenerateOTPTL(input, tr uint) (string, uint, error) {
	byteSec, err := oo.ConvertSecretToBytes()
	if err != nil {
		return "", 0, err
	}
	hasher := hmac.New(oo.hasher.Digest, byteSec)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(input)) // Potential Protability Problem
	hasher.Write(bs)
	hmacHash := hasher.Sum(nil)

	whereInHash := int(hmacHash[len(hmacHash)-1] & 0xf)
	code := uint(binary.BigEndian.Uint32(hmacHash[whereInHash:whereInHash+4]) & 0x7fffffff)
	code %= uint(math.Pow10(oo.nDigits))
	return fmt.Sprintf(fmt.Sprintf("%%0%dd", oo.nDigits), code), tr, nil
}

// From Copyright (c) 2018 xlzd, https://github.com/xlzd/gotp, MIT Licensed.
func (oo *OTP) ConvertSecretToBytes() (bytes []byte, err error) {
	missingPadding := len(oo.secret) % 8
	if missingPadding != 0 {
		oo.secret = PadRight(oo.secret, "=", len(oo.secret)+missingPadding)
	}
	bytes, e0 := base32.StdEncoding.DecodeString(oo.secret)
	if e0 != nil {
		err = fmt.Errorf("Failed to encode (base32.StdEncoding) the secret: error: %s", e0)
	}
	return
}
