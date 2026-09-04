package twin

import (
	"strings"
	"testing"
)

// consumeEncodedEvent() (screen.go) matches escapeSequenceToKeyCode entries
// with strings.HasPrefix() while ranging over the map, and Go randomizes map
// iteration order. So if one entry's key were a prefix of another entry's key,
// which one wins the match would depend on that random order, and matching the
// shorter one would leave the rest of the longer sequence to be misparsed as
// something else entirely.
//
// This test keeps that from happening by asserting no key is a prefix of any
// other key.
func TestEscapeSequenceToKeyCodeNoPrefixCollisions(t *testing.T) {
	for candidate := range escapeSequenceToKeyCode {
		for other := range escapeSequenceToKeyCode {
			if candidate == other {
				continue
			}

			if !strings.HasPrefix(other, candidate) {
				continue
			}

			t.Errorf("%q is a prefix of %q, matching between them would depend on map iteration order", candidate, other)
		}
	}
}
