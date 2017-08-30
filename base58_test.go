package base58

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

type addressTest struct {
	short, address, hexcode string
	big                     uint64
}

var base58FlickrPaths = []addressTest{
	{short: "brXijP", hexcode: "0198b9a10f", big: 6857269519},
	{short: "6hKMCS", hexcode: "cee93986", big: 3471391110},
	{short: "6hDrmR", hexcode: "ced65225", big: 3470152229},
	{short: "6hHHZB", hexcode: "cee31559", big: 3470988633},
	{short: "6hHKum", hexcode: "cee32900", big: 3470993664},
	{short: "6hLgFW", hexcode: "ceeaaa28", big: 3471485480},
	{short: "6hBRKR", hexcode: "ced19e6b", big: 3469844075},
	{short: "6hGRTd", hexcode: "cee082de", big: 3470820062},
	{short: "6hCuie", hexcode: "ced37e97", big: 3469966999},
	{short: "6hJuXN", hexcode: "cee56444", big: 3471139908},
	{short: "6hJsyS", hexcode: "cee544ca", big: 3471131850},
	{short: "6hFWFb", hexcode: "ceddc7b0", big: 3470641072},
	{short: "6hENdZ", hexcode: "ceda5e79", big: 3470417529},
	{short: "6hEJqg", hexcode: "ceda2c77", big: 3470404727},
	{short: "6hGNaq", hexcode: "cee051fa", big: 3470807546},
	{short: "6hDRoZ", hexcode: "ced78e01", big: 3470233089},
	{short: "6hKkP9", hexcode: "cee7e632", big: 3471304242},
	{short: "6hHVZ3", hexcode: "cee3b2e8", big: 3471028968},
	{short: "6hNcfE", hexcode: "cef0642e", big: 3471860782},
	{short: "6hJBqs", hexcode: "cee5b926", big: 3471161638},
	{short: "6hCPyc", hexcode: "ced47ba7", big: 3470031783},
	{short: "6hJNrC", hexcode: "cee649f6", big: 3471198710},
	{short: "6hKmkd", hexcode: "cee7ed02", big: 3471305986},
	{short: "6hFUYs", hexcode: "ceddb152", big: 3470635346},
	{short: "6hK6UC", hexcode: "cee72f78", big: 3471257464},
	{short: "6hBmiv", hexcode: "ced01b5f", big: 3469744991},
	{short: "6hKex1", hexcode: "cee793b2", big: 3471283122},
	{short: "6hFHQj", hexcode: "cedd1eee", big: 3470597870},
	{short: "6hCA2n", hexcode: "ced3c9d7", big: 3469986263},
	{short: "6hBTgt", hexcode: "ced1b245", big: 3469849157},
	{short: "6hHEss", hexcode: "cee2e6de", big: 3470976734},
	{short: "6hLows", hexcode: "ceeb03fe", big: 3471508478},
	{short: "6hD95z", hexcode: "ced56f11", big: 3470094097},
	{short: "6hKjcq", hexcode: "cee7d0f6", big: 3471298806},
	{short: "6hGEbd", hexcode: "cedfe908", big: 3470780680},
	{short: "6hKSNS", hexcode: "cee97d7e", big: 3471408510},
	{short: "6hG8hv", hexcode: "cede5319", big: 3470676761},
	{short: "6hEmj6", hexcode: "ced909f9", big: 3470330361},
	{short: "6hGjpn", hexcode: "cedee533", big: 3470714163},
	{short: "6hEsUr", hexcode: "ced96099", big: 3470352537},
	{short: "6hJEhy", hexcode: "cee5dec8", big: 3471171272},
	{short: "6hKBHn", hexcode: "cee8b723", big: 3471357731},
	{short: "6hG3gi", hexcode: "cede111f", big: 3470659871},
	{short: "6hFJTT", hexcode: "cedd2ce1", big: 3470601441},
	{short: "6hLZDs", hexcode: "ceecd180", big: 3471626624},
	{short: "6hGdL7", hexcode: "cede9b0e", big: 3470695182},
	{short: "6hBpi4", hexcode: "ced042b1", big: 3469755057},
	{short: "6hEuFV", hexcode: "ced9780b", big: 3470358539},
	{short: "6hGVw1", hexcode: "cee0b2a0", big: 3470832288},
	{short: "6hLdm1", hexcode: "ceea7e38", big: 3471474232},
	{short: "6hFcCK", hexcode: "cedb9217", big: 3470496279},
	{short: "6hDZmR", hexcode: "ced7f6a5", big: 3470259877},
	{short: "6hG8iX", hexcode: "cede536d", big: 3470676845},
	{short: "6hFZZL", hexcode: "ceddf352", big: 3470652242},
	{short: "6hJ79u", hexcode: "cee43874", big: 3471063156},
	{short: "6hMsrS", hexcode: "ceee31ac", big: 3471716780},
	{short: "6hGH3G", hexcode: "cee00ec0", big: 3470790336},
	{short: "6hKqD3", hexcode: "cee8259c", big: 3471320476},
	{short: "6hKxEY", hexcode: "cee88208", big: 3471344136},
	{short: "6hHVF1", hexcode: "cee3aed2", big: 3471027922},
}

var base58BitcoinAddresses = []addressTest{
	{hexcode: "0065a16059864a2fdbc7c99a4723a8395bc6f188eb", address: "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i"},
	{hexcode: "0574f209f6ea907e2ea48f74fae05782ae8a665257", address: "3CMNFxN1oHBc4R1EpboAL5yzHGgE611Xou"},
	{hexcode: "6f53c0307d6851aa0ce7825ba883c6bd9ad242b486", address: "mo9ncXisMeAoXwqcV5EWuyncbmCcQN4rVs"},
	{hexcode: "c46349a418fc4578d10a372b54b45c280cc8c4382f", address: "2N2JD6wb56AfK4tfmM6PwdVmoYk2dCKf4Br"},
	{hexcode: "80eddbdc1168f1daeadbd3e44c1e3f8f5a284c2029f78ad26af98583a499de5b19", address: "5Kd3NBUAdUnhyzenEwVLy9pBKxSwXvE9FMPyR4UKZvpe6E3AgLr"},
	{hexcode: "8055c9bccb9ed68446d1b75273bbce89d7fe013a8acd1625514420fb2aca1a21c401", address: "Kz6UJmQACJmLtaQj5A3JAge4kVTNQ8gbvXuwbmCj7bsaabudb3RD"},
	{hexcode: "ef36cb93b9ab1bdabf7fb9f2c04f1b9cc879933530ae7842398eef5a63a56800c2", address: "9213qJab2HNEpMpYNBa7wHGFKKbkDn24jpANDs2huN3yi4J11ko"},
	{hexcode: "efb9f4892c9e8282028fea1d2667c4dc5213564d41fc5783896a0d843fc15089f301", address: "cTpB4YiyKiBcPxnefsDpbnDxFDffjqJob8wGCEDXxgQ7zQoMXJdH"},
	{hexcode: "006d23156cbbdcc82a5a47eee4c2c7c583c18b6bf4", address: "1Ax4gZtb7gAit2TivwejZHYtNNLT18PUXJ"},
	{hexcode: "05fcc5460dd6e2487c7d75b1963625da0e8f4c5975", address: "3QjYXhTkvuj8qPaXHTTWb5wjXhdsLAAWVy"},
	{hexcode: "6ff1d470f9b02370fdec2e6b708b08ac431bf7a5f7", address: "n3ZddxzLvAY9o7184TB4c6FJasAybsw4HZ"},
	{hexcode: "c4c579342c2c4c9220205e2cdc285617040c924a0a", address: "2NBFNJTktNa7GZusGbDbGKRZTxdK9VVez3n"},
	{hexcode: "80a326b95ebae30164217d7a7f57d72ab2b54e3be64928a19da0210b9568d4015e", address: "5K494XZwps2bGyeL71pWid4noiSNA2cfCibrvRWqcHSptoFn7rc"},
	{hexcode: "807d998b45c219a1e38e99e7cbd312ef67f77a455a9b50c730c27f02c6f730dfb401", address: "L1RrrnXkcKut5DEMwtDthjwRcTTwED36thyL1DebVrKuwvohjMNi"},
	{hexcode: "efd6bca256b5abc5602ec2e1c121a08b0da2556587430bcf7e1898af2224885203", address: "93DVKyFYwSN6wEo3E2fCrFPUp17FtrtNi2Lf7n4G3garFb16CRj"},
	{hexcode: "efa81ca4e8f90181ec4b61b6a7eb998af17b2cb04de8a03b504b9e34c4c61db7d901", address: "cTDVKtMGVYWTHCb1AFjmVbEbWjvKpKqKgMaR3QJxToMSQAhmCeTN"},
	{hexcode: "007987ccaa53d02c8873487ef919677cd3db7a6912", address: "1C5bSj1iEGUgSTbziymG7Cn18ENQuT36vv"},
	{hexcode: "0563bcc565f9e68ee0189dd5cc67f1b0e5f02f45cb", address: "3AnNxabYGoTxYiTEZwFEnerUoeFXK2Zoks"},
	{hexcode: "6fef66444b5b17f14e8fae6e7e19b045a78c54fd79", address: "n3LnJXCqbPjghuVs8ph9CYsAe4Sh4j97wk"},
	{hexcode: "c4c3e55fceceaa4391ed2a9677f4a4d34eacd021a0", address: "2NB72XtkjpnATMggui83aEtPawyyKvnbX2o"},
	{hexcode: "80e75d936d56377f432f404aabb406601f892fd49da90eb6ac558a733c93b47252", address: "5KaBW9vNtWNhc3ZEDyNCiXLPdVPHCikRxSBWwV9NrpLLa4LsXi9"},
	{hexcode: "808248bd0375f2f75d7e274ae544fb920f51784480866b102384190b1addfbaa5c01", address: "L1axzbSyynNYA8mCAhzxkipKkfHtAXYF4YQnhSKcLV8YXA874fgT"},
	{hexcode: "ef44c4f6a096eac5238291a94cc24c01e3b19b8d8cef72874a079e00a242237a52", address: "927CnUkUbasYtDwYwVn2j8GdTuACNnKkjZ1rpZd2yBB1CLcnXpo"},
	{hexcode: "efd1de707020a9059d6d3abaf85e17967c6555151143db13dbb06db78df0f15c6901", address: "cUcfCMRjiQf85YMzzQEk9d1s5A4K7xL5SmBCLrezqXFuTVefyhY7"},
	{hexcode: "00adc1cc2081a27206fae25792f28bbc55b831549d", address: "1Gqk4Tv79P91Cc1STQtU3s1W6277M2CVWu"},
	{hexcode: "05188f91a931947eddd7432d6e614387e32b244709", address: "33vt8ViH5jsr115AGkW6cEmEz9MpvJSwDk"},
	{hexcode: "6f1694f5bc1a7295b600f40018a618a6ea48eeb498", address: "mhaMcBxNh5cqXm4aTQ6EcVbKtfL6LGyK2H"},
	{hexcode: "c43b9b3fd7a50d4f08d1a5b0f62f644fa7115ae2f3", address: "2MxgPqX1iThW3oZVk9KoFcE5M4JpiETssVN"},
	{hexcode: "80091035445ef105fa1bb125eccfb1882f3fe69592265956ade751fd095033d8d0", address: "5HtH6GdcwCJA4ggWEL1B3jzBBUB8HPiBi9SBc5h9i4Wk4PSeApR"},
	{hexcode: "80ab2b4bcdfc91d34dee0ae2a8c6b6668dadaeb3a88b9859743156f462325187af01", address: "L2xSYmMeVo3Zek3ZTsv9xUrXVAmrWxJ8Ua4cw8pkfbQhcEFhkXT8"},
	{hexcode: "efb4204389cef18bbe2b353623cbf93e8678fbc92a475b664ae98ed594e6cf0856", address: "92xFEve1Z9N8Z641KQQS7ByCSb8kGjsDzw6fAmjHN1LZGKQXyMq"},
	{hexcode: "efe7b230133f1b5489843260236b06edca25f66adb1be455fbd38d4010d48faeef01", address: "cVM65tdYu1YK37tNoAyGoJTR13VBYFva1vg9FLuPAsJijGvG6NEA"},
	{hexcode: "00c4c1b72491ede1eedaca00618407ee0b772cad0d", address: "1JwMWBVLtiqtscbaRHai4pqHokhFCbtoB4"},
	{hexcode: "05f6fe69bcb548a829cce4c57bf6fff8af3a5981f9", address: "3QCzvfL4ZRvmJFiWWBVwxfdaNBT8EtxB5y"},
	{hexcode: "6f261f83568a098a8638844bd7aeca039d5f2352c0", address: "mizXiucXRCsEriQCHUkCqef9ph9qtPbZZ6"},
	{hexcode: "c4e930e1834a4d234702773951d627cce82fbb5d2e", address: "2NEWDzHWwY5ZZp8CQWbB7ouNMLqCia6YRda"},
	{hexcode: "80d1fab7ab7385ad26872237f1eb9789aa25cc986bacc695e07ac571d6cdac8bc0", address: "5KQmDryMNDcisTzRp3zEq9e4awRmJrEVU1j5vFRTKpRNYPqYrMg"},
	{hexcode: "80b0bbede33ef254e8376aceb1510253fc3550efd0fcf84dcd0c9998b288f166b301", address: "L39Fy7AC2Hhj95gh3Yb2AU5YHh1mQSAHgpNixvm27poizcJyLtUi"},
	{hexcode: "ef037f4192c630f399d9271e26c575269b1d15be553ea1a7217f0cb8513cef41cb", address: "91cTVUcgydqyZLgaANpf1fvL55FH53QMm4BsnCADVNYuWuqdVys"},
	{hexcode: "ef6251e205e8ad508bab5596bee086ef16cd4b239e0cc0c5d7c4e6035441e7d5de01", address: "cQspfSzsgLeiJGB2u8vrAiWpCU4MxUT6JseWo2SjXy4Qbzn2fwDw"},
	{hexcode: "005eadaf9bb7121f0f192561a5a62f5e5f54210292", address: "19dcawoKcZdQz365WpXWMhX6QCUpR9SY4r"},
	{hexcode: "053f210e7277c899c3a155cc1c90f4106cbddeec6e", address: "37Sp6Rv3y4kVd1nQ1JV5pfqXccHNyZm1x3"},
	{hexcode: "6fc8a3c2a09a298592c3e180f02487cd91ba3400b5", address: "myoqcgYiehufrsnnkqdqbp69dddVDMopJu"},
	{hexcode: "c499b31df7c9068d1481b596578ddbb4d3bd90baeb", address: "2N7FuwuUuoTBrDFdrAZ9KxBmtqMLxce9i1C"},
	{hexcode: "80c7666842503db6dc6ea061f092cfb9c388448629a6fe868d068c42a488b478ae", address: "5KL6zEaMtPRXZKo1bbMq7JDjjo1bJuQcsgL33je3oY8uSJCR5b4"},
	{hexcode: "8007f0803fc5399e773555ab1e8939907e9badacc17ca129e67a2f5f2ff84351dd01", address: "KwV9KAfwbwt51veZWNscRTeZs9CKpojyu1MsPnaKTF5kz69H1UN2"},
	{hexcode: "efea577acfb5d1d14d3b7b195c321566f12f87d2b77ea3a53f68df7ebf8604a801", address: "93N87D6uxSBzwXvpokpzg8FFmfQPmvX4xHoWQe3pLdYpbiwT5YV"},
	{hexcode: "ef0b3b34f0958d8a268193a9814da92c3e8b58b4a4378a542863e34ac289cd830c01", address: "cMxXusSihaX58wpJ3tNuuUcZEQGt6DKJ1wEpxys88FFaQCYjku9h"},
	{hexcode: "001ed467017f043e91ed4c44b4e8dd674db211c4e6", address: "13p1ijLwsnrcuyqcTvJXkq2ASdXqcnEBLE"},
	{hexcode: "055ece0cadddc415b1980f001785947120acdb36fc", address: "3ALJH9Y951VCGcVZYAdpA3KchoP9McEj1G"},
	{hexcode: "00", address: "0"},
	{hexcode: "00000000", address: "0"},
}

func TestBitcoinEncodingCheck(t *testing.T) {
	for _, ah := range base58BitcoinAddresses {
		b, err := hex.DecodeString(ah.hexcode)
		if err != nil {
			t.Errorf("decoding hex: [%s] %v", ah.hexcode, err)
		}
		have := BitcoinEncoding.EncodeToString(b)
		if have != ah.address {
			t.Errorf("want: %s have %s", ah.address, have)
		}
	}
}

func TestFlickrEncodingCheck(t *testing.T) {
	for _, ah := range base58FlickrPaths {
		b, err := hex.DecodeString(ah.hexcode)
		if err != nil {
			t.Errorf("decoding hex: [%s] %v", ah.hexcode, err)
		}
		have := FlickrEncoding.EncodeToString(b)
		if have != ah.short {
			t.Errorf("want: %s have %s", ah.short, have)
		}
	}
}

func TestEncodingCheck(t *testing.T) {
	for _, ah := range base58BitcoinAddresses {
		b, err := hex.DecodeString(ah.hexcode)
		if err != nil {
			t.Errorf("decoding hex: [%s] %v", ah.hexcode, err)
		}
		have := BitcoinEncoding.EncodeToString(b)
		if have != ah.address {
			t.Errorf("want: %s have %s", ah.address, have)
		}
	}
}

func TestBitcoinDecodingCheck(t *testing.T) {
	for _, ah := range base58BitcoinAddresses {
		have, err := BitcoinEncoding.DecodeString(ah.address)
		if ah.address == "0" && err == nil {
			t.Errorf("decoding address: [%s] should have error", ah.address)
		} else if ah.address == "0" {
			continue // skip test
		}
		if err != nil {
			t.Errorf("decoding address: [%s] %v", ah.address, err)
		}
		if hex.EncodeToString(have) != ah.hexcode {
			t.Errorf("want: %s have %x", ah.hexcode, have)
		}
	}
}

func TestFlickrDecodingCheck(t *testing.T) {
	for _, ah := range base58FlickrPaths {
		have, err := FlickrEncoding.DecodeString(ah.short)
		if ah.short == "0" && err == nil {
			t.Errorf("decoding address: [%s] should have error", ah.address)
		} else if ah.short == "0" {
			continue // skip test
		}
		if err != nil {
			t.Errorf("decoding address: [%s] %v", ah.short, err)
		}
		if hex.EncodeToString(have) != ah.hexcode {
			t.Errorf("want: %s have %x", ah.hexcode, have)
		}
	}
}

func TestEncodingAndDecodingEquality(t *testing.T) {
	for j := 1; j < 256; j++ {
		var b = make([]byte, j)
		for i := 0; i < 100; i++ {
			rand.Read(b)

			// check if all 0's skip this special case
			if hex.EncodeToString(b) == hex.EncodeToString(make([]byte, len(b))) {
				continue
			}

			te := radixEncoding(b)
			be := StdEncoding.EncodeToString(b)

			if be != te {
				t.Errorf("encoding err: %#v", hex.EncodeToString(b))
			}

			bd, eerr := StdEncoding.DecodeString(be)
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
		rand.Read(data)
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
