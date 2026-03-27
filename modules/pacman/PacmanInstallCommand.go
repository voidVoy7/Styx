package pacman

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func PacmanInstallCommand(args []string) {
	installFlag := pflag.NewFlagSet("install", pflag.ExitOnError)
	installFlag.Usage = func() {
		fmt.Println("Usage: styx install [option] <package>")
		fmt.Println()
		fmt.Println("Options")
		installFlag.PrintDefaults()
	}

	help := installFlag.BoolP("help", "h", false, "show helpful information")
	verbose := installFlag.BoolP("verbose", "v", false, "show extra output")

	installFlag.Parse(args)

	if *help {
		installFlag.Usage()
		return
	}

	pkgs := installFlag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: styx install requires a package name")
		fmt.Println()
		installFlag.Usage()
		return
	}

	pacmanArgs := []string{"-S"}

	pacmanArgs = append(pacmanArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: pacman -S")
	}

	cmd := exec.Command("pacman", pacmanArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Pacman -S failed:", err)
		os.Exit(1)
	}
}
