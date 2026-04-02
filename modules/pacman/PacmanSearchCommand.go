package pacman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func PacmanSearchCommand(args []string) {
	searchFlag := pflag.NewFlagSet("search", pflag.ExitOnError)
	searchFlag.Usage = func() {
		fmt.Println("Usage: ow search [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		searchFlag.PrintDefaults()
	}

	help := searchFlag.BoolP("help", "h", false, "show helpful information")
	verbose := searchFlag.BoolP("verbose", "v", false, "show extra output")

	searchFlag.Parse(args)

	if *help {
		searchFlag.Usage()
		return
	}

	pkgs := searchFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: ow search requires a package name")
		fmt.Println()
		searchFlag.Usage()
		return
	}

	pacmanArgs := []string{"-Ss"}

	pacmanArgs = append(pacmanArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: pacman -Ss")
	}

	cmd := exec.Command("pacman", pacmanArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("pacman -Ss failed", err)
		os.Exit(1)
	}
}
