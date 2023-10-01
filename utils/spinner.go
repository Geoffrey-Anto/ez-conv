package utils

import (
	"time"

	"github.com/briandowns/spinner"
)

func GetSpinner(type_ []string, time_ time.Duration, suffix string, end_text string) *spinner.Spinner {
	s := spinner.New(type_, time_)
	s.UpdateSpeed(300 * time.Millisecond)
	color := Colors{}
	s.Suffix = color.Blue() + " " + suffix
	s.FinalMSG = color.Green() + end_text + color.Reset()
	return s
}
