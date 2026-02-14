package modules

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewRemoveCommand(args []string) {
	removeFlags := pflag.NewFlagSet("remove", pflag.ExitOnError)

	verbose := removeFlags.BoolP("verbose", "v", false, "show extra output")
	useForce := removeFlags.BoolP("force", "r", false, "force the uninstall process")
	useIgnore := removeFlags.BoolP("ignore-dependencies", "i", false, "continue the process even if the package is a dependency of something")
	useZap := removeFlags.BoolP("zap", "z", false, "remove all files associated with a cask")
	useFormula := removeFlags.BoolP("formula", "f", false, "treat arguments as formulas")
	useCask := removeFlags.BoolP("cask", "c", false, "treat arguments as casks")

	removeFlags.Parse(args)

	pkgs := removeFlags.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: bowser remove requires a package name")
		fmt.Println("A few examples:")
		fmt.Println("  bowser remove wget")
		fmt.Println("  bowser remove --force htop")
		fmt.Println("  bowser remove --ignore-dependencies fastfetch")
		fmt.Println("  bowser remove --zap google-chrome")
	}

	brewArgs := []string{"remove"}

	if *useForce {
		brewArgs = append(brewArgs, "--force")
	}
	if *useIgnore {
		brewArgs = append(brewArgs, "--ignore-dependencies")
	}
	if *useZap {
		brewArgs = append(brewArgs, "--zap")
	}
	if *useFormula {
		brewArgs = append(brewArgs, "--formula")
	}
	if *useCask {
		brewArgs = append(brewArgs, "--cask")
	}

	brewArgs = append(brewArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: brew remove")
	}

	cmd := exec.Command("brew", brewArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("brew remove failed", err)
		os.Exit(1)
	}
}
