package utils

import (
	"fmt"
	"os"
)

func FatalLog(err error) {
	colors := Colors{}
	fmt.Println(string(colors.Red()), "Error: ", err)
	fmt.Println("Conversion aborted.", string(colors.Reset()))
	os.Exit(1)
}
