package twin

import (
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type interruptableReader struct {
	base *os.File

	interrupted atomic.Bool

	// Ensures we can either read or be paused, but not both at the same time
	pauseOrRead sync.Mutex

	// Re-apply the terminal mode we need. Replaced in tests, which have no
	// terminal to observe; production code always wants reassertTtyInMode.
	reassert func(ttyIn *os.File)
}

// Basically how long we wait between interrupt checks
const interruptableReaderMaxWait = 100 * time.Millisecond

func newInterruptableReader(base *os.File) interruptableReader {
	return interruptableReader{
		base: base,

		reassert: reassertTtyInMode,
	}
}

// Interrupt unblocks the read call, either now or eventually.
func (r *interruptableReader) Interrupt() {
	r.interrupted.Store(true)

	log.Info("Interruptable reader interrupted")
}

func (r *interruptableReader) SetPaused(paused bool) {
	if paused {
		r.pauseOrRead.Lock()
	} else {
		r.pauseOrRead.Unlock()
	}
}

func (r *interruptableReader) Read(p []byte) (n int, err error) {
	for {
		if r.interrupted.Load() {
			log.Info("Interruptable reader already interrupted, returning fabricated EOF")
			return 0, io.EOF
		}

		// Other processes can (and do) reset the terminal mode behind our back.
		// We cannot detect that happening, so we just keep re-applying the mode
		// we need.
		//
		// This must happen even when there is nothing to read: in cooked mode
		// keystrokes are line buffered by the kernel, so they don't make our fd
		// readable until the user presses enter. Waiting with the re-assert
		// until we have something to read would be waiting for input that
		// cannot arrive.
		//
		// The semaphore makes sure we never re-assert while the terminal mode
		// has intentionally been restored, by PauseAndCall() or by Close().
		//
		// Refs:
		//   - https://github.com/walles/moor/issues/443
		//   - https://github.com/walles/moor/issues/394
		if r.pauseOrRead.TryLock() {
			r.reassert(r.base)
			r.pauseOrRead.Unlock()
		}

		// A reset while we're waiting here is fine: the wait is bounded, and
		// the next re-assert picks it up. Keystrokes made in the meantime stay
		// buffered by the kernel until then.
		ready, waitErr := r.waitForReadReady(interruptableReaderMaxWait)
		if waitErr != nil {
			return 0, waitErr
		}

		if !ready {
			continue
		}

		if r.interrupted.Load() {
			log.Info("Interruptable reader interrupted while waiting, returning fabricated EOF")
			return 0, io.EOF
		}

		r.pauseOrRead.Lock()

		// The acquire above can have waited out a whole pause, with every
		// re-assert at the top of the loop skipped throughout it. Re-assert
		// here, because the read below blocks while holding the semaphore: if
		// the mode is wrong by then, the top of the loop won't get to fix it
		// either.
		r.reassert(r.base)

		// Waiting out a pause can also have made our readiness stale: whoever
		// paused us is likely to have consumed the input we were told about.
		// Check again rather than block in the read below while holding the
		// semaphore, which would leave pausing waiting for a keypress that
		// nobody is going to make.
		//
		// Zero timeout, so this is a poll rather than a wait.
		ready, waitErr = r.waitForReadReady(0)
		if waitErr != nil {
			r.pauseOrRead.Unlock()
			return 0, waitErr
		}
		if !ready {
			r.pauseOrRead.Unlock()
			continue
		}

		n, err = r.base.Read(p)
		r.pauseOrRead.Unlock()

		if r.interrupted.Load() {
			log.Info("Interruptable reader interrupted while reading, returning fabricated EOF")
			return 0, io.EOF
		}

		if err == io.EOF {
			log.Info("Interruptable reader base returned a genuine EOF")
		}

		return
	}
}
