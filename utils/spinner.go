package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

func GetSpinner(type_ []string, time_ time.Duration, suffix string) *spinner.Spinner {
	s := spinner.New(type_, time_)
	s.UpdateSpeed(300 * time.Millisecond)
	s.Suffix = suffix
	return s
}
