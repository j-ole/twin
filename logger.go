package twin

// Logger receives twin's own log messages. Implement it and pass it to
// SetLogger to route those messages wherever you want.
type Logger interface {
	// Debug logs low-level diagnostic messages. This level is recommended for
	// messages that the user has to explicitly ask to see.
	Debug(message string)

	// Info logs high-level status messages. This level and up is recommended
	// for adding to panic reports.
	Info(message string)

	// Error logs a problem the user should be told about whether or not they
	// asked for it.
	Error(message string)
}

type noopLogger struct{}

func (l *noopLogger) Debug(message string) {}

func (l *noopLogger) Info(message string) {}

func (l *noopLogger) Error(message string) {}

var log Logger = &noopLogger{}

// SetLogger routes twin's own log messages to newLogger. Pass nil to disable
// logging.
//
// NOTE: This must be called before any other twin package functions to ensure
// that log messages are not lost.
func SetLogger(newLogger Logger) {
	if newLogger != nil {
		log = newLogger
	} else {
		log = &noopLogger{}
	}
}
