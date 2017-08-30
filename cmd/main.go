package main

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/njones/base58"
)

func main() {
	var (
		err      error
		exitCode int

		help     = flag.Bool("h", false, "display this message")
		lnBreak  = flag.Int("b", 76, "break encoded string into num character lines. Use 0 to disable line wrapping")
		input    = flag.String("i", "-", `input file (use: "-" for stdin)`)
		output   = flag.String("o", "-", `output file (use: "-" for stdout)`)
		decode   = flag.Bool("d", false, `decode input`)
		check    = flag.Bool("k", false, `use sha256 check`)
		useError = flag.Bool("e", false, `write error to stderr`)
	)

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	fin, fout := os.Stdin, os.Stdout
	if *input != "-" {
		if fin, err = os.Open(*input); err != nil {
			fmt.Fprintf(os.Stderr, "input file err: %v\n", err)
			os.Exit(1)
		}
	}

	if *output != "-" {
		if fout, err = os.Create(*output); err != nil {
			fmt.Fprintf(os.Stderr, "output file err: %v\n", err)
			os.Exit(1)
		}
	}

	// separated out for better testing
	err, exitCode = command(fin, fout, decode, check, useError, *lnBreak)
	if err != nil {
		fmt.Fprintf(os.Stderr, "input file err: %v\n", err)
	}
	os.Exit(exitCode)
}

func command(fin io.Reader, fout io.Writer, decode, check, useError *bool, lnBreak int) (err error, code int) {
	var bin, decoded []byte

	if bin, err = ioutil.ReadAll(fin); err != nil {
		return fmt.Errorf("read input err: %v\n", err), 1
	}

	if *decode {
		decodeString := base58.StdEncoding.DecodeString
		if *check {
			decodeString = base58.BitcoinEncoding.DecodeString
		}

		decoded, err = decodeString(strings.TrimSpace(string(bin)))
		if err != nil && err != base58.ErrInvalidChecksum {
			return fmt.Errorf("decode input err: %v\n", err), 1
		}

		io.Copy(fout, bytes.NewReader(decoded))

		if *check && err == base58.ErrInvalidChecksum {
			if *useError {
				return err, 3
			}
			return nil, 3
		}
		return nil, 0
	}

	encodeToString := base58.StdEncoding.EncodeToString
	if *check {
		encodeToString = base58.BitcoinEncoding.EncodeToString
	}

	encoded := encodeToString(bin)
	if lnBreak > 0 {
		lines := (len(encoded) / lnBreak) + 1
		for i := 0; i < lines; i++ {
			start := i * lnBreak
			end := start + lnBreak
			if i == lines-1 {
				fmt.Fprintln(fout, encoded[start:])
				return
			}
			fmt.Fprintln(fout, encoded[start:end])
		}
	}
	fmt.Fprintln(fout, encoded)
	return nil, 0
}

func checkSum(b []byte) []byte {
	sh1, sh2 := sha256.New(), sha256.New()
	sh1.Write(b)
	sh2.Write(sh1.Sum(nil))
	return sh2.Sum(nil)
}
