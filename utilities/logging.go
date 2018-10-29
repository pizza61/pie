package utilities

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Log log in blue
func Log(text string) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("> %s %s\n", blue("[LOG]"), text)
}

func LogErr(text string) {
	red := color.New(color.FgHiRed).SprintFunc()
	fmt.Printf("> %s %s\n", red("[ERROR]"), text)
}

func LogWarn(text string) {
	yello := color.New(color.FgHiYellow).SprintFunc()
	fmt.Printf("> %s %s\n", yello("[WARN]"), text)
}
func CheckErr(err error) {
	if err != nil {
		LogErr(err.Error())
		os.Exit(1)
	}
}
