package core

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/geoffrey-anto/ez-convert/parser"
	"github.com/geoffrey-anto/ez-convert/utils"
)

func Convert(args parser.Args) error {
	from := args.FromType
	to := args.ToType

	s := utils.GetSpinner(spinner.CharSets[9], 2*time.Second, "  Converting files from "+from+" to "+to+"...", "  Conversion complete!\n")
	s.Start()

	switch from {
	case "pdf":
		switch to {
		case "docx":
			err := PDF_to_DOCX(args)
			if err != nil {
				return err
			}
		case "pptx":
			PDF_to_PPTX()
		default:
			return fmt.Errorf("invalid to type, must be docx or pptx")
		}
	case "docx":
		switch to {
		case "pdf":
			DOCX_to_PDF()
		default:
			return fmt.Errorf("invalid to type, must be pdf")
		}
	case "pptx":
		switch to {
		case "pdf":
			PPTX_to_PDF()
		default:
			return fmt.Errorf("invalid to type, must be pdf")
		}
	}

	s.Stop()

	return nil
}
