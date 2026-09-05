package twin

import (
	"fmt"
	"slices"
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

// unicodeRange is a contiguous, inclusive range of Unicode code points.
type unicodeRange struct {
	lo, hi rune
}

// Blocks added in Unicode 15.1 (2023), 16.0 (2024), and 17.0 (2025).
//
// Go's unicode package ships whatever Unicode version was current when that Go
// version was released, so older Go toolchains can be out of sync with the
// latest Unicode standard and unicode.IsPrint() won't recognize these blocks.
// Most code points in these blocks are assigned real characters, so on such a
// toolchain we'd rather pass all of them through than mask real characters with
// our own '?' just because a few code points in the block are still unassigned.
//
// Must be sorted by `lo` ascending, with no overlaps — enforced by
// TestUnicodePost15PrintableRangesSorted.
var unicodePost15PrintableRanges = []unicodeRange{
	{0x105C0, 0x105FF}, // Todhri (16.0)
	{0x10940, 0x1095F}, // Sidetic (17.0)
	{0x10D40, 0x10D8F}, // Garay (16.0)
	{0x11380, 0x113FF}, // Tulu-Tigalari (16.0)
	{0x116D0, 0x116FF}, // Myanmar Extended-C (16.0)
	{0x11B60, 0x11B7F}, // Sharada Supplement (17.0)
	{0x11BC0, 0x11BFF}, // Sunuwar (16.0)
	{0x11DB0, 0x11DEF}, // Tolong Siki (17.0)
	{0x13460, 0x143FF}, // Egyptian Hieroglyphs Extended-A (16.0)
	{0x16100, 0x1613F}, // Gurung Khema (16.0)
	{0x16D40, 0x16D7F}, // Kirat Rai (16.0)
	{0x16EA0, 0x16EDF}, // Beria Erfe (17.0)
	{0x18D80, 0x18DFF}, // Tangut Components Supplement (17.0)
	{0x1CC00, 0x1CEFF}, // Symbols for Legacy Computing Supplement (16.0) + Misc Symbols Supplement (17.0)
	{0x1E5D0, 0x1E5FF}, // Ol Onal (16.0)
	{0x1E6C0, 0x1E6FF}, // Tai Yo (17.0)
	{0x2EBF0, 0x2EE5F}, // CJK Unified Ideographs Extension I (15.1)
	{0x323B0, 0x3347F}, // CJK Unified Ideographs Extension J (17.0)
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

	_, found := slices.BinarySearchFunc(unicodePost15PrintableRanges, char,
		func(r unicodeRange, char rune) int {
			switch {
			case char < r.lo:
				return 1
			case char > r.hi:
				return -1
			default:
				return 0
			}
		})

	return found
}
