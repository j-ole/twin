package twin

import (
	"fmt"
	"unicode"

	"github.com/rivo/uniseg"
)

// StyledRune is a rune with a style to be written to a one or more cells on the
// screen. Note that a StyledRune may use more than one cell on the screen ('午'
// for example).
type StyledRune struct {
	Rune  rune
	Style Style
}

func (styledRune StyledRune) String() string {
	return fmt.Sprint("rune='", string(styledRune.Rune), "' ", styledRune.Style)
}

// Width returns how many screen cells this rune will cover. Most runes cover
// one, but some like '午' will cover two.
func (styledRune StyledRune) Width() int {
	return uniseg.StringWidth(string(styledRune.Rune))
}

// Equal reports whether styledRune and other have the same rune and style.
func (styledRune StyledRune) Equal(other StyledRune) bool {
	return styledRune.Rune == other.Rune && styledRune.Style.Equal(other.Style)
}

// Printable reports whether char should be rendered as-is rather than
// escaped, covering some cases that unicode.IsPrint() gets wrong for
// terminal output.
func Printable(char rune) bool {
	if unicode.IsPrint(char) {
		return true
	}

	if unicode.Is(unicode.Co, char) {
		// Co == "Private Use": https://www.compart.com/en/unicode/category
		//
		// This space is used by Font Awesome, for "fa-battery-empty" for
		// example: https://fontawesome.com/v4/icon/battery-empty
		//
		// So we want to print these and let the rendering engine deal with
		// outputting them in a way that's helpful to the user.
		return true
	}

	if char == 0xa0 {
		// 0xa0 is a non-breaking space, which is printable, despite what
		// unicode.IsPrint() says.
		return true
	}

	return false
}
