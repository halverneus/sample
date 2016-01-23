package cli

import (
	"os"
)

// Args parsed from the command-line.
type Args []string

// Parse command-line arguments into Args. Value is returned to support daisy
// chaining.
func (args Args) Parse() Args {
	args = os.Args[1:]
	return args
}

// Matches is used to determine if the arguments match the provided pattern.
func (args Args) Matches(pattern ...string) bool {
	if len(pattern) > len(args) {
		return false
	}

	for index, value := range pattern {
		if "*" != value && value != args[index] {
			return false
		}
	}

	return true
}
