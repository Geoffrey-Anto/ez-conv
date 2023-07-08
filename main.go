package main

import (
	"fmt"
	"log"
	"os"

	"github.com/geoffrey-anto/ez-convert/app"
	"github.com/geoffrey-anto/ez-convert/config"
	"github.com/geoffrey-anto/ez-convert/parser"
)

func main() {

	args, err := parser.Parse_Arguments(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	if config.IsDebug() {
		fmt.Println("Input Type: ", args.FromType)
		fmt.Println("Output Type: ", args.ToType)
		fmt.Println("Input File: ", args.Inputs)
	}

	err = app.Run(args)

	if err != nil {
		log.Fatal(err)
	}
}
