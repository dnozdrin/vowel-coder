package main

import "testing"

var encodeTests = []struct {
	in, out string
}{
	{"aeiou", "12345"},
	{"hello", "h2ll4"},
	{"hi there", "h3 th2r2"},
	{"Hello", "h2ll4"},
	{"HI THERE", "h3 th2r2"},
}

func TestEncode(t *testing.T) {
	testTextProcessing(t, encodeTests, encoderMap)
}

func testTextProcessing(t *testing.T, samples []struct{ in, out string }, mapping []string) {
	for _, tt := range samples {
		t.Run(tt.in, func(t *testing.T) {
			s := processText(mapping, tt.in)
			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

var decodeTests = []struct {
	in, out string
}{
	{"12345", "aeiou"},
	{"h2ll4", "hello"},
	{"h3 th2r2", "hi there"},
}

func TestDecode(t *testing.T) {
	testTextProcessing(t, decodeTests, decoderMap)
}
