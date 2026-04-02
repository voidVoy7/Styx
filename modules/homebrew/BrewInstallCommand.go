package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewInstallCommand(args []string) {
	installFlags := pflag.NewFlagSet("install", pflag.ExitOnError)
	installFlags.Usage = func() {
		fmt.Println("Usage: ow install [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		installFlags.PrintDefaults()
	}

	help := installFlags.BoolP("help", "h", false, "show helpful information")
	verbose := installFlags.BoolP("verbose", "v", false, "show extra output")
	useFormula := installFlags.BoolP("formula", "f", false, "formula to install")
	useCask := installFlags.BoolP("cask", "c", false, "cask to install")

	installFlags.Parse(args)

	if *help {
		installFlags.Usage()
		return
	}

	pkgs := installFlags.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: ow install requires a package name")
		fmt.Println()
		installFlags.Usage()
		return
	}

	brewArgs := []string{"install"}

	if *useFormula {
		brewArgs = append(brewArgs, "--formula")
	}
	if *useCask {
		brewArgs = append(brewArgs, "--cask")
	}

	brewArgs = append(brewArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: brew install")
	}

	cmd := exec.Command("brew", brewArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("brew install failed!", err)
		os.Exit(1)
	}
}
