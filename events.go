package twin

// Event is the type of value received on a Screen's event channel.
//
// EventRune, EventKeyCode and EventMouse can be constructed directly by
// embedding applications, to programmatically feed input into the screen.
//
// Ref: https://github.com/walles/moor/pull/456
type Event any

// EventRune is sent when the user types a printable rune.
type EventRune struct {
	Rune rune
}

// EventKeyCode is sent when the user presses a non-printable key.
type EventKeyCode struct {
	KeyCode KeyCode
}

// MouseButtonMask is a bitmask of the MouseWheel* constants, used in
// EventMouse.Buttons.
type MouseButtonMask uint16

const (
	// MouseWheelUp is set in EventMouse.Buttons when the wheel scrolls up.
	MouseWheelUp MouseButtonMask = 1 << iota

	// MouseWheelDown is set in EventMouse.Buttons when the wheel scrolls down.
	MouseWheelDown

	// MouseWheelLeft is set in EventMouse.Buttons when the wheel scrolls left.
	MouseWheelLeft

	// MouseWheelRight is set in EventMouse.Buttons when the wheel scrolls right.
	MouseWheelRight
)

// EventMouse is sent on mouse wheel activity.
type EventMouse struct {
	Buttons MouseButtonMask
}

// EventResize is sent when the terminal window is resized. Query Screen.Size()
// after receiving this to get the new size.
type EventResize struct {
	// This interface intentionally left blank
}

// EventExit is sent, and the application should exit, when we're unable to
// continue showing the screen.
//
// Ref: https://github.com/walles/moor/issues/126
type EventExit struct {
	// This interface intentionally left blank
}
