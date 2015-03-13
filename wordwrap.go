package wordwrap

import (
	"bytes"
	"unicode"
)

// WrapString wraps the given string within lim width in characters.
//
// Wrapping is currently naive and only happens at white-space. A future
// version of the library will implement smarter wrapping. This means that
// pathological cases can dramatically reach past the limit, such as a very
// long word.
func WrapString(s string, lim uint) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)
	implicit := false

	var current uint
	var spaceBuf bytes.Buffer
	for _, char := range s {
		current++

		// If we got a newline, then we honor it and output it. But we have
		// to reset our limit count.
		if char == '\n' {
			if current == 1 && implicit {
				current = 0
				continue
			}
			implicit = false
			goto LINEBREAK
		}

		// Track the whitespace in our whitespace buffer.
		if unicode.IsSpace(char) {
			// Consume any whitespace if we just linebroke implicitly.
			if current == 1 && implicit {
				current = 0
				continue
			}

			// If we're over the limit already, then output a newline
			// and reset.
			if current > lim {
				implicit = true
				goto LINEBREAK
			}

			// If this whitespace would put us over the limit, break
			if current+uint(spaceBuf.Len()) >= lim {
				implicit = true
				goto LINEBREAK
			}

			spaceBuf.WriteRune(char)
			continue
		}

		// Output our buffered whitespace if we have any
		if spaceBuf.Len() > 0 {
			if _, err := spaceBuf.WriteTo(buf); err != nil {
				panic(err)
			}
			current += uint(spaceBuf.Len())
		}

		buf.WriteRune(char)

		continue
	LINEBREAK:
		buf.WriteRune('\n')
		current = 0
		spaceBuf.Reset()
	}

	return buf.String()
}
