package main

import (
	"ez-convert/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		panic("Invalid number of arguments")
	}

	inputFile := args[0]
	outputFile := args[1]

	inputFileType, errInput := utils.GetTypeFromFilename(inputFile)
	outputFileType, errOutput := utils.GetTypeFromFilename(outputFile)

	if errInput != nil || errOutput != nil {
		panic("Invalid file type")
	}

	if inputFileType == outputFileType {
		panic("Input and output file types are same")
	}

	switch inputFileType {
	case ".odt":
		switch outputFileType {
		case ".pdf":
			utils.Convert_ODT_2_PDF(inputFile, outputFile)
		default:
			panic("Invalid output file type")

		}
	case ".pdf":
		switch outputFileType {
		case ".odt":
			utils.Convert_PDF_2_ODT(inputFile, outputFile)
		default:
			panic("Invalid output file type")
		}
	case ".pptx":
		switch outputFileType {
		case ".pdf":
			utils.PPT_2_PDF(inputFile, outputFile)
		default:
			panic("Invalid output file type")
		}
	default:
		panic("Invalid input file type")

	}
}
