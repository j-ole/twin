package twin

// KeyCode identifies a non-printable key the user pressed
type KeyCode uint16

// KeyCode values for non-printable keys, sent via EventKeyCode
const (
	KeyEscape KeyCode = iota
	KeyEnter

	KeyBackspace
	KeyDelete

	KeyUp
	KeyDown
	KeyRight
	KeyLeft

	KeyAltUp
	KeyAltDown
	KeyAltRight
	KeyAltLeft

	KeyHome
	KeyEnd
	KeyPgUp
	KeyPgDown

	RuneKeyAltLeft
	RuneKeyAltRight
	KeyF1
)

// Map incoming escape keystrokes to keycodes, used in consumeEncodedEvent() in
// screen.go.
var escapeSequenceToKeyCode = map[string]KeyCode{
	// NOTE: Please keep this list in the same order as the KeyCode const()
	// section above.

	// KeyEscape intentionally left out: it's a single-byte prefix of every
	// other entry here, so consumeEncodedEvent() handles it separately,
	// after this map fails to match.

	// KeyEnter intentionally left out: consumeEncodedEvent() handles '\r'
	// separately, after this map fails to match.

	"\x7f":    KeyBackspace,
	"\x1b[3~": KeyDelete,

	"\x1b[A": KeyUp,
	"\x1b[B": KeyDown,
	"\x1b[C": KeyRight,
	"\x1b[D": KeyLeft,

	// Ref: https://github.com/walles/moor/issues/138#issuecomment-1579199274
	"\x1bOA": KeyUp,
	"\x1bOB": KeyDown,
	"\x1bOC": KeyRight,
	"\x1bOD": KeyLeft,

	"\x1b\x1b[A": KeyAltUp,    // Alt + up arrow
	"\x1b\x1b[B": KeyAltDown,  // Alt + down arrow
	"\x1b\x1b[C": KeyAltRight, // Alt + right arrow
	"\x1b\x1b[D": KeyAltLeft,  // Alt + left arrow

	// macBook option-arrows
	"\x1b[1;3A": KeyAltUp,
	"\x1b[1;3B": KeyAltDown,
	"\x1b[1;3C": KeyAltRight,
	"\x1b[1;3D": KeyAltLeft,

	"\x1b[H":  KeyHome,
	"\x1b[F":  KeyEnd,
	"\x1b[1~": KeyHome,
	"\x1b[4~": KeyEnd,
	"\x1b[5~": KeyPgUp,
	"\x1b[6~": KeyPgDown,

	"\x1bl":  RuneKeyAltRight, // Alt + l
	"\x1bh":  RuneKeyAltLeft,  // Alt + h
	"\x1bOP": KeyF1,
}
