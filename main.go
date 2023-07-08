package main

import (
	"fmt"
	"os"

	"github.com/geoffrey-anto/ez-convert/app"
	"github.com/geoffrey-anto/ez-convert/config"
	"github.com/geoffrey-anto/ez-convert/parser"
	"github.com/geoffrey-anto/ez-convert/utils"
)

func main() {

	args, err := parser.Parse_Arguments(os.Args[1:])

	if err != nil {
		utils.FatalLog(err)
	}

	if config.IsDebug() {
		fmt.Println("Input Type: ", args.FromType)
		fmt.Println("Output Type: ", args.ToType)
		fmt.Println("Input File: ", args.Inputs)
	}

	err = app.Run(args)

	if err != nil {
		utils.FatalLog(err)
	}
}
