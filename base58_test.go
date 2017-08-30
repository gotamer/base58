package base58

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

type testPair struct {
	String, Hex string
	Big         uint64
}

var testPairsN = 3000000
var testPairs = make([]testPair, 0, testPairsN)

func initTestPairs() {
	if len(testPairs) > 0 {
		return
	}
	// pre-make the test pairs, so it doesn't take up benchmark time...
	data := make([]byte, 32)
	for i := 0; i < testPairsN; i++ {
		rand.Read(data)
		testPairs = append(testPairs, testPair{String: BitcoinEncoding.EncodeToString(data)})
	}
}

// pulled from - https://github.com/jimeh/node-base58/blob/master/test/examples.js
var base58FlickrTestPairs = []testPair{
	{String: "6hKMCS", Hex: "cee93986", Big: 3471391110},
	{String: "6hDrmR", Hex: "ced65225", Big: 3470152229},
	{String: "6hHHZB", Hex: "cee31559", Big: 3470988633},
	{String: "6hHKum", Hex: "cee32900", Big: 3470993664},
	{String: "6hLgFW", Hex: "ceeaaa28", Big: 3471485480},
	{String: "6hBRKR", Hex: "ced19e6b", Big: 3469844075},
	{String: "6hGRTd", Hex: "cee082de", Big: 3470820062},
	{String: "6hCuie", Hex: "ced37e97", Big: 3469966999},
	{String: "6hJuXN", Hex: "cee56444", Big: 3471139908},
	{String: "6hJsyS", Hex: "cee544ca", Big: 3471131850},
	{String: "6hFWFb", Hex: "ceddc7b0", Big: 3470641072},
	{String: "6hENdZ", Hex: "ceda5e79", Big: 3470417529},
	{String: "6hEJqg", Hex: "ceda2c77", Big: 3470404727},
	{String: "6hGNaq", Hex: "cee051fa", Big: 3470807546},
	{String: "6hDRoZ", Hex: "ced78e01", Big: 3470233089},
	{String: "6hKkP9", Hex: "cee7e632", Big: 3471304242},
	{String: "6hHVZ3", Hex: "cee3b2e8", Big: 3471028968},
	{String: "6hNcfE", Hex: "cef0642e", Big: 3471860782},
	{String: "6hJBqs", Hex: "cee5b926", Big: 3471161638},
	{String: "6hCPyc", Hex: "ced47ba7", Big: 3470031783},
	{String: "6hJNrC", Hex: "cee649f6", Big: 3471198710},
	{String: "6hKmkd", Hex: "cee7ed02", Big: 3471305986},
	{String: "6hFUYs", Hex: "ceddb152", Big: 3470635346},
	{String: "6hK6UC", Hex: "cee72f78", Big: 3471257464},
	{String: "6hBmiv", Hex: "ced01b5f", Big: 3469744991},
	{String: "6hKex1", Hex: "cee793b2", Big: 3471283122},
	{String: "6hFHQj", Hex: "cedd1eee", Big: 3470597870},
	{String: "6hCA2n", Hex: "ced3c9d7", Big: 3469986263},
	{String: "6hBTgt", Hex: "ced1b245", Big: 3469849157},
	{String: "6hHEss", Hex: "cee2e6de", Big: 3470976734},
	{String: "6hLows", Hex: "ceeb03fe", Big: 3471508478},
	{String: "6hD95z", Hex: "ced56f11", Big: 3470094097},
	{String: "6hKjcq", Hex: "cee7d0f6", Big: 3471298806},
	{String: "6hGEbd", Hex: "cedfe908", Big: 3470780680},
	{String: "6hKSNS", Hex: "cee97d7e", Big: 3471408510},
	{String: "6hG8hv", Hex: "cede5319", Big: 3470676761},
	{String: "6hEmj6", Hex: "ced909f9", Big: 3470330361},
	{String: "6hGjpn", Hex: "cedee533", Big: 3470714163},
	{String: "6hEsUr", Hex: "ced96099", Big: 3470352537},
	{String: "6hJEhy", Hex: "cee5dec8", Big: 3471171272},
	{String: "6hKBHn", Hex: "cee8b723", Big: 3471357731},
	{String: "6hG3gi", Hex: "cede111f", Big: 3470659871},
	{String: "6hFJTT", Hex: "cedd2ce1", Big: 3470601441},
	{String: "6hLZDs", Hex: "ceecd180", Big: 3471626624},
	{String: "6hGdL7", Hex: "cede9b0e", Big: 3470695182},
	{String: "6hBpi4", Hex: "ced042b1", Big: 3469755057},
	{String: "6hEuFV", Hex: "ced9780b", Big: 3470358539},
	{String: "6hGVw1", Hex: "cee0b2a0", Big: 3470832288},
	{String: "6hLdm1", Hex: "ceea7e38", Big: 3471474232},
	{String: "6hFcCK", Hex: "cedb9217", Big: 3470496279},
	{String: "6hDZmR", Hex: "ced7f6a5", Big: 3470259877},
	{String: "6hG8iX", Hex: "cede536d", Big: 3470676845},
	{String: "6hFZZL", Hex: "ceddf352", Big: 3470652242},
	{String: "6hJ79u", Hex: "cee43874", Big: 3471063156},
	{String: "6hMsrS", Hex: "ceee31ac", Big: 3471716780},
	{String: "6hGH3G", Hex: "cee00ec0", Big: 3470790336},
	{String: "6hKqD3", Hex: "cee8259c", Big: 3471320476},
	{String: "6hKxEY", Hex: "cee88208", Big: 3471344136},
	{String: "6hHVF1", Hex: "cee3aed2", Big: 3471027922},
}

// pulled from - https://github.com/trezor/trezor-crypto/blob/master/test_check.c
var base58BitcoinTestPairs = []testPair{
	{String: "1AGNa15ZQXAZUgFiqJ2i7Z2DPU2J6hW62i", Hex: "0065a16059864a2fdbc7c99a4723a8395bc6f188eb"},
	{String: "3CMNFxN1oHBc4R1EpboAL5yzHGgE611Xou", Hex: "0574f209f6ea907e2ea48f74fae05782ae8a665257"},
	{String: "mo9ncXisMeAoXwqcV5EWuyncbmCcQN4rVs", Hex: "6f53c0307d6851aa0ce7825ba883c6bd9ad242b486"},
	{String: "2N2JD6wb56AfK4tfmM6PwdVmoYk2dCKf4Br", Hex: "c46349a418fc4578d10a372b54b45c280cc8c4382f"},
	{String: "5Kd3NBUAdUnhyzenEwVLy9pBKxSwXvE9FMPyR4UKZvpe6E3AgLr", Hex: "80eddbdc1168f1daeadbd3e44c1e3f8f5a284c2029f78ad26af98583a499de5b19"},
	{String: "Kz6UJmQACJmLtaQj5A3JAge4kVTNQ8gbvXuwbmCj7bsaabudb3RD", Hex: "8055c9bccb9ed68446d1b75273bbce89d7fe013a8acd1625514420fb2aca1a21c401"},
	{String: "9213qJab2HNEpMpYNBa7wHGFKKbkDn24jpANDs2huN3yi4J11ko", Hex: "ef36cb93b9ab1bdabf7fb9f2c04f1b9cc879933530ae7842398eef5a63a56800c2"},
	{String: "cTpB4YiyKiBcPxnefsDpbnDxFDffjqJob8wGCEDXxgQ7zQoMXJdH", Hex: "efb9f4892c9e8282028fea1d2667c4dc5213564d41fc5783896a0d843fc15089f301"},
	{String: "1Ax4gZtb7gAit2TivwejZHYtNNLT18PUXJ", Hex: "006d23156cbbdcc82a5a47eee4c2c7c583c18b6bf4"},
	{String: "3QjYXhTkvuj8qPaXHTTWb5wjXhdsLAAWVy", Hex: "05fcc5460dd6e2487c7d75b1963625da0e8f4c5975"},
	{String: "n3ZddxzLvAY9o7184TB4c6FJasAybsw4HZ", Hex: "6ff1d470f9b02370fdec2e6b708b08ac431bf7a5f7"},
	{String: "2NBFNJTktNa7GZusGbDbGKRZTxdK9VVez3n", Hex: "c4c579342c2c4c9220205e2cdc285617040c924a0a"},
	{String: "5K494XZwps2bGyeL71pWid4noiSNA2cfCibrvRWqcHSptoFn7rc", Hex: "80a326b95ebae30164217d7a7f57d72ab2b54e3be64928a19da0210b9568d4015e"},
	{String: "L1RrrnXkcKut5DEMwtDthjwRcTTwED36thyL1DebVrKuwvohjMNi", Hex: "807d998b45c219a1e38e99e7cbd312ef67f77a455a9b50c730c27f02c6f730dfb401"},
	{String: "93DVKyFYwSN6wEo3E2fCrFPUp17FtrtNi2Lf7n4G3garFb16CRj", Hex: "efd6bca256b5abc5602ec2e1c121a08b0da2556587430bcf7e1898af2224885203"},
	{String: "cTDVKtMGVYWTHCb1AFjmVbEbWjvKpKqKgMaR3QJxToMSQAhmCeTN", Hex: "efa81ca4e8f90181ec4b61b6a7eb998af17b2cb04de8a03b504b9e34c4c61db7d901"},
	{String: "1C5bSj1iEGUgSTbziymG7Cn18ENQuT36vv", Hex: "007987ccaa53d02c8873487ef919677cd3db7a6912"},
	{String: "3AnNxabYGoTxYiTEZwFEnerUoeFXK2Zoks", Hex: "0563bcc565f9e68ee0189dd5cc67f1b0e5f02f45cb"},
	{String: "n3LnJXCqbPjghuVs8ph9CYsAe4Sh4j97wk", Hex: "6fef66444b5b17f14e8fae6e7e19b045a78c54fd79"},
	{String: "2NB72XtkjpnATMggui83aEtPawyyKvnbX2o", Hex: "c4c3e55fceceaa4391ed2a9677f4a4d34eacd021a0"},
	{String: "5KaBW9vNtWNhc3ZEDyNCiXLPdVPHCikRxSBWwV9NrpLLa4LsXi9", Hex: "80e75d936d56377f432f404aabb406601f892fd49da90eb6ac558a733c93b47252"},
	{String: "L1axzbSyynNYA8mCAhzxkipKkfHtAXYF4YQnhSKcLV8YXA874fgT", Hex: "808248bd0375f2f75d7e274ae544fb920f51784480866b102384190b1addfbaa5c01"},
	{String: "927CnUkUbasYtDwYwVn2j8GdTuACNnKkjZ1rpZd2yBB1CLcnXpo", Hex: "ef44c4f6a096eac5238291a94cc24c01e3b19b8d8cef72874a079e00a242237a52"},
	{String: "cUcfCMRjiQf85YMzzQEk9d1s5A4K7xL5SmBCLrezqXFuTVefyhY7", Hex: "efd1de707020a9059d6d3abaf85e17967c6555151143db13dbb06db78df0f15c6901"},
	{String: "1Gqk4Tv79P91Cc1STQtU3s1W6277M2CVWu", Hex: "00adc1cc2081a27206fae25792f28bbc55b831549d"},
	{String: "33vt8ViH5jsr115AGkW6cEmEz9MpvJSwDk", Hex: "05188f91a931947eddd7432d6e614387e32b244709"},
	{String: "mhaMcBxNh5cqXm4aTQ6EcVbKtfL6LGyK2H", Hex: "6f1694f5bc1a7295b600f40018a618a6ea48eeb498"},
	{String: "2MxgPqX1iThW3oZVk9KoFcE5M4JpiETssVN", Hex: "c43b9b3fd7a50d4f08d1a5b0f62f644fa7115ae2f3"},
	{String: "5HtH6GdcwCJA4ggWEL1B3jzBBUB8HPiBi9SBc5h9i4Wk4PSeApR", Hex: "80091035445ef105fa1bb125eccfb1882f3fe69592265956ade751fd095033d8d0"},
	{String: "L2xSYmMeVo3Zek3ZTsv9xUrXVAmrWxJ8Ua4cw8pkfbQhcEFhkXT8", Hex: "80ab2b4bcdfc91d34dee0ae2a8c6b6668dadaeb3a88b9859743156f462325187af01"},
	{String: "92xFEve1Z9N8Z641KQQS7ByCSb8kGjsDzw6fAmjHN1LZGKQXyMq", Hex: "efb4204389cef18bbe2b353623cbf93e8678fbc92a475b664ae98ed594e6cf0856"},
	{String: "cVM65tdYu1YK37tNoAyGoJTR13VBYFva1vg9FLuPAsJijGvG6NEA", Hex: "efe7b230133f1b5489843260236b06edca25f66adb1be455fbd38d4010d48faeef01"},
	{String: "1JwMWBVLtiqtscbaRHai4pqHokhFCbtoB4", Hex: "00c4c1b72491ede1eedaca00618407ee0b772cad0d"},
	{String: "3QCzvfL4ZRvmJFiWWBVwxfdaNBT8EtxB5y", Hex: "05f6fe69bcb548a829cce4c57bf6fff8af3a5981f9"},
	{String: "mizXiucXRCsEriQCHUkCqef9ph9qtPbZZ6", Hex: "6f261f83568a098a8638844bd7aeca039d5f2352c0"},
	{String: "2NEWDzHWwY5ZZp8CQWbB7ouNMLqCia6YRda", Hex: "c4e930e1834a4d234702773951d627cce82fbb5d2e"},
	{String: "5KQmDryMNDcisTzRp3zEq9e4awRmJrEVU1j5vFRTKpRNYPqYrMg", Hex: "80d1fab7ab7385ad26872237f1eb9789aa25cc986bacc695e07ac571d6cdac8bc0"},
	{String: "L39Fy7AC2Hhj95gh3Yb2AU5YHh1mQSAHgpNixvm27poizcJyLtUi", Hex: "80b0bbede33ef254e8376aceb1510253fc3550efd0fcf84dcd0c9998b288f166b301"},
	{String: "91cTVUcgydqyZLgaANpf1fvL55FH53QMm4BsnCADVNYuWuqdVys", Hex: "ef037f4192c630f399d9271e26c575269b1d15be553ea1a7217f0cb8513cef41cb"},
	{String: "cQspfSzsgLeiJGB2u8vrAiWpCU4MxUT6JseWo2SjXy4Qbzn2fwDw", Hex: "ef6251e205e8ad508bab5596bee086ef16cd4b239e0cc0c5d7c4e6035441e7d5de01"},
	{String: "19dcawoKcZdQz365WpXWMhX6QCUpR9SY4r", Hex: "005eadaf9bb7121f0f192561a5a62f5e5f54210292"},
	{String: "37Sp6Rv3y4kVd1nQ1JV5pfqXccHNyZm1x3", Hex: "053f210e7277c899c3a155cc1c90f4106cbddeec6e"},
	{String: "myoqcgYiehufrsnnkqdqbp69dddVDMopJu", Hex: "6fc8a3c2a09a298592c3e180f02487cd91ba3400b5"},
	{String: "2N7FuwuUuoTBrDFdrAZ9KxBmtqMLxce9i1C", Hex: "c499b31df7c9068d1481b596578ddbb4d3bd90baeb"},
	{String: "5KL6zEaMtPRXZKo1bbMq7JDjjo1bJuQcsgL33je3oY8uSJCR5b4", Hex: "80c7666842503db6dc6ea061f092cfb9c388448629a6fe868d068c42a488b478ae"},
	{String: "KwV9KAfwbwt51veZWNscRTeZs9CKpojyu1MsPnaKTF5kz69H1UN2", Hex: "8007f0803fc5399e773555ab1e8939907e9badacc17ca129e67a2f5f2ff84351dd01"},
	{String: "93N87D6uxSBzwXvpokpzg8FFmfQPmvX4xHoWQe3pLdYpbiwT5YV", Hex: "efea577acfb5d1d14d3b7b195c321566f12f87d2b77ea3a53f68df7ebf8604a801"},
	{String: "cMxXusSihaX58wpJ3tNuuUcZEQGt6DKJ1wEpxys88FFaQCYjku9h", Hex: "ef0b3b34f0958d8a268193a9814da92c3e8b58b4a4378a542863e34ac289cd830c01"},
	{String: "13p1ijLwsnrcuyqcTvJXkq2ASdXqcnEBLE", Hex: "001ed467017f043e91ed4c44b4e8dd674db211c4e6"},
	{String: "3ALJH9Y951VCGcVZYAdpA3KchoP9McEj1G", Hex: "055ece0cadddc415b1980f001785947120acdb36fc"},
	{String: "0", Hex: "00"},
	{String: "0", Hex: "00000000"},
}

func TestBitcoinEncodingCheck(t *testing.T) {
	for _, pair := range base58BitcoinTestPairs {
		b, err := hex.DecodeString(pair.Hex)
		if err != nil {
			t.Errorf("decoding hex: [%s] %v", pair.Hex, err)
		}
		want := pair.String
		have := BitcoinEncoding.EncodeToString(b)
		if want != have {
			t.Errorf("want: %s have %s", want, have)
		}
	}
}

func TestFlickrEncodingCheck(t *testing.T) {
	for _, pair := range base58FlickrTestPairs {
		b, err := hex.DecodeString(pair.Hex)
		if err != nil {
			t.Errorf("decoding hex: [%s] %v", pair.Hex, err)
		}
		want := pair.String
		have := FlickrEncoding.EncodeToString(b)
		if want != have {
			t.Errorf("want: %s have %s", want, have)
		}
	}
}

func TestBitcoinDecodingCheck(t *testing.T) {
	for _, pair := range base58BitcoinTestPairs {
		b, err := BitcoinEncoding.DecodeString(pair.String)
		if pair.String == "0" {
			if err == nil {
				t.Errorf("decoding address: [0] should have error")
			}
			continue
		}
		if err != nil {
			t.Errorf("decoding address: [%s] %v", pair.String, err)
		}
		want := pair.Hex
		have := hex.EncodeToString(b)
		if want != have {
			t.Errorf("want: %s have %x", want, have)
		}
	}
}

func TestDecodingErrorCheck(t *testing.T) {
	var (
		addr1 = "1Cwvi9VZSR3sXBS1pG59UowQRVc"       // \0ABCDEFGHIJKLMNOPQRS
		addr2 = "15y3"                              // \0AB
		addr3 = "12MBGHp6dG1Ray3cazhw1mer7DLVJDHwG" // \0ABCDEFGHIJKLMNOPQRS
		addr4 = "12MBGHp6dGlRay3cazhw1mer7DLVJDHwG"
		addr5 = ""
	)
	want1 := ErrInvalidChecksum
	_, have1 := BitcoinEncoding.DecodeString(addr1)
	if want1 != have1 {
		t.Errorf("ErrInvalidChecksum:: want: %q have: %q", want1, have1)
	}

	want2 := ErrInvalidChecksumLength
	_, have2 := BitcoinEncoding.DecodeString(addr2)
	if want2 != have2 {
		t.Errorf("ErrInvalidChecksumLength:: want: %q have: %q", want2, have2)
	}

	want3 := ErrUnexpectedEOF
	dst3 := make([]byte, 3)
	_, have3 := BitcoinEncoding.Decode(dst3, []byte(addr3))
	if want3 != have3 {
		t.Errorf("ErrUnexpectedEOF:: want: %q have: %q", want3, have3)
	}

	want4 := fmt.Errorf("invalid base58 digit (%q)", 'l')
	_, have4 := BitcoinEncoding.DecodeString(addr4)
	if want4.Error() != have4.Error() { // because the error objects aren't the same
		t.Errorf("Invalid Base58 Digit:: want: %q have: %q", want4, have4)
	}

	want5 := ErrZeroLength
	_, have5 := BitcoinEncoding.DecodeString(addr5)
	if want5 != have5 { // because the error objects aren't the same
		t.Errorf("ErrZeroLength:: want: %q have: %q", want5, have5)
	}
}

func TestNewEncodingErrorCheck(t *testing.T) {
	defer func() {
		want := "encoding alphabet is not 58-bytes long"
		if have := recover(); have != nil {
			t.Log("Expected panic was called...")
			if want != have {
				t.Errorf("want: %q have: %q", want, have)
			}
		}
	}()
	NewEncoding("ABC123")
}
func TestFlickrDecodingCheck(t *testing.T) {
	for _, pair := range base58FlickrTestPairs {
		b, err := FlickrEncoding.DecodeString(pair.String)
		if pair.String == "0" {
			if err == nil {
				t.Errorf("decoding address: [0] should have error")
			}
			continue
		}
		if err != nil {
			t.Errorf("decoding address: [%s] %v", pair.String, err)
		}
		want := pair.Hex
		have := hex.EncodeToString(b)
		if want != have {
			t.Errorf("want: %s have %x", want, have)
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
		radixDecoding(testPairs[i].String)
	}
}

func BenchmarkFastBase58Decoding(b *testing.B) {
	b.ReportAllocs()

	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		BitcoinEncoding.DecodeString(testPairs[i].String)
	}
}

// Keep radix based endcoding/decoding for benchmark comparisons

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
