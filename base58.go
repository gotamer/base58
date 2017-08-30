// Package base58 implements very fast base58 encoding as used by Bitcoin and Flickr
package base58

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type errString string

func (e errString) Error() string {
	return string(e)
}

// ErrInvalidChecksum is the error return when the checksum does not match
const ErrInvalidChecksum = errString("the checksum is invalid")

const bitcoinAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const flickrAlphabet = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

var (
	decodeBlockSizes = [...]int{0, 0, 1, 2, 3, 3, 4, 5, 6, 6, 7, 8}
	encodeBlockSizes = [...]int{0, 2, 3, 5, 6, 7, 9, 10, 11}
)

// StdEncoding is the standard base58 encoding based on the bitcoin alphabet
var StdEncoding = NewEncoding(bitcoinAlphabet)

// BitcoinEncoding is the standard base58 encoding with a checksum
var BitcoinEncoding = NewEncoding(bitcoinAlphabet, WithChecksum(4))

// flickrEncoding is the standard base58 encoding with a checksum
var FlickrEncoding = NewEncoding(flickrAlphabet)

// An Encoding is a radix 58 encoding/decoding scheme, defined by a
// 58-character alphabet. The most common encoding is the "base58"
// check encoding for bitcoin
type Encoding struct {
	encode    string
	decodeMap [256]int8

	checkNum  int
	checkFunc func([]byte) []byte
}

// opts is the functional option type
type opts func(*Encoding)

func WithChecksum(i int) func(*Encoding) {
	return func(enc *Encoding) {
		enc.checkNum = i
	}
}

func WithChecksumFunc(fn func([]byte) []byte) func(*Encoding) {
	return func(enc *Encoding) {
		enc.checkFunc = fn
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
	e.checkFunc = func(b []byte) []byte {
		sh1, sh2 := sha256.New(), sha256.New()
		sh1.Write(b)
		sh2.Write(sh1.Sum(nil))
		return sh2.Sum(nil)
	}
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

	for zcount < binsz && src[zcount] == 0 {
		zcount++
	}

	if zcount == len(src) {
		copy(dst, []byte("0"))
		return 1
	}

	if enc.checkNum > 0 {
		checkSum := enc.checkFunc(src)
		src = append(src, checkSum[:enc.checkNum]...)
		binsz = len(src)
	}

	size := enc.EncodedLen(binsz-zcount) - enc.checkNum
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

	for j = 0; j < size && buf[j] == 0; j++ {
	}

	n = size - j + zcount
	if zcount != 0 {
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
	var buf = make([]byte, enc.EncodedLen(len(src))+enc.checkNum)
	n := enc.Encode(buf, src)
	return string(buf[:n])
}

// EncodedLen returns the length in bytes of the base58 encoding
// of an input buffer of length n.
func (enc *Encoding) EncodedLen(n int) int {
	return ((n / 8) * 11) + encodeBlockSizes[n%8] + enc.checkNum
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
	for ; zcount < size && src[zcount] == '1'; zcount++ {
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

	if enc.checkNum > 0 {
		if n < enc.checkNum {
			return n, fmt.Errorf("Invalid checksum length")
		}

		n -= enc.checkNum
		checkSum, dst := dst[n:], dst[:n]
		checkChecksum := enc.checkFunc(dst)
		if hex.EncodeToString(checkSum[:enc.checkNum]) != hex.EncodeToString(checkChecksum[:enc.checkNum]) {
			return n, ErrInvalidChecksum
		}
	}

	return n, nil
}

// DecodeString returns the bytes represented by the base58 string str.
func (enc *Encoding) DecodeString(str string) ([]byte, error) {
	var zcount int
	for ; zcount < len(str) && str[zcount] == '1'; zcount++ {
	}

	buf := make([]byte, enc.DecodedLen(len(str)))
	n, err := enc.Decode(buf, []byte(str))
	return buf[:n], err
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of base58-encoded data.
func (enc *Encoding) DecodedLen(n int) int {
	return (((n / 11) * 8) + decodeBlockSizes[n%11]) + 3
}
