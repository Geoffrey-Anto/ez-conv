package app

import (
	"fmt"

	"github.com/geoffrey-anto/ez-convert/core"
	"github.com/geoffrey-anto/ez-convert/parser"
	"github.com/geoffrey-anto/ez-convert/utils"
)

func Run(args parser.Args) error {
	fmt.Println("Welcome to ez-convert! Please wait while we convert your files...")
	fmt.Println("Press y to continue, or any other key to abort.")

	if !utils.CanProceed() {
		fmt.Println("Conversion aborted.")
		return fmt.Errorf("conversion aborted")
	} else {
		core.Convert(args)
		return nil
	}
}
