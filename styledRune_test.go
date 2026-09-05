package twin

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestRuneWidth(t *testing.T) {
	assert.Equal(t, StyledRune{Rune: 'x', Style: Style{}}.Width(), 1)
	assert.Equal(t, StyledRune{Rune: '午', Style: Style{}}.Width(), 2)
}

// Added after a moor user reported that recently-added Unicode characters
// (in their case, Legacy Computing Supplement glyphs used by jj) were
// rendering as '?'.
//
// This checks the boundaries of our own unicodePost15PrintableRanges table,
// not "the last real character Unicode has assigned in this block" — that
// shifts with every Unicode release. The boundaries below are true by
// construction, independent of whichever Unicode version the running Go
// toolchain knows about.
//
// Ported from https://github.com/walles/moor/pull/408.
func TestPrintableUnicodePost15(t *testing.T) {
	for _, r := range unicodePost15PrintableRanges {
		assert.Assert(t, Printable(r.lo), "expected U+%04X (range start) to be printable", r.lo)
		assert.Assert(t, Printable(r.hi), "expected U+%04X (range end) to be printable", r.hi)
	}

	// jj's motivating use case: Large Type Piece glyphs from the Symbols for
	// Legacy Computing Supplement block (Unicode 16.0), known to be assigned
	// real characters.
	assert.Assert(t, Printable(0x1CE1A), "expected Large Type Piece U+1CE1A to be printable")
}

// Printable()'s binary search over unicodePost15PrintableRanges assumes the
// table is sorted by `lo` with no overlaps. Guard that invariant, since a
// violation would make the binary search silently miss some runes rather than
// fail loudly.
func TestUnicodePost15PrintableRangesSorted(t *testing.T) {
	prevHi := rune(-1)
	for _, r := range unicodePost15PrintableRanges {
		assert.Assert(t, r.lo > prevHi,
			"range %X..%X overlaps or is out of order with previous (hi=%X)",
			r.lo, r.hi, prevHi)
		assert.Assert(t, r.lo <= r.hi,
			"range %X..%X has lo > hi", r.lo, r.hi)
		prevHi = r.hi
	}
}
