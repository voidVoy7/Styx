package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewInstallCommand(args []string) {
	installFlags := pflag.NewFlagSet("install", pflag.ExitOnError)

	verbose := installFlags.BoolP("verbose", "v", false, "show extra output")
	useFormula := installFlags.BoolP("formula", "f", false, "formula to install")
	useCask := installFlags.BoolP("cask", "c", false, "cask to install")

	installFlags.Parse(args)

	pkgs := installFlags.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: bowser install requires a package name")
		fmt.Println("A few examples:")
		fmt.Println("  bowser install wget")
		fmt.Println("  bowser install --formula wget")
		fmt.Println("  bowser install --cask google-chrome")
		os.Exit(1)
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
