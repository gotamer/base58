package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSimple(t *testing.T) {
	have := "Hello world"
	want := "JxF12TrwXzT5jvT\n"
	gather := new(bytes.Buffer)

	var decode, check, useError bool
	var doDecode = !decode

	err, code := command(strings.NewReader(have), gather, &decode, &check, &useError, 200)
	if err != nil {
		t.Errorf("command err:", err)
	}
	if code != 0 {
		t.Errorf("code not 0:", code)
	}
	if gather.String() != want {
		t.Errorf("want: %q have: %q", gather.String(), want)
	}

	gather.Reset()
	err1, code1 := command(strings.NewReader(want), gather, &doDecode, &check, &useError, 200)
	if err1 != nil {
		t.Errorf("command err:", err1)
	}
	if code != 0 {
		t.Errorf("code not 0:", code1)
	}

	if gather.String() != have {
		t.Errorf("want: %q have: %q", gather.String(), have)
	}
}

func TestSimpleCheck(t *testing.T) {
	have := "Hello world"
	want := "32UWxgjUJd9s6KywDtjJL\n"
	gather := new(bytes.Buffer)

	var decode, check, useError bool
	check = true
	var doDecode = !decode

	err, code := command(strings.NewReader(have), gather, &decode, &check, &useError, 200)
	if err != nil {
		t.Errorf("command err:", err)
	}
	if code != 0 {
		t.Errorf("code not 0:", code)
	}
	if gather.String() != want {
		t.Errorf("want: %q have: %q", gather.String(), want)
	}

	gather.Reset()
	err1, code1 := command(strings.NewReader(want), gather, &doDecode, &check, &useError, 200)
	if err1 != nil {
		t.Errorf("command err:", err1)
	}
	if code != 0 {
		t.Errorf("code not 0:", code1)
	}

	if gather.String() != have {
		t.Errorf("want: %q have: %q", gather.String(), have)
	}
}
