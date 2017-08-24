package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

type testValues struct {
	dec, enc string // decoded hex value
}

var n = 3000000
var testPairs = make([]testValues, 0, n)

func initTestPairs() {
	if len(testPairs) > 0 {
		return
	}
	// pre-make the test pairs, so it doesn't take up benchmark time...
	data := make([]byte, 32)
	for i := 0; i < n; i++ {
		rand.Read(data)
		testPairs = append(testPairs, testValues{dec: hex.EncodeToString(data), enc: BitcoinEncoding.EncodeToString(data)})
	}
}

func TestEncodingAndDecoding(t *testing.T) {
	for j := 1; j < 256; j++ {
		var b = make([]byte, j)
		for i := 0; i < 100; i++ {
			rand.Read(b)
			te := radixEncoding(b)
			be := BitcoinEncoding.EncodeToString(b)

			if be != te {
				t.Errorf("encoding err: %#v", hex.EncodeToString(b))
			}

			bd, eerr := BitcoinEncoding.DecodeString(be)
			if eerr != nil {
				t.Errorf("fast encoding error: %v", eerr)
			}

			td, terr := radixDecoding(te)
			if terr != nil {
				t.Errorf("trivial error: %v", terr)
			}

			if hex.EncodeToString(bd) != hex.EncodeToString(td) {
				t.Errorf("encoding decoding err: [%x] %q != %q", b, hex.EncodeToString(bd), hex.EncodeToString(td))
			}
		}
	}
}

func BenchmarkTrivialBase58Encoding(b *testing.B) {
	b.ReportAllocs()

	data := make([]byte, 32)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Read(data)
		b.StartTimer()
		radixEncoding(data)
	}
}

func BenchmarkFastBase58Encoding(b *testing.B) {
	b.ReportAllocs()

	data := make([]byte, 32)
	for i := 0; i < b.N; i++ {
		rand.Read(data)
		BitcoinEncoding.EncodeToString(data)
	}
}

func BenchmarkTrivialBase58Decoding(b *testing.B) {
	b.ReportAllocs()

	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		radixDecoding(testPairs[i].enc)
	}
}

func BenchmarkFastBase58Decoding(b *testing.B) {
	b.ReportAllocs()

	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BitcoinEncoding.DecodeString(testPairs[i].enc)
	}
}

// Kept for benchmark comparisons

var (
	bn0  = big.NewInt(0)
	bn58 = big.NewInt(58)
)

var decodeMap [256]byte

func init() {
	for i := range decodeMap {
		decodeMap[i] = 0xFF
	}
	for i, b := range bitcoinAlphabet {
		decodeMap[b] = byte(i)
	}
}

// radixEncoding encodes to a base58 string using a big int radix
// which is slower than bit shifting. This function is used for
// benchmark comparisons
func radixEncoding(a []byte) string {
	idx := len(a)*138/100 + 1
	buf := make([]byte, idx)
	bn := new(big.Int).SetBytes(a)
	var mo *big.Int
	for bn.Cmp(bn0) != 0 {
		bn, mo = bn.DivMod(bn, bn58, new(big.Int))
		idx--
		buf[idx] = bitcoinAlphabet[mo.Int64()]
	}
	for i := range a {
		if a[i] != 0 {
			break
		}
		idx--
		buf[idx] = bitcoinAlphabet[0]
	}
	return string(buf[idx:])
}

// radixDecoding decodes a base58 string using a big int radix
// which is slower than bit shifting. This function is used for
// benchmark comparisons
func radixDecoding(str string) ([]byte, error) {
	src := []byte(str)
	var zcnt int
	for ; zcnt < len(src) && src[zcnt] == '1'; zcnt++ {
	}

	n := new(big.Int)
	for i := zcnt; i < len(src); i++ {
		b := decodeMap[src[i]]
		if b == 0xFF {
			return nil, fmt.Errorf("illegal base58 data at input byte %d ", i)
		}
		n.Mul(n, bn58)
		n.Add(n, big.NewInt(int64(b)))
	}

	dst := n.Bytes()
	buf := make([]byte, zcnt+len(dst))
	copy(buf[zcnt:], dst)

	return buf, nil
}
