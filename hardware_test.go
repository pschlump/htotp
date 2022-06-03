package htotp

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// This test will fail if you do not have 64 bit integers or if your hardware is not BigEndian.
func Test_Hardware(t *testing.T) {
	var buffer bytes.Buffer
	// verify that the hardware will treat an integer in the correct order
	var fx = func(input int) {
		bs := make([]byte, 8)
		binary.BigEndian.PutUint64(bs, uint64(input)) // Potential Protability Problem
		buffer.Write(bs)
	}
	fx(123456789033)

	//                        0     1     2     3     4     5     6     7
	expect_bytes := []byte{0x00, 0x00, 0x00, 0x1c, 0xbe, 0x99, 0x1a, 0x29}
	got_bytes := buffer.Bytes()
	if len(got_bytes) != 8 {
		t.Errorf("Error: wrong length, expected 8, got %d\n", len(got_bytes))
	}
	for ii := 0; ii < 8; ii++ {
		if expect_bytes[ii] != got_bytes[ii] {
			t.Errorf("Error: on ii=%d byte, expected %x got %x\n", ii, expect_bytes[ii], got_bytes[ii])
		}
	}
	/*
		whereInHash := int(hmacHash[len(hmacHash)-1] & 0xf)
		code := ((uint(hmacHash[whereInHash]) & 0x7f) << 24) |
			((uint(hmacHash[whereInHash+1] & 0xff)) << 16) |
			((uint(hmacHash[whereInHash+2] & 0xff)) << 8) |
			(uint(hmacHash[whereInHash+3]) & 0xff)
		code2 := binary.BigEndian.Uint32(hmacHash[whereInHash:whereInHash+4]) & 0x7fffffff
		if code != uint(code2) {
			fmt.Printf("%s Bad code ->%x<- code2 ->%x<- %s\n", dbgo.ColorRed, code, code2, dbgo.ColorReset)
		}
	*/
}
