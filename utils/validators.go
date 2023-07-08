package utils

import "fmt"

func CanProceed() bool {
	var canproceed string
	fmt.Scanln(&canproceed)
	return canproceed == "y"
}

func ValidateFromAndTo(from string, to string) error {
	if from == to {
		return fmt.Errorf("from and to cannot be the same")
	}

	if from != "pdf" && from != "docx" && from != "pptx" {
		return fmt.Errorf("invalid from type, must be pdf, docx or pptx")
	}

	if to != "pdf" && to != "docx" && to != "pptx" {
		return fmt.Errorf("invalid to type, must be pdf, docx or pptx")
	}

	return nil
}
