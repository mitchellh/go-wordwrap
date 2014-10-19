package wordwrap

import (
	"testing"
)

func TestWrapString(t *testing.T) {
	cases := []struct{
		Input, Output string
		Lim uint
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
	}

	for _, tc := range cases {
		actual := WrapString(tc.Input, tc.Lim)
		if actual != tc.Output {
			t.Fatalf("Input:\n\n%s\n\nActual Output:\n\n%s", tc.Input, actual)
		}
	}
}
