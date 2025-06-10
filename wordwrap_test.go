package wordwrap

import (
	"strings"
	"testing"
)

func TestWrapString(t *testing.T) {
	cases := []struct {
		Input, Output, BreakSeq string
		Lim                     uint
	}{
		// A simple word passes through.
		{
			"foo",
			"foo",
			"\n",
			4,
		},
		// A single word that is too long passes through.
		// We do not break words.
		{
			"foobarbaz",
			"foobarbaz",
			"\n",
			4,
		},
		// Lines are broken at whitespace.
		{
			"foo bar baz",
			"foo\nbar\nbaz",
			"\n",
			4,
		},
		// Lines are broken at whitespace, even if words
		// are too long. We do not break words.
		{
			"foo bars bazzes",
			"foo\nbars\nbazzes",
			"\n",
			4,
		},
		// A word that would run beyond the width is wrapped.
		{
			"fo sop",
			"fo\nsop",
			"\n",
			4,
		},
		// Do not break on non-breaking space.
		{
			"foo bar\u00A0baz",
			"foo\nbar\u00A0baz",
			"\n",
			10,
		},
		// Whitespace that trails a line and fits the width
		// passes through, as does whitespace prefixing an
		// explicit line break. A tab counts as one character.
		{
			"foo\nb\t r\n baz",
			"foo\nb\t r\n baz",
			"\n",
			4,
		},
		// Trailing whitespace is removed if it doesn't fit the width.
		// Runs of whitespace on which a line is broken are removed.
		{
			"foo    \nb   ar   ",
			"foo\nb\nar",
			"\n",
			4,
		},
		// An explicit line break at the end of the input is preserved.
		{
			"foo bar baz\n",
			"foo\nbar\nbaz\n",
			"\n",
			4,
		},
		// Explicit break are always preserved.
		{
			"\nfoo bar\n\n\nbaz\n",
			"\nfoo\nbar\n\n\nbaz\n",
			"\n",
			4,
		},
		// Complete example:
		{
			" This is a list: \n\n\t* foo\n\t* bar\n\n\n\t* baz  \nBAM    ",
			" This\nis a\nlist: \n\n\t* foo\n\t* bar\n\n\n\t* baz\nBAM",
			"\n",
			6,
		},
		// Multi-byte characters
		{
			strings.Repeat("\u2584 ", 4),
			"\u2584 \u2584" + "\n" +
				strings.Repeat("\u2584 ", 2),
			"\n",
			4,
		},
		// BASH snippets, to be broken with " \\\n"
		{
			"echo \"deb [signed-by=/etc/apt/keyrings/charm.gpg] https://repo.charm.sh/apt/ * *\" | sudo tee /etc/apt/sources.list.d/charm.list",
			"echo \"deb [signed-by=/etc/apt/keyrings/charm.gpg] https://repo.charm.sh/apt/ * \\\n*\" | sudo tee /etc/apt/sources.list.d/charm.list",
			" \\\n",
			80,
		},
		{
			"curl -fsSL https://repo.charm.sh/apt/gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/charm.gpg",
			"curl -fsSL https://repo.charm.sh/apt/gpg.key | sudo gpg --dearmor -o \\\n/etc/apt/keyrings/charm.gpg",
			" \\\n",
			80,
		},
	}

	for i, tc := range cases {
		actual := WrapString(tc.Input, tc.Lim, tc.BreakSeq)
		if actual != tc.Output {
			t.Fatalf("Case %d Input:\n\n`%s`\n\nExpected Output:\n\n`%s`\n\nActual Output:\n\n`%s`", i, tc.Input, tc.Output, actual)
		}
	}
}
