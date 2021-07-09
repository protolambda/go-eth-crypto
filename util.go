package crypto

import (
	"encoding/hex"
	"math/big"
	"strings"
)

const AddressLength = 20

type Address = [AddressLength]byte

var big0 = big.NewInt(0)
var big1 = big.NewInt(1)

// paddedBigBytes encodes a big integer as a big-endian byte slice. The length
// of the slice is at least n bytes.
func paddedBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	readBits(bigint, ret)
	return ret
}

const (
	// number of bits in a big.Word
	wordBits = 32 << (uint64(^big.Word(0)) >> 63)
	// number of bytes in a big.Word
	wordBytes = wordBits / 8
)

// readBits encodes the absolute value of bigint as big-endian bytes. Callers must ensure
// that buf has enough space. If buf is too short the result will be incomplete.
func readBits(bigint *big.Int, buf []byte) {
	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

// copyBytes returns an exact copy of the provided bytes.
func copyBytes(b []byte) (copiedBytes []byte) {
	if b == nil {
		return nil
	}
	copiedBytes = make([]byte, len(b))
	copy(copiedBytes, b)

	return
}

func mustDecodeHex(input string) []byte {
	if strings.HasPrefix(input, "0x") {
		input = input[2:]
	}
	b, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return b
}

func mustDecodeBig(input string) *big.Int {
	if strings.HasPrefix(input, "0x") {
		input = input[2:]
	}
	v, ok := new(big.Int).SetString(input, 16)
	if !ok {
		panic("bad big hex input")
	}
	return v
}

func mustParseBig256(input string) *big.Int {
	if strings.HasPrefix(input, "0x") {
		input = input[2:]
	}
	v, ok := new(big.Int).SetString(input, 16)
	if !ok || v.BitLen() > 256 {
		panic("bad big256 input")
	}
	return v
}

func hexToAddress(input string) Address {
	b := mustDecodeHex(input)
	if len(b) != AddressLength {
		panic("bad address length")
	}
	var addr Address
	copy(addr[:], b)
	return addr
}

func bytesToAddress(b []byte) Address {
	if len(b) > AddressLength {
		b = b[len(b)-AddressLength:]
	}
	var a Address
	copy(a[AddressLength-len(b):], b)
	return a
}
