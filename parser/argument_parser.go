package parser

import (
	"flag"
	"fmt"

	"github.com/geoffrey-anto/ez-convert/utils"
)

type Args struct {
	FromType string
	ToType   string
	Inputs   []string
}

func Parse_Arguments(args []string) (Args, error) {

	fromFlag := flag.String("f", "", "The type to convert from")
	toFlag := flag.String("t", "", "The type to convert to")

	flag.Parse()

	if *fromFlag == "" || *toFlag == "" {
		return Args{}, fmt.Errorf("invalid arguments usage: go run main.go -f <from> -t <to> <inputfiles...>")
	}

	if len(flag.Args()) == 0 {
		return Args{}, fmt.Errorf("no input files specified")
	}

	err := utils.ValidateFromAndTo(*fromFlag, *toFlag)

	if err != nil {
		return Args{}, err
	}

	return Args{
		FromType: *fromFlag,
		ToType:   *toFlag,
		Inputs:   flag.Args(),
	}, nil
}
