package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	applicationMap := map[rune]string{
		'🔍': "search",
		'❗': "recommendation",
		'☀': "weather",
	}

	for _, char := range log {
		application, ok := applicationMap[char]
		if ok {
			return application
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog []rune

	for _, char := range log {
		if char == oldRune {
			newLog = append(newLog, newRune)
		} else {
			newLog = append(newLog, char)
		}
	}

	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
