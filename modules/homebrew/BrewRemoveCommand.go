package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewRemoveCommand(args []string) {
	removeFlags := pflag.NewFlagSet("remove", pflag.ExitOnError)
	removeFlags.Usage = func() {
		fmt.Println("Usage: ow remove [option] <package>")
		fmt.Println()
		fmt.Println("Options:")
		removeFlags.PrintDefaults()
	}

	help := removeFlags.BoolP("help", "h", false, "show helpful information")
	verbose := removeFlags.BoolP("verbose", "v", false, "show extra output")
	useForce := removeFlags.BoolP("force", "r", false, "force the uninstall process")
	useIgnore := removeFlags.BoolP("ignore-dependencies", "i", false, "continue the process even if the package is a dependency of something")
	useZap := removeFlags.BoolP("zap", "z", false, "remove all files associated with a cask")
	useFormula := removeFlags.BoolP("formula", "f", false, "treat arguments as formulas")
	useCask := removeFlags.BoolP("cask", "c", false, "treat arguments as casks")

	removeFlags.Parse(args)

	if *help {
		removeFlags.Usage()
		return
	}

	pkgs := removeFlags.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: ow remove requires a package name")
		fmt.Println()
		removeFlags.Usage()
		return
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
