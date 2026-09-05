package twin

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestRuneWidth(t *testing.T) {
	assert.Equal(t, StyledRune{Rune: 'x', Style: Style{}}.Width(), 1)
	assert.Equal(t, StyledRune{Rune: '午', Style: Style{}}.Width(), 2)
}

// Go's unicode tables lag behind the latest Unicode release, so
// unicode.IsPrint() does not recognize blocks added in Unicode 15.1 (2023),
// 16.0 (2024), and 17.0 (2025). Printable() should still consider characters
// in these blocks printable.
//
// Ported from https://github.com/walles/moor/pull/408.
func TestPrintableUnicodePost15(t *testing.T) {
	cases := []struct {
		name string
		r    rune
	}{
		// Unicode 15.1 (2023)
		{"CJK Ext I start", 0x2EBF0},
		{"CJK Ext I end", 0x2EE5F},

		// Unicode 16.0 (2024)
		{"Todhri start", 0x105C0},
		{"Todhri end", 0x105F3},
		{"Garay start", 0x10D40},
		{"Garay end", 0x10D8E},
		{"Tulu-Tigalari start", 0x11380},
		{"Tulu-Tigalari end", 0x113D5},
		{"Sunuwar start", 0x11BC0},
		{"Sunuwar end", 0x11BF2},
		{"Egyptian Hieroglyphs Ext-A start", 0x13460},
		{"Egyptian Hieroglyphs Ext-A end", 0x143FA},
		{"Gurung Khema start", 0x16100},
		{"Gurung Khema end", 0x16139},
		{"Kirat Rai start", 0x16D40},
		{"Kirat Rai end", 0x16D79},
		{"Legacy Computing Supplement start", 0x1CC00},
		{"Large Type Piece (used by jj)", 0x1CE1A},
		{"Large Type Piece end", 0x1CE50},
		{"Legacy Computing Supplement end", 0x1CEBF},
		{"Ol Onal start", 0x1E5D0},
		{"Ol Onal end", 0x1E5FA},

		// Unicode 17.0 (2025)
		{"Sidetic start", 0x10940},
		{"Sidetic end", 0x1095F},
		{"Sharada Supplement start", 0x11B60},
		{"Sharada Supplement end", 0x11B7F},
		{"Tolong Siki start", 0x11DB0},
		{"Tolong Siki end", 0x11DEF},
		{"Beria Erfe start", 0x16EA0},
		{"Beria Erfe end", 0x16EDF},
		{"Tangut Components Supplement start", 0x18D80},
		{"Tangut Components Supplement end", 0x18DFF},
		{"Misc Symbols Supplement start", 0x1CEC0},
		{"Misc Symbols Supplement end", 0x1CEFF},
		{"Tai Yo start", 0x1E6C0},
		{"Tai Yo end", 0x1E6FF},
		{"CJK Ext J start", 0x323B0},
		{"CJK Ext J end", 0x3347F},
	}

	for _, tc := range cases {
		assert.Assert(t, Printable(tc.r),
			"expected U+%04X (%s) to be printable", tc.r, tc.name)
	}
}
