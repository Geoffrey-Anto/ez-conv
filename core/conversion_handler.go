package core

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/geoffrey-anto/ez-convert/parser"
	"github.com/geoffrey-anto/ez-convert/utils"
)

func Convert(args parser.Args) error {
	from := args.FromType
	to := args.ToType

	s := utils.GetSpinner(spinner.CharSets[35], 2*time.Second, "  Converting files from "+from+" to "+to+"...")
	s.Start()

	// TODO: Implement conversion logic here
	time.Sleep(4 * time.Second)
	// TODO: Implement conversion logic here

	s.Stop()

	return nil
}
