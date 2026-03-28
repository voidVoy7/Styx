package pacman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func PacmanUpdateCommand(args []string) {
	updateFlag := pflag.NewFlagSet("update", pflag.ExitOnError)
	updateFlag.Usage = func() {
		fmt.Println("Usage: styx update [option]")
		fmt.Println()
		fmt.Println("Options:")
		updateFlag.PrintDefaults()
	}

	help := updateFlag.BoolP("help", "h", false, "show helpful information")
	verbose := updateFlag.BoolP("verbose", "v", false, "show extra output")

	updateFlag.Parse(args)

	pacmanArgs := []string{"-Syu"}

	if *verbose {
		fmt.Println("Running: pacman -Syu")
	}

	if *help {
		updateFlag.Usage()
		return
	}

	cmd := exec.Command("pacman", pacmanArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("pacman -Syu failed:", err)
		os.Exit(1)
	}
}
