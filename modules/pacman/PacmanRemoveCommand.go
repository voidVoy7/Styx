package pacman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func PacmanRemoveCommand(args []string) {
	removeFlag := pflag.NewFlagSet("remove", pflag.ExitOnError)
	removeFlag.Usage = func() {
		fmt.Println("Usage: styx remove [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		removeFlag.PrintDefaults()
	}

	help := removeFlag.BoolP("help", "h", false, "show helpful information")
	verbose := removeFlag.BoolP("verbose", "v", false, "show extra output")

	removeFlag.Parse(args)

	if *help {
		removeFlag.Usage()
		return
	}

	pkgs := removeFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: styx remove requires a package name")
		fmt.Println()
		removeFlag.Usage()
		return
	}

	pacmanArgs := []string{"-Rcns"}

	pacmanArgs = append(pacmanArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: pacman -Rcns")
	}

	cmd := exec.Command("pacman", pacmanArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("pacman -Rcns failed:", err)
		os.Exit(1)
	}
}
