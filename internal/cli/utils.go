package cli

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/common-nighthawk/go-figure"
)

func clearTerminal() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func clearAndPrintHeader() {
	clearTerminal()
	myFigure := figure.NewFigure(CliName, "", true)
	myFigure.Print()
}

func addReturnOption(items []string, first bool) ([]string, int) {
	returnOption := ReturnOption
	if first {
		returnOption = ExitOption
	}
	itemsWithReturn := append(items, returnOption)
	returnIndex := len(items)
	return itemsWithReturn, returnIndex
}
