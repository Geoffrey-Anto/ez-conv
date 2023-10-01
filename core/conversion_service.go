package core

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/geoffrey-anto/ez-convert/parser"
)

func PDF_to_DOCX(args parser.Args) error {
	from := args.FromType
	to := args.ToType

	files := args.Inputs

	for _, file := range files {
		if strings.Split(file, ".")[1] != from {
			return fmt.Errorf("invalid file type, on of the types has been provided as %s must be %s", strings.Split(file, ".")[1], from)
		}
	}

	switch runtime.GOOS {
	case "windows":
		for _, file := range files {
			err := exec.Command("powershell", "soffice", "--infilter=\"writer_pdf_import\"", "--convert-to", to, file).Run()

			if err != nil {
				return err
			}
		}

	default: //Mac & Linux
		for _, file := range files {
			err := exec.Command("soffice", "--infilter=\"writer_pdf_import\"", "--convert-to", to, file).Run()

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func PDF_to_PPTX() {

}

func DOCX_to_PDF() {

}

func PPTX_to_PDF() {

}
