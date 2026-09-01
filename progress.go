package twin

import (
	"fmt"
	"strconv"
)

// ProgressState is a terminal progress bar's state, one of the
// ProgressState* constants.
type ProgressState int

// ProgressState values. The numbers match the ones from
// https://rockorager.dev/misc/osc-9-4-progress-bars/.
const (
	// ProgressStateRemove hides the progress bar.
	ProgressStateRemove ProgressState = 0

	// ProgressStateSet shows the progress bar at Progress.Percent.
	ProgressStateSet ProgressState = 1

	// ProgressStateError shows the progress bar, in an error color, at
	// Progress.Percent.
	ProgressStateError ProgressState = 2

	// ProgressStateIndeterminate shows a busy progress bar with no known
	// percentage.
	ProgressStateIndeterminate ProgressState = 3

	// ProgressStatePause shows the progress bar, in a paused color, at
	// Progress.Percent.
	ProgressStatePause ProgressState = 4
)

// Tell the terminal to remove the progress bar.
//
// See renderProgress() below for details.
const progressRemoveSequence = "\x1b]9;4;0\x07"

// Progress is a terminal progress bar's state and completion percentage.
//
// Ref: https://rockorager.dev/misc/osc-9-4-progress-bars/
type Progress struct {
	State   ProgressState
	Percent int
}

// SetProgress sets the terminal's progress bar to state, showing percent
// complete. percent is ignored for ProgressStateRemove and
// ProgressStateIndeterminate, and clamped to 0-100 otherwise.
//
// Panics if state is not one of the ProgressState* constants.
func (screen *UnixScreen) SetProgress(state ProgressState, percent int) {
	if state < ProgressStateRemove || state > ProgressStatePause {
		panic(fmt.Errorf("invalid progress state: %d", state))
	}

	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}

	screen.renderLock.Lock()
	defer screen.renderLock.Unlock()

	screen.progress = Progress{
		State:   state,
		Percent: percent,
	}
}

// You must hold renderLock when calling this method.
func (screen *UnixScreen) renderProgressLocked() string {
	osc := "\x1b]9;4;"
	osc += strconv.Itoa(int(screen.progress.State))

	if screen.progress.State != ProgressStateRemove && screen.progress.State != ProgressStateIndeterminate {
		osc += ";"
		osc += strconv.Itoa(screen.progress.Percent)
	}

	osc += "\x07"

	return osc
}
