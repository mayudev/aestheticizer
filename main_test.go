package main

import "testing"

func TestConvertString(t *testing.T) {
	input := "abcABCdefDEFxyzXYZ2006ąść``~\""
	want := "ａｂｃＡＢＣｄｅｆＤＥＦｘｙｚＸＹＺ２００６ąść｀｀～＂"
	output := convertString(input)

	if output != want {
		t.Fatalf("convertString(%q) = %q, want %q", input, output, want)
	}
}

func TestConvertRune(t *testing.T) {
	input := '!'
	want := '！'
	output := convertRune(input)

	if output != want {
		t.Fatalf("convertRune(%q) = %q, want %q", input, output, want)
	}
}
