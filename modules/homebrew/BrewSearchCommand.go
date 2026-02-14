package homebrew

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/pflag"
)

func BrewSearchCommand(args []string) {
	searchFlags := pflag.NewFlagSet("search", pflag.ExitOnError)

	verbose := searchFlags.BoolP("verbose", "v", false, "show extra output")
	useFormula := searchFlags.BoolP("formula", "f", false, "search for formulas/formulae")
	useCask := searchFlags.BoolP("cask", "c", false, "Search for casks")
	useDesc := searchFlags.BoolP("desc", "d", false, "Show searches with extra descriptions")
	useEvalAll := searchFlags.BoolP("eval-all", "e", false, "Search for additional formulae/casks from third party taps")

	searchFlags.Parse(args)

	pkgs := searchFlags.Args()

	if len(pkgs) == 0 {
		fmt.Println("Invalid usage: bowser search requires a package name")
		fmt.Println("A few examples:")
		fmt.Println("  bowser search wget")
		fmt.Println("  bowser search htop")
		fmt.Println("  bowser search google-chrome")
		os.Exit(1)
	}

	brewArgs := []string{"search"}

	if *useFormula {
		brewArgs = append(brewArgs, "--formula")
	}

	if *useCask {
		brewArgs = append(brewArgs, "--cask")
	}

	if *useDesc {
		brewArgs = append(brewArgs, "--desc")
	}

	if *useEvalAll {
		brewArgs = append(brewArgs, "--eval-all")
	}

	brewArgs = append(brewArgs, pkgs...)

	if *verbose {
		fmt.Println("Running: brew search")
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
