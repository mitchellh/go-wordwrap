package wordwrap

import (
	"testing"
)

func TestWrapString(t *testing.T) {
	cases := []struct {
		Input, Output string
		Lim           uint
	}{
		{
			"foo",
			"foo",
			4,
		},
		{
			"foobarbaz",
			"foobarbaz",
			4,
		},
		{
			"foo bar baz",
			"foo\nbar\nbaz",
			4,
		},
		{
			"foo\nb ar\nbaz",
			"foo\nb ar\nbaz",
			4,
		},
		{
			"foo       \nb ar\nbaz",
			"foo\nb ar\nbaz",
			4,
		},
		{
			"foo       \nb ar\nbaz",
			"foo\nb ar\nbaz",
			4,
		},
		{
			"fo sop       \nb ar\n baz",
			"fo sop\nb ar\n baz",
			4,
		},
		{
			" This is a list:\n\n\t* foo\n\t* bar\n\n\n\t* baz\nBAM",
			" This\nis a\nlist:\n\n\t* foo\n\t* bar\n\n\n\t* baz\nBAM",
			4,
		},
	}

	for _, tc := range cases {
		actual := WrapString(tc.Input, tc.Lim)
		if actual != tc.Output {
			t.Fatalf("Input:\n\n%s\n\nActual Output:\n\n%s", tc.Input, actual)
		}
	}
}
