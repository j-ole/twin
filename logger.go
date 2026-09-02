package twin

// Logger receives twin's log messages. Implement it and pass it in
// Options.Logger to NewScreen to consume those messages.
type Logger interface {
	// Debug logs low-level diagnostic messages. This level is recommended for
	// messages that the user has to explicitly ask to see.
	Debug(message string)

	// Info logs high-level status messages. This level and up is recommended
	// for adding to panic reports.
	Info(message string)

	// Error logs a problem the user should be told about whether they asked for
	// it or not.
	Error(message string)
}

type noopLogger struct{}

func (l *noopLogger) Debug(message string) {}

func (l *noopLogger) Info(message string) {}

func (l *noopLogger) Error(message string) {}

var log Logger = &noopLogger{}
