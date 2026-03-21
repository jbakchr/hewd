package cliutils

import (
	"fmt"
	"strings"
)

// HewdError is a structured error type used across all commands.
// It separates the main error message from an optional hint.
//
// Msg  -> the actual error message (clean, lowercase, no trailing period)
// Hint -> optional hint for the user (will be shown on a new line)
//
// Example printing:
//
//	error: cannot combine --json and --yaml
//	hint: use only one machine-readable format at a time
type HewdError struct {
	Msg  string
	Hint string
}

func (e HewdError) Error() string {
	return e.Msg
}

// Err creates a basic HewdError with no hint.
func Err(message string) error {
	message = strings.TrimSpace(message)
	return HewdError{
		Msg: message,
	}
}

// ErrHint creates a HewdError with both message and hint.
func ErrHint(message, hint string) error {
	message = strings.TrimSpace(message)
	hint = strings.TrimSpace(hint)
	return HewdError{
		Msg:  message,
		Hint: hint,
	}
}

// RootErr formats a top-level HewdError with command prefix.
// Used ONLY in main.go when rootCmd.Execute() returns an error.
func RootErr(commandName, message, hint string) string {
	message = strings.TrimSpace(message)
	hint = strings.TrimSpace(hint)

	// red error prefix
	prefix := fmt.Sprintf("%serror (hewd %s):%s", Red, commandName, Reset)

	if hint == "" {
		return fmt.Sprintf("%s %s", prefix, message)
	}

	return fmt.Sprintf("%s %s\nhint: %s", prefix, message, hint)
}
