package main

import (
	"fmt"
)

const bitcoinAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// BitcoinEncoding is the standard base58 encoding, for Bitcoin
var BitcoinEncoding = NewEncoding(bitcoinAlphabet, With1Padding())

// An Encoding is a radix 58 encoding/decoding scheme, defined by a
// 58-character alphabet. The most common encoding is the "base58"
// check encoding for bitcoin
type Encoding struct {
	encode    string
	decodeMap [256]int8

	pad1 bool
}

// opts is the functional option type
type opts func(*Encoding)

// With1Padding adds a 1 at the beginning for each 0-padded prefix
// of the base58 address
func With1Padding() func(*Encoding) {
	return func(e *Encoding) {
		e.pad1 = true
	}
}

// NewEncoding returns a new Encoding defined by the given alphabet,
// which must be a 58-byte string.
func NewEncoding(encoder string, options ...opts) *Encoding {
	if len(encoder) != 58 {
		panic("encoding alphabet is not 58-bytes long")
	}

	e := new(Encoding)
	e.encode = encoder
	for i := 0; i < len(e.decodeMap); i++ {
		e.decodeMap[i] = -1
	}
	for i := 0; i < len(encoder); i++ {
		e.decodeMap[encoder[i]] = int8(i)
	}

	for _, opt := range options {
		opt(e)
	}

	return e
}

// Encode encodes src using the encoding enc, writing
// EncodedLen(len(src)) bytes to dst.
func (enc *Encoding) Encode(dst, src []byte) (n int) {
	binsz := len(src)
	var i, j, high, zcount, carry int

	if enc.pad1 {
		for zcount < binsz && src[zcount] == 0 {
			zcount++
		}
	}

	size := enc.EncodedLen(binsz - zcount)
	var buf = make([]byte, size)

	high = size - 1
	for i = zcount; i < binsz; i++ {
		j = size - 1
		for carry = int(src[i]); j > high || carry != 0; j-- {
			carry = carry + 256*int(buf[j])
			buf[j] = byte(carry % 58)
			carry /= 58
		}
		high = j
	}

	if enc.pad1 {
		for j = 0; j < size && buf[j] == 0; j++ {
		}
	}

	n = size - j + zcount
	if enc.pad1 && zcount != 0 {
		for i = 0; i < zcount; i++ {
			dst[i] = '1'
		}
	}

	for i = zcount; j < size; i++ {
		dst[i] = enc.encode[buf[j]]
		j++
	}

	return n
}

// EncodeToString returns the base58 encoding of src.
func (enc *Encoding) EncodeToString(src []byte) string {
	var buf = make([]byte, enc.EncodedLen(len(src)))
	n := enc.Encode(buf, src)
	return string(buf[:n])
}

// EncodedLen returns the length in bytes of the base58 encoding
// of an input buffer of length n.
func (enc *Encoding) EncodedLen(i int) int {
	return i*138/100 + 1
}

// Decode decodes src using the encoding enc. It writes at most
// DecodedLen(len(src)) bytes to dst and returns the number of bytes
// written. If src contains invalid base58 data, it will return the
// number of bytes successfully written and an error.
func (enc *Encoding) Decode(dst, src []byte) (n int, err error) {
	if len(src) == 0 {
		return n, fmt.Errorf("zero length string")
	}

	var size = len(src)

	var zmask uint32
	bytesleft := size % 4
	if bytesleft > 0 {
		zmask = (0xffffffff << uint32(bytesleft*8))
	}

	var zcount int
	var buf = make([]uint32, (size+3)/4)
	if enc.pad1 {
		for ; zcount < size && src[zcount] == '1'; zcount++ {
		}
	}

	for i := zcount; i < size; i++ {
		if src[i]&0x80 != 0 {
			return n, fmt.Errorf("High-bit set on invalid digit")
		}

		if enc.decodeMap[src[i]] == -1 {
			return n, fmt.Errorf("Invalid base58 digit (%q)", src[i])
		}

		c := uint32(enc.decodeMap[src[i]])

		for j := len(buf) - 1; j >= 0; j-- {
			t := uint64(buf[j])*58 + uint64(c)
			c = uint32((t & 0x3f00000000) >> 32)
			buf[j] = uint32(t)
		}

		if c > 0 {
			return n, fmt.Errorf("Output number too big (carry to the next int32)")
		}

		if buf[0]&zmask != 0 {
			return n, fmt.Errorf("Output number too big (last int32 filled too far)")
		}
	}

	n = zcount
	var mark bool
	for j := 0; j < len(buf); j++ {
		for k, mask := range []uint32{0x18, 0x10, 0x8, 0x0} {
			if j == 0 && bytesleft > 0 && k < 4-bytesleft {
				continue // skip the first bytes left over
			}
			if n > len(dst)-1 {
				return n - 1, fmt.Errorf("Unexpected end")
			}
			dst[n] = byte(buf[j] >> mask)
			if !mark && dst[n] == 0 {
				continue
			} else if !mark && dst[n] > 0 {
				mark = true
			}
			n++
		}
	}

	return n, nil
}

// DecodeString returns the bytes represented by the base58 string str.
func (enc *Encoding) DecodeString(str string) ([]byte, error) {
	var zcount int
	if enc.pad1 {
		for ; zcount < len(str) && str[zcount] == '1'; zcount++ {
		}
	}
	buf := make([]byte, enc.DecodedLen(len(str))+zcount)
	n, err := enc.Decode(buf, []byte(str))
	return buf[:n], err
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base58-encoded data.
func (enc *Encoding) DecodedLen(n int) int {
	return n
}
