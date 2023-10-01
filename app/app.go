package app

import (
	"fmt"

	"github.com/geoffrey-anto/ez-convert/core"
	"github.com/geoffrey-anto/ez-convert/parser"
	"github.com/geoffrey-anto/ez-convert/utils"
)

func Run(args parser.Args) error {
	colors := utils.Colors{}
	fmt.Println(string(colors.Purple()), " Welcome to ez-convert! Please wait while we convert your files...")
	fmt.Print(string(colors.Reset()), " Press y to continue, or any other key to abort.  ")

	if !utils.CanProceed() {
		fmt.Println("Conversion aborted.")
		return fmt.Errorf("conversion aborted")
	} else {
		err := core.Convert(args)
		if err != nil {
			return err
		}
		return nil
	}
}
